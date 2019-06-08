package window

import (
	"unsafe"

	"github.com/akiyosi/w32"
	"github.com/therecipe/qt/core"
)

// SetupNativeEvent overrides native events.
func (f *QFramelessWindow) SetupNativeEvent() {
	f.WinId()
	f.ConnectNativeEvent(func(eventType *core.QByteArray, message unsafe.Pointer, result *int) bool {
		msg := (*w32.MSG)(message)
		hwnd := msg.Hwnd

		switch msg.Message {
		case w32.WM_NCCALCSIZE:
			if msg.WParam == 1 {
				*result = 0
				return true
			}
			*result = (int)(w32.DefWindowProc(msg.Hwnd, w32.WM_NCCALCSIZE, msg.WParam, msg.LParam))
			return true

		case w32.WM_GETMINMAXINFO:
			mm := (*w32.MINMAXINFO)((unsafe.Pointer)(msg.LParam))
			mm.PtMinTrackSize.X = int32(f.minimumWidth)
			mm.PtMinTrackSize.Y = int32(f.minimumHeight)

			return true

		case w32.WM_ACTIVATEAPP:
			f.putShadow(hwnd)

		}
		return false
	})
}

// SetupNativeEvent2 overrides native events.
func (f *QFramelessWindow) SetupNativeEvent2() {
	filterObj := core.NewQAbstractNativeEventFilter()
	filterObj.ConnectNativeEventFilter(func(eventType *core.QByteArray, message unsafe.Pointer, result *int) bool {
		msg := (*w32.MSG)(message)
		hwnd := msg.Hwnd

		switch msg.Message {
		case w32.WM_CREATE:
			style := w32.GetWindowLong(hwnd, w32.GWL_STYLE)
			style = style | w32.WS_THICKFRAME
			w32.SetWindowLong(hwnd, w32.GWL_STYLE, uint32(style))

		case w32.WM_NCCALCSIZE:
			if msg.WParam == 1 {
				*result = 0
				return true
			}
			return false

		case w32.WM_GETMINMAXINFO:
			mm := (*w32.MINMAXINFO)((unsafe.Pointer)(msg.LParam))
			mm.PtMinTrackSize.X = int32(f.minimumWidth)
			mm.PtMinTrackSize.Y = int32(f.minimumHeight)

			return true
		}

		return false
	})
	core.QCoreApplication_Instance().InstallNativeEventFilter(filterObj)
}

func (f *QFramelessWindow) putShadow(hwnd w32.HWND) {
	if f.borderless {
		return
	}

	style := w32.GetWindowLong(hwnd, w32.GWL_STYLE)
	style = style | w32.WS_THICKFRAME ^ w32.WS_CAPTION
	w32.SetWindowLong(hwnd, w32.GWL_STYLE, uint32(style))

	// Set shadow.
	shadow := &w32.MARGINS{1, 1, 1, 1}
	w32.DwmExtendFrameIntoClientArea(hwnd, shadow)

	var uflag uint
	uflag = w32.SWP_NOZORDER | w32.SWP_NOOWNERZORDER | w32.SWP_NOMOVE | w32.SWP_NOSIZE | w32.SWP_FRAMECHANGED
	var nullptr w32.HWND
	w32.SetWindowPos(hwnd, nullptr, 0, 0, 0, 0, uflag)

	w32.UpdateWindow(hwnd)

	f.borderless = true
}
