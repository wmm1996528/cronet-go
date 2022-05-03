package cronet

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -lcronet.100.0.4896.60
// #include <stdlib.h>
// #include <stdbool.h>
// #include <cronet_c.h>
import "C"

import (
	"unsafe"
)

type PublicKeyPins struct {
	ptr C.Cronet_PublicKeyPinsPtr
}

func NewPublicKeyPins() *PublicKeyPins {
	return &PublicKeyPins{C.Cronet_PublicKeyPins_Create()}
}

func (p *PublicKeyPins) Destroy() {
	C.Cronet_PublicKeyPins_Destroy(p.ptr)
}

func (p *PublicKeyPins) Host() string {
	return C.GoString(C.Cronet_PublicKeyPins_host_get(p.ptr))
}

func (p *PublicKeyPins) SetHost(host string) {
	cHost := C.CString(host)
	C.Cronet_PublicKeyPins_host_set(p.ptr, cHost)
	C.free(unsafe.Pointer(cHost))
}

func (p *PublicKeyPins) AddPinnedSHa256(hash string) {
	cHash := C.CString(hash)
	C.Cronet_PublicKeyPins_pins_sha256_add(p.ptr, cHash)
	C.free(unsafe.Pointer(cHash))
}

func (p *PublicKeyPins) PinnedSha256Len() int {
	return int(C.Cronet_PublicKeyPins_pins_sha256_size(p.ptr))
}

func (p *PublicKeyPins) PinnedSHA256At(index int) string {
	return C.GoString(C.Cronet_PublicKeyPins_pins_sha256_at(p.ptr, C.uint32_t(index)))
}

func (p *PublicKeyPins) ClearPinnedSHA256() {
	C.Cronet_PublicKeyPins_pins_sha256_clear(p.ptr)
}

func (p *PublicKeyPins) IncludeSubdomains() bool {
	return bool(C.Cronet_PublicKeyPins_include_subdomains_get(p.ptr))
}

func (p *PublicKeyPins) SetIncludeSubdomains(includeSubdomains bool) {
	C.Cronet_PublicKeyPins_include_subdomains_set(p.ptr, C.bool(includeSubdomains))
}

func (p *PublicKeyPins) ExpirationDate() int64 {
	return int64(C.Cronet_PublicKeyPins_expiration_date_get(p.ptr))
}

func (p *PublicKeyPins) SetExpirationDate(date int64) {
	C.Cronet_PublicKeyPins_expiration_date_set(p.ptr, C.int64_t(date))
}