// +build darwin

package window

//void allowMinimize(void*);
import "C"
import "unsafe"

// AllowMinimize will make the window minimizable on OSX.
func AllowMinimize(winID uintptr) {
	C.allowMinimize(unsafe.Pointer(winID))
}
