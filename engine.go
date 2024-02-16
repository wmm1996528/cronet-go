package cronet

// #include <stdlib.h>
// #include <stdbool.h>
// #include <cronet_c.h>
import "C"

import (
	"unsafe"
)

// Engine is an engine to process URLRequest, which uses the best HTTP stack
// available on the current platform. An instance of this class can be started
// using StartWithParams.
type Engine struct {
	ptr C.Cronet_EnginePtr
}

func NewEngine() Engine {
	return Engine{C.Cronet_Engine_Create()}
}

func (e Engine) Destroy() {
	C.Cronet_Engine_Destroy(e.ptr)
}

// StartWithParams starts Engine using given |params|. The engine must be started once
// and only once before other methods can be used.
func (e Engine) StartWithParams(params EngineParams) Result {
	return Result(C.Cronet_Engine_StartWithParams(e.ptr, params.ptr))
}

// StartNetLogToFile starts NetLog logging to a file. The NetLog will contain events emitted
// by all live Engines. The NetLog is useful for debugging.
// The file can be viewed using a Chrome browser navigated to
// chrome://net-internals/#import
// Returns |true| if netlog has started successfully, |false| otherwise.
// Parameter |fileName| the complete file path. It must not be empty. If the file
// exists, it is truncated before starting. If actively logging, this method is ignored.
// Parameter |logAll| to include basic events, user cookies, credentials and all transferred
// bytes in the log. This option presentsa privacy risk, since it exposes the user's credentials,
// and should only be used with the user's consent and in situations where the log won't be public.
// false to just include basic events.
func (e Engine) StartNetLogToFile(fileName string, logAll bool) bool {
	cPath := C.CString(fileName)
	result := C.Cronet_Engine_StartNetLogToFile(e.ptr, cPath, C.bool(logAll))
	C.free(unsafe.Pointer(cPath))
	return bool(result)
}

// StopNetLog Stops NetLog logging and flushes file to disk. If a logging session is
// not in progress, this call is ignored. This method blocks until the log is
// closed to ensure that log file is complete and available.
func (e Engine) StopNetLog() {
	C.Cronet_Engine_StopNetLog(e.ptr)
}

// Shutdown shuts down the Engine if there are no active requests,
// otherwise returns a failure Result.
//
// Cannot be called on network thread - the thread Cronet calls into
// Executor on (which is different from the thread the Executor invokes
// callbacks on). This method blocks until all the Engine's resources have
// been cleaned up.
func (e Engine) Shutdown() Result {
	return Result(C.Cronet_Engine_Shutdown(e.ptr))
}

// Version returns a human-readable version string of the engine.
func (e Engine) Version() string {
	return C.GoString(C.Cronet_Engine_GetVersionString(e.ptr))
}

// DefaultUserAgent Returns default human-readable version string of the engine. Can be used
// before StartWithParams() is called.
func (e Engine) DefaultUserAgent() string {
	return C.GoString(C.Cronet_Engine_GetDefaultUserAgent(e.ptr))
}

// AddRequestFinishedListener registers a listener that gets called at the end of each request.
//
// The listener is called on Executor.
//
// The listener is called before URLRequestCallbackHandler.OnCanceled(),
// URLRequestCallbackHandler.OnFailed() or
// URLRequestCallbackHandler.OnSucceeded() is called -- note that if Executor
// runs the listener asynchronously, the actual call to the listener
// may happen after a URLRequestCallbackHandler method is called.
//
// Listeners are only guaranteed to be called for requests that are started
// after the listener is added.
//
// Ownership is **not** taken for listener or Executor.
//
// Assuming the listener won't run again (there are no pending requests with
// the listener attached, either via Engine or UrlRequest),
// the app may destroy it once its OnRequestFinished() has started,
// even inside that method.
//
// Similarly, the app may destroy executor in or after OnRequestFinished()}.
//
// It's also OK to destroy executor in or after one of
// URLRequestCallbackHandler.OnCanceled(), URLRequestCallbackHandler.OnFailed() or
// URLRequestCallbackHandler.OnSucceeded().
//
// Of course, both of these are only true if listener won't run again
// and executor isn't being used for anything else that might start
// running in the future.
//
// @param listener the listener for finished requests.
// @param executor the executor upon which to run listener.
func (e Engine) AddRequestFinishedListener(listener URLRequestFinishedInfoListener, executor Executor) {
	C.Cronet_Engine_AddRequestFinishedListener(e.ptr, listener.ptr, executor.ptr)
}

// RemoveRequestFinishedListener unregisters a RequestFinishedInfoListener,
// including its association with its registered Executor.
func (e Engine) RemoveRequestFinishedListener(listener URLRequestFinishedInfoListener) {
	C.Cronet_Engine_RemoveRequestFinishedListener(e.ptr, listener.ptr)
}

// SetClientCertificate Configures all subsequent connections to server designated with {@code hostPortPair}
// to authenticate with {@code client_cert_data} and {@code private_key_data} when requested.
// {@code clientCertData} can be PEM/DER encoded.
// {@code privateKeyData} is supposed to be PEM encoded.
//
// The method can be called only after the engine is started.
func (e Engine) SetClientCertificate(hostPortPair string, clientCertData []byte, privateKeyData []byte) {
	cHostPortPair := C.CString(hostPortPair)
	clientCertBuffer := NewBuffer()
	clientCertBuffer.InitWithDataAndCallback(clientCertData, NewBufferCallback(nil))
	privateKeyBuffer := NewBuffer()
	privateKeyBuffer.InitWithDataAndCallback(privateKeyData, NewBufferCallback(nil))
	C.Cronet_Engine_SetClientCertificate(e.ptr, cHostPortPair, clientCertBuffer.ptr, privateKeyBuffer.ptr)
	clientCertBuffer.Destroy()
	privateKeyBuffer.Destroy()
	C.free(unsafe.Pointer(cHostPortPair))
}

// ClearClientCertificate Clears a client certificate preference for server designated with {@code hostPortPair}
// set by SetClientCertificate(). Returns true if one was removed and false otherwise.
//
// The method can be called only after the engine is started.
func (e Engine) ClearClientCertificate(hostPortPair string) Result {
	cHostPortPair := C.CString(hostPortPair)
	result := C.Cronet_Engine_ClearClientCertificate(e.ptr, cHostPortPair)
	C.free(unsafe.Pointer(cHostPortPair))
	return Result(result)
}
