package window

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>

void setStyleMask(long *wid) {
    NSView* view = (NSView*)wid;
    NSWindow *window = view.window;

    window.styleMask |= NSWindowStyleMaskResizable;
    window.styleMask |= NSWindowStyleMaskMiniaturizable;
    window.styleMask |= NSWindowStyleMaskFullSizeContentView;
    window.opaque = NO;
    window.backgroundColor = [NSColor clearColor];
    window.hasShadow = NO;

    return;
}

void setStyleMaskAndShadow(long *wid) {
    NSView* view = (NSView*)wid;
    NSWindow *window = view.window;

    window.styleMask |= NSWindowStyleMaskResizable;
    window.styleMask |= NSWindowStyleMaskMiniaturizable;
    window.styleMask |= NSWindowStyleMaskFullSizeContentView;
    window.opaque = NO;
    window.backgroundColor = [NSColor clearColor];
    window.hasShadow = YES;

    return;
}
*/
import "C"

// SetStyleMask ...
func (f *QFramelessWindow) SetStyleMask() {}

// SetupNativeEvent ...
func (f *QFramelessWindow) SetupNativeEvent() {}

// SetupNativeEvent2 ...
func (f *QFramelessWindow) SetupNativeEvent2() {}
