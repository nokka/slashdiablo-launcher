package window

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>
*/
import "C"

// SetStyleMask ...
func (f *QFramelessWindow) SetStyleMask() {}

// SetupNativeEvent ...
func (f *QFramelessWindow) SetupNativeEvent() {}

// SetupNativeEvent2 ...
func (f *QFramelessWindow) SetupNativeEvent2() {}
