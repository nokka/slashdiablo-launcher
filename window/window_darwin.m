#import <AppKit/AppKit.h>

void allowMinimize(void* w)
{
	[[(NSView*)(w) window] setStyleMask:NSWindowStyleMaskMiniaturizable];
}