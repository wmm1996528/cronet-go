package cronet

import (
	"context"
	"errors"
	"io"
	"net"
	"net/http"
	"net/textproto"
	"os"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"sync"
)

// RoundTripper is a wrapper from URLRequest to http.RoundTripper
type RoundTripper struct {
	FollowRedirect bool
	Engine         Engine
	Executor       Executor

	closeEngine   bool
	closeExecutor bool
}

func NewCronetTransport(params EngineParams, FollowRedirect bool) *RoundTripper {
	t := &RoundTripper{
		FollowRedirect: FollowRedirect,
	}
	t.Engine = NewEngine()
	t.Engine.StartWithParams(params)
	params.Destroy()
	t.closeEngine = true

	t.Executor = NewExecutor(func(executor Executor, command Runnable) {
		go func() {
			command.Run()
			command.Destroy()
		}()
	})
	t.closeExecutor = true
	runtime.SetFinalizer(t, (*RoundTripper).Close)
	return t
}

func NewCronetTransportWithDefaultParams() *RoundTripper {
	engineParams := NewEngineParams()
	engineParams.SetEnableHTTP2(true)
	engineParams.SetEnableQuic(true)
	engineParams.SetEnableBrotli(true)
	engineParams.SetUserAgent("Go-cronet-http-client")
	return NewCronetTransport(engineParams, true)
}

func (t *RoundTripper) Close() error {
	if t.closeEngine {
		result := t.Engine.Shutdown()
		if result != ResultSuccess {
			return errors.New("engine still has active requests, so couldn't shutdown")
		}
		t.Engine.Destroy()
	}
	if t.closeExecutor {
		t.Executor.Destroy()
	}
	return nil
}

