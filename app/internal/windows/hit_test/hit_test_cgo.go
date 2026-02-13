//go:build windows && cgo
// +build windows,cgo

package hit_test

import "golang.org/x/sys/windows"

/*
#include <windows.h>
#include <stdint.h>

extern int32_t go_ElementProviderFromPoint(uintptr_t pThis, double x, double y, uintptr_t* retVal);

HRESULT STDMETHODCALLTYPE ElementProviderFromPoint(uintptr_t pThis, double x, double y, uintptr_t* retVal) {
    return (HRESULT)go_ElementProviderFromPoint(pThis, x, y, retVal);
}
*/
import "C"

var ElementProviderFromPoint = windows.NewCallbackCDecl(C.ElementProviderFromPoint)

// FIXME: wait for https://github.com/golang/go/issues/45300 to be fixed, but
// for the time being, provide dummy implementation.
func go_ElementProviderFromPoint(pThis uintptr, x, y uintptr, retVal *uintptr) uintptr {
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
}
