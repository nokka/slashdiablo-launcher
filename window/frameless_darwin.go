package window

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>
*/
import "C"

// SetupNativeEvent overrides native events.
func (f *QFramelessWindow) SetupNativeEvent() {}

// SetupNativeEvent2 overrides native events.
func (f *QFramelessWindow) SetupNativeEvent2() {}
