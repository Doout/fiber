//go:build s390x || ppc64 || ppc64le

package fiber

import "unsafe"

// isReadOnly reports whether the pointer references memory that belongs to
// Go's read-only data section. On architectures that enforce strict alignment
// (such as s390x and ppc64), converting the runtime section symbols to
// pointers can trigger alignment faults, so we conservatively return false.
func isReadOnly(unsafe.Pointer) bool {
	return false
}