func (t *RoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {

	requestParams := NewURLRequestParams()
	if request.Method == "" {
		requestParams.SetMethod("GET")
	} else {
		requestParams.SetMethod(request.Method)
	}
	for key, values := range request.Header {
		for _, value := range values {
			if len(value) == 0 {
				continue
			}
			header := NewHTTPHeader()
			header.SetName(key)
			header.SetValue(value)
			requestParams.AddHeader(header)
			header.Destroy()

		}
	}
	if request.Body != nil {
		uploadProvider := NewUploadDataProvider(&bodyUploadProvider{request.Body, request.GetBody, request.ContentLength})
		requestParams.SetUploadDataProvider(uploadProvider)
		requestParams.SetUploadDataExecutor(t.Executor)
	}
	m := &sync.Mutex{}
	responseHandler := urlResponse{
		FollowRedirect: t.FollowRedirect,
		response: http.Response{
			Request:    request,
			Proto:      request.Proto,
			ProtoMajor: request.ProtoMajor,
			ProtoMinor: request.ProtoMinor,
			Header:     make(http.Header),
		},
		complete: sync.NewCond(m),
		read:     make(chan int),
		cancel:   make(chan struct{}),
		done:     make(chan struct{}),
	}
	responseHandler.response.Body = &responseHandler
	go responseHandler.monitorContext(request.Context())

	callback := NewURLRequestCallback(&responseHandler)
	urlRequest := NewURLRequest()
	responseHandler.request = urlRequest
	urlRequest.InitWithParams(t.Engine, request.URL.String(), requestParams, callback, t.Executor)
	requestParams.Destroy()
	urlRequest.Start()
	m.Lock()
	responseHandler.complete.Wait()
	return &responseHandler.response, responseHandler.err
}

type urlResponse struct {
	FollowRedirect bool

	complete *sync.Cond
	request  URLRequest
	response http.Response
	err      error

	access     sync.Mutex
	read       chan int
	readBuffer Buffer
	cancel     chan struct{}
	done       chan struct{}
}

func (r *urlResponse) monitorContext(ctx context.Context) {
	if ctx.Done() == nil {
		return
	}
	select {
	case <-r.cancel:
	case <-r.done:
	case <-ctx.Done():
		r.err = ctx.Err()
		r.Close()
	}
}

func (r *urlResponse) Read(p []byte) (n int, err error) {
	select {
	case <-r.done:
		return 0, r.err
	default:
	}

	r.access.Lock()

	select {
	case <-r.done:
		r.access.Unlock()
		return 0, r.err
	default:
	}

	r.readBuffer = NewBuffer()
	r.readBuffer.InitWithDataAndCallback(p, NewBufferCallback(nil))
	r.request.Read(r.readBuffer)
	r.access.Unlock()

	select {
	case bytesRead := <-r.read:
		return bytesRead, nil
	case <-r.cancel:
		return 0, net.ErrClosed
	case <-r.done:
		return 0, r.err
	}
}

func (r *urlResponse) Close() error {
	r.access.Lock()
	defer r.access.Unlock()
	select {
	case <-r.cancel:
		return os.ErrClosed
	case <-r.done:
		return nil
	default:
		close(r.cancel)
		r.request.Cancel()
	}
	return nil
}

// Cronet automatically decompresses body content if one of these encodings is used
var cronetEncodings = []string{"br", "deflate", "gzip", "x-gzip", "zstd"}


func (r *urlResponse) OnRedirectReceived(self URLRequestCallback, request URLRequest, info URLResponseInfo, newLocationUrl string) {
	if r.FollowRedirect {
		request.FollowRedirect()
		return
	}
	// No need to let cronet follow further redirect after first HTTP response
	r.response.Status = info.StatusText()
	r.response.StatusCode = info.StatusCode()
	headerLen := info.HeaderSize()
	for i := 0; i < headerLen; i++ {
		header := info.HeaderAt(i)
		r.response.Header.Set(header.Name(), header.Value())
	}
	r.response.Body = io.NopCloser(io.MultiReader())
	request.Cancel()
	r.complete.Signal()
}

func (r *urlResponse) OnResponseStarted(self URLRequestCallback, request URLRequest, info URLResponseInfo) {
	r.response.Status = info.StatusText()
	r.response.StatusCode = info.StatusCode()
	headerLen := info.HeaderSize()

	resetContentLength := false
	for i := 0; i < headerLen; i++ {
		header := info.HeaderAt(i)
		// Drop Content-Encoding header if body has been decompressed already
		// and reset Content-Length to unknown after loop completes
		if textproto.CanonicalMIMEHeaderKey(header.Name()) == "Content-Encoding" &&
			slices.Contains(cronetEncodings, strings.ToLower(header.Value())) {
			resetContentLength = true
			continue
		}
		r.response.Header.Set(header.Name(), header.Value())
	}
	if resetContentLength {
		r.response.Uncompressed = true
		r.response.ContentLength = -1
		r.response.Header.Del("Content-Length")
	} else {
		r.response.ContentLength, _ = strconv.ParseInt(r.response.Header.Get("Content-Length"), 10, 64)
	}
	r.response.TransferEncoding = r.response.Header.Values("Content-Transfer-Encoding")
	r.response.Close = true
	r.complete.Signal()
}

func (r *urlResponse) OnReadCompleted(self URLRequestCallback, request URLRequest, info URLResponseInfo, buffer Buffer, bytesRead int64) {
	r.access.Lock()
	defer r.access.Unlock()

	if bytesRead == 0 {
		r.close(request, io.EOF)
		return
	}

	select {
	case <-r.cancel:
	case <-r.done:
	case r.read <- int(bytesRead):
		r.readBuffer.Destroy()
		r.readBuffer = Buffer{}
	}
}

func (r *urlResponse) OnSucceeded(self URLRequestCallback, request URLRequest, info URLResponseInfo) {
	r.close(request, io.EOF)
}

func (r *urlResponse) OnFailed(self URLRequestCallback, request URLRequest, info URLResponseInfo, error Error) {
	r.close(request, ErrorFromError(error))
}

func (r *urlResponse) OnCanceled(self URLRequestCallback, request URLRequest, info URLResponseInfo) {
	r.close(request, context.Canceled)
}

func (r *urlResponse) close(request URLRequest, err error) {
	r.access.Lock()
	defer r.access.Unlock()

	select {
	case <-r.done:
		return
	default:
	}

	if r.err == nil {
		r.err = err
	}

	close(r.done)
	r.complete.Signal()
	request.Destroy()
}

type bodyUploadProvider struct {
	body          io.ReadCloser
	getBody       func() (io.ReadCloser, error)
	contentLength int64
}

func (p *bodyUploadProvider) Length(self UploadDataProvider) int64 {
	return p.contentLength
}

func (p *bodyUploadProvider) Read(self UploadDataProvider, sink UploadDataSink, buffer Buffer) {
	n, err := p.body.Read(buffer.DataSlice())
	if err != nil {
		if err == io.EOF {
			if p.contentLength == -1 {
				// Case for chunked uploads
				sink.OnReadSucceeded(0, true)
			} else {
				sink.OnReadSucceeded(int64(n), false)
			}
			return
		}
		sink.OnReadError(err.Error())
	} else {
		sink.OnReadSucceeded(int64(n), false)
	}
}

func (p *bodyUploadProvider) Rewind(self UploadDataProvider, sink UploadDataSink) {
	if p.getBody == nil {
		sink.OnRewindError("unsupported")
		return
	}
	p.body.Close()
	newBody, err := p.getBody()
	if err != nil {
		sink.OnRewindError(err.Error())
		return
	}
	p.body = newBody
	sink.OnRewindSucceeded()
}

func (p *bodyUploadProvider) Close(self UploadDataProvider) {
	self.Destroy()
	p.body.Close()
}
