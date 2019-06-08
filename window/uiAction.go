// +build !darwin

package window

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
)

func (f *QFramelessWindow) SetupTitleBarActions() {
	t := f.TitleBar

	f.IconMinimize.Widget.ConnectEnterEvent(func(event *core.QEvent) {
		f.IconMinimize.SetStyle(&RGB{
			R: 0,
			G: 162,
			B: 232,
		})
	})
	f.IconMaximize.Widget.ConnectEnterEvent(func(event *core.QEvent) {
		f.IconMaximize.SetStyle(&RGB{
			R: 0,
			G: 162,
			B: 232,
		})
	})
	f.IconRestore.Widget.ConnectEnterEvent(func(event *core.QEvent) {
		f.IconRestore.SetStyle(&RGB{
			R: 0,
			G: 162,
			B: 232,
		})
	})
	f.IconClose.Widget.ConnectEnterEvent(func(event *core.QEvent) {
		f.IconClose.SetStyle(&RGB{
			R: 0,
			G: 162,
			B: 232,
		})
	})

	f.IconMinimize.Widget.ConnectLeaveEvent(func(event *core.QEvent) {
		f.IconMinimize.SetStyle(nil)
	})
	f.IconMaximize.Widget.ConnectLeaveEvent(func(event *core.QEvent) {
		f.IconMaximize.SetStyle(nil)
	})
	f.IconRestore.Widget.ConnectLeaveEvent(func(event *core.QEvent) {
		f.IconRestore.SetStyle(nil)
	})
	f.IconClose.Widget.ConnectLeaveEvent(func(event *core.QEvent) {
		f.IconClose.SetStyle(nil)
	})

	// Button Actions
	f.IconMinimize.Widget.ConnectMousePressEvent(func(e *gui.QMouseEvent) {
		f.IsTitleBarPressed = false
	})

	f.IconMaximize.Widget.ConnectMousePressEvent(func(e *gui.QMouseEvent) {
		f.IsTitleBarPressed = false
	})

	f.IconRestore.Widget.ConnectMousePressEvent(func(e *gui.QMouseEvent) {
		f.IsTitleBarPressed = false
	})

	f.IconClose.Widget.ConnectMousePressEvent(func(e *gui.QMouseEvent) {
		f.IsTitleBarPressed = false
	})

	f.IconMinimize.Widget.ConnectMouseReleaseEvent(func(e *gui.QMouseEvent) {
		isContain := f.IconMinimize.Widget.Rect().Contains(e.Pos(), false)
		if !isContain {
			return
		}
		f.SetWindowState(core.Qt__WindowMinimized)
		f.Widget.Hide()
		f.Widget.Show()
	})

	f.IconMaximize.Widget.ConnectMouseReleaseEvent(func(e *gui.QMouseEvent) {
		isContain := f.IconMinimize.Widget.Rect().Contains(e.Pos(), false)
		if !isContain {
			return
		}
		f.windowMaximize()
		f.Widget.Hide()
		f.Widget.Show()
	})

	f.IconRestore.Widget.ConnectMouseReleaseEvent(func(e *gui.QMouseEvent) {
		isContain := f.IconMinimize.Widget.Rect().Contains(e.Pos(), false)
		if !isContain {
			return
		}
		f.windowRestore()
		f.Widget.Hide()
		f.Widget.Show()
	})

	f.IconClose.Widget.ConnectMouseReleaseEvent(func(e *gui.QMouseEvent) {
		isContain := f.IconMinimize.Widget.Rect().Contains(e.Pos(), false)
		if !isContain {
			return
		}
		f.Close()
	})

	// TitleBar Actions
	t.ConnectMousePressEvent(func(e *gui.QMouseEvent) {
		f.Widget.Raise()
		f.IsTitleBarPressed = true
		f.TitleBarMousePos = e.GlobalPos()
		f.Position = f.Pos()
	})

	t.ConnectMouseReleaseEvent(func(e *gui.QMouseEvent) {
		f.IsTitleBarPressed = false
	})

	t.ConnectMouseMoveEvent(func(e *gui.QMouseEvent) {
		if !f.IsTitleBarPressed {
			return
		}
		x := f.Position.X() + e.GlobalPos().X() - f.TitleBarMousePos.X()
		y := f.Position.Y() + e.GlobalPos().Y() - f.TitleBarMousePos.Y()
		newPos := core.NewQPoint2(x, y)
		f.Move(newPos)
	})

	t.ConnectMouseDoubleClickEvent(func(e *gui.QMouseEvent) {
		if f.IconMaximize.Widget.IsVisible() {
			f.windowMaximize()
		} else {
			f.windowRestore()
		}
	})
}

func (f *QFramelessWindow) windowMaximize() {
	f.IconMaximize.Widget.SetVisible(false)
	f.IconRestore.Widget.SetVisible(true)
	f.Layout.SetContentsMargins(0, 0, 0, 0)
	f.SetWindowState(core.Qt__WindowMaximized)
	f.IconRestore.SetStyle(nil)
}

func (f *QFramelessWindow) windowRestore() {
	f.IconMaximize.Widget.SetVisible(true)
	f.IconRestore.Widget.SetVisible(false)
	f.Layout.SetContentsMargins(f.shadowMargin, f.shadowMargin, f.shadowMargin, f.shadowMargin)
	f.SetWindowState(core.Qt__WindowNoState)
}
