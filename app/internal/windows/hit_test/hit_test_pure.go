//go:build windows && !cgo
// +build windows,!cgo

package hit_test

import "golang.org/x/sys/windows"

var ElementProviderFromPoint = windows.NewCallback(_ElementProviderFromPoint)

// FIXME: wait for https://github.com/golang/go/issues/45300 to be fixed, but
// for the time being, provide dummy implementation.
func _ElementProviderFromPoint(pThis uintptr, x, y uintptr, retVal *uintptr) uintptr {
	// this := FromFragmentRootPointer(pThis)
	//
	// if retVal == nil {
	//      return windows.E_POINTER
	// }
	//
	// p := f32.Pt(float32(x), float32(y))
	// id, exists := this.w.SemanticAt(p)
	// if !exists {
	//      *retVal = 0
	//      return windows.S_OK
	// }

	return 0
}
