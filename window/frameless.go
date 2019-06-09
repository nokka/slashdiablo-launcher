package window

import (
	"fmt"
	"math"
	"runtime"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/svg"
	"github.com/therecipe/qt/widgets"
)

// Edge represents a position.
type Edge int

// Edge positions.
const (
	None        Edge = 0x0
	Left        Edge = 0x1
	Top         Edge = 0x2
	Right       Edge = 0x4
	Bottom      Edge = 0x8
	TopLeft     Edge = 0x10
	TopRight    Edge = 0x20
	BottomLeft  Edge = 0x40
	BottomRight Edge = 0x80
)

// RGB represents a color.
type RGB struct {
	R uint16
	G uint16
	B uint16
}

// QToolButtonForNotDarwin toolbar for everything but darwin.
type QToolButtonForNotDarwin struct {
	f       *QFramelessWindow
	Widget  *widgets.QWidget
	IconBtn *svg.QSvgWidget
	isHover bool
}

// QFramelessWindow is the main frameless window.
type QFramelessWindow struct {
	widgets.QMainWindow
	WindowColor       *RGB
	WindowColorAlpha  float64
	Widget            *widgets.QWidget
	Layout            *widgets.QVBoxLayout
	Content           *widgets.QWidget
	WindowWidget      *widgets.QFrame
	WindowVLayout     *widgets.QVBoxLayout
	shadowMargin      int
	borderSize        int
	minimumWidth      int
	minimumHeight     int
	TitleBar          *widgets.QWidget
	TitleBarLayout    *widgets.QHBoxLayout
	TitleLabel        *widgets.QLabel
	TitleBarBtnWidget *widgets.QWidget
	TitleBarBtnLayout *widgets.QHBoxLayout
	TitleColor        *RGB
	TitleBarMousePos  *core.QPoint
	IsTitleBarPressed bool

	// For darwin.
	BtnMinimize *widgets.QToolButton
	BtnRestore  *widgets.QToolButton
	BtnClose    *widgets.QToolButton

	// For Windows and Linux.
	IconMinimize *QToolButtonForNotDarwin
	IconRestore  *QToolButtonForNotDarwin
	IconClose    *QToolButtonForNotDarwin

	isCursorChanged     bool
	isDragStart         bool
	isLeftButtonPressed bool
	dragPos             *core.QPoint
	hoverEdge           Edge
	Position            *core.QPoint
	MousePos            [2]int
	borderless          bool
}

// NewFramelessWindow creates a new frameless window.
func NewFramelessWindow(alpha float64, width int, height int) *QFramelessWindow {
	f := NewQFramelessWindow(nil, 0)
	f.WindowColorAlpha = alpha

	f.Widget = widgets.NewQWidget(nil, 0)
	f.SetCentralWidget(f.Widget)

	f.shadowMargin = 0
	f.borderSize = 1
	
	f.SetupUI(f.Widget)
	f.SetupWindowFlags()

	// DO NOT NEED.
	//f.SetupAttributes()

	// NEED TO HAVE.
	f.SetupWindowActions()

	// NEED TO HAVE.
	f.SetupTitleBarActions()

	f.SetFixedSize2(width, height)
	f.SetupWidgetColor(0, 0, 0)

	return f
}

// SetupWindowShadow ...
func (f *QFramelessWindow) SetupWindowShadow(size int) {
	f.shadowMargin = size
	f.Layout.SetContentsMargins(f.shadowMargin, f.shadowMargin, f.shadowMargin, f.shadowMargin)
	if f.shadowMargin == 0 {
		return
	}
	shadow := widgets.NewQGraphicsDropShadowEffect(nil)
	var alpha int
	if runtime.GOOS == "darwin" {
		alpha = 220
		shadow.SetOffset3(0, 10)
	} else {
		alpha = 100
		shadow.SetOffset3(0, 0)
	}
	shadow.SetBlurRadius((float64)(f.shadowMargin))
	shadow.SetColor(gui.NewQColor3(0, 0, 0, alpha))
	f.WindowWidget.SetGraphicsEffect(shadow)
}

// SetupUI ...
func (f *QFramelessWindow) SetupUI(widget *widgets.QWidget) {
	f.InstallEventFilter(f)

	widget.SetSizePolicy2(widgets.QSizePolicy__Expanding|widgets.QSizePolicy__Maximum, widgets.QSizePolicy__Expanding|widgets.QSizePolicy__Maximum)
	f.Layout = widgets.NewQVBoxLayout2(widget)
	f.Layout.SetSpacing(0)

	f.WindowWidget = widgets.NewQFrame(widget, 0)
	f.WindowWidget.SetObjectName("QFramelessWidget")
	f.WindowWidget.SetSizePolicy2(widgets.QSizePolicy__Expanding|widgets.QSizePolicy__Maximum, widgets.QSizePolicy__Expanding|widgets.QSizePolicy__Maximum)

	f.Layout.SetContentsMargins(f.shadowMargin, f.shadowMargin, f.shadowMargin, f.shadowMargin)

	f.WindowVLayout = widgets.NewQVBoxLayout2(f.WindowWidget)
	f.WindowVLayout.SetContentsMargins(f.borderSize, f.borderSize, f.borderSize, 0)
	f.WindowVLayout.SetContentsMargins(0, 0, 0, 0)
	f.WindowVLayout.SetSpacing(0)
	f.WindowWidget.SetLayout(f.WindowVLayout)

	// create titlebar widget
	f.TitleBar = widgets.NewQWidget(f.WindowWidget, 0)
	f.TitleBar.SetObjectName("titleBar")
	f.TitleBar.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Fixed)
	f.TitleBarLayout = widgets.NewQHBoxLayout2(f.TitleBar)
	f.TitleBarLayout.SetContentsMargins(0, 0, 0, 0)

	f.TitleLabel = widgets.NewQLabel(nil, 0)
	f.TitleLabel.SetObjectName("TitleLabel")
	f.TitleLabel.SetAlignment(core.Qt__AlignCenter)

	if runtime.GOOS == "darwin" {
		f.SetTitleBarButtonsForDarwin()
	} else {
		f.SetTitleBarButtons()
	}

	// create window content
	f.Content = widgets.NewQWidget(f.WindowWidget, 0)

	// Set widget to layout
	f.WindowVLayout.AddWidget(f.TitleBar, 0, 0)
	f.WindowVLayout.AddWidget(f.Content, 0, 0)

	f.Layout.AddWidget(f.WindowWidget, 0, 0)
}

// SetupWidgetColor ...
func (f *QFramelessWindow) SetupWidgetColor(red uint16, green uint16, blue uint16) {
	alpha := f.WindowColorAlpha
	f.WindowColor = &RGB{
		R: red,
		G: green,
		B: blue,
	}
	color := f.WindowColor
	style := fmt.Sprintf("background-color: rgba(%d, %d, %d, %f);", color.R, color.G, color.B, alpha)
	f.Widget.SetStyleSheet(" * { background-color: rgba(0, 0, 0, 0.0); color: rgba(0, 0, 0, 0); }")

	borderSizeString := fmt.Sprintf("%d", f.borderSize*2) + "px"

	var roundSizeString string
	if runtime.GOOS != "windows" {
		roundSizeString = "10px" //fmt.Sprintf("%d", f.borderSize*2) + "px"
	} else {
		roundSizeString = "0px"
	}

	f.WindowWidget.SetStyleSheet(fmt.Sprintf(`
	#QFramelessWidget {
		border: 0px solid %s; 
		padding-top: 2px; padding-right: %s; padding-bottom: %s; padding-left: %s; 
		border-radius: %s;
		%s; 
	}`, color.Hex(), borderSizeString, borderSizeString, borderSizeString, roundSizeString, style))
}

// NewQToolButtonForNotDarwin ...
func NewQToolButtonForNotDarwin(parent widgets.QWidget_ITF) *QToolButtonForNotDarwin {
	iconSize := 15
	marginTB := iconSize / 6
	marginLR := 1
	if runtime.GOOS == "linux" {
		iconSize = 18
		marginLR = int(float64(iconSize) / float64(3.5))
	} else {
		marginLR = int(float64(iconSize) / float64(2.5))
	}

	widget := widgets.NewQWidget(parent, 0)
	widget.SetSizePolicy2(widgets.QSizePolicy__Fixed, widgets.QSizePolicy__Fixed)
	layout := widgets.NewQVBoxLayout2(widget)
	layout.SetContentsMargins(marginLR, marginTB, marginLR, marginTB)
	icon := svg.NewQSvgWidget(nil)
	icon.SetFixedSize2(iconSize, iconSize)

	layout.AddWidget(icon, 0, 0)
	layout.SetAlignment(icon, core.Qt__AlignCenter)

	return &QToolButtonForNotDarwin{
		Widget:  widget,
		IconBtn: icon,
	}
}

// SetObjectName ...
func (b *QToolButtonForNotDarwin) SetObjectName(name string) {
	b.IconBtn.SetObjectName(name)
}

// Hide ...
func (b *QToolButtonForNotDarwin) Hide() {
	b.Widget.Hide()
}

// Show ...
func (b *QToolButtonForNotDarwin) Show() {
	b.Widget.Show()
}

// SetStyle ...
func (b *QToolButtonForNotDarwin) SetStyle(color *RGB) {
	var backgroundColor string
	if color == nil {
		backgroundColor = "background-color:none;"
	} else {
		hoverColor := color.Brend(b.f.WindowColor, 0.75)
		backgroundColor = fmt.Sprintf("background-color: rgba(%d, %d, %d, %f);", hoverColor.R, hoverColor.G, hoverColor.B, b.f.WindowColorAlpha)
	}

	b.Widget.SetStyleSheet(fmt.Sprintf(`
	.QWidget { 
		%s;
		border:none;
	}
	`, backgroundColor))
}

func (f *QFramelessWindow) SetTitleBarButtons() {
	iconSize := 15
	f.TitleBarLayout.SetSpacing(1)

	f.IconMinimize = NewQToolButtonForNotDarwin(nil)
	f.IconMinimize.f = f
	f.IconMinimize.IconBtn.SetFixedSize2(iconSize, iconSize)
	f.IconMinimize.SetObjectName("IconMinimize")
	f.IconRestore = NewQToolButtonForNotDarwin(nil)
	f.IconRestore.f = f
	f.IconRestore.IconBtn.SetFixedSize2(iconSize, iconSize)
	f.IconRestore.SetObjectName("IconRestore")
	f.IconClose = NewQToolButtonForNotDarwin(nil)
	f.IconClose.f = f
	f.IconClose.IconBtn.SetFixedSize2(iconSize, iconSize)
	f.IconClose.SetObjectName("IconClose")

	f.SetIconsStyle(nil)

	f.IconMinimize.Hide()
	f.IconRestore.Hide()
	f.IconClose.Hide()

	f.TitleBarLayout.SetAlignment(f.TitleBarBtnWidget, core.Qt__AlignRight)
	f.TitleBarLayout.AddWidget(f.TitleLabel, 0, 0)
	f.TitleBarLayout.AddWidget(f.IconMinimize.Widget, 0, 0)
	f.TitleBarLayout.AddWidget(f.IconRestore.Widget, 0, 0)
	f.TitleBarLayout.AddWidget(f.IconClose.Widget, 0, 0)
}

// SetIconsStyle ...
func (f *QFramelessWindow) SetIconsStyle(color *RGB) {
	for _, b := range []*QToolButtonForNotDarwin{
		f.IconMinimize,
		f.IconRestore,
		f.IconClose,
	} {
		b.SetStyle(color)
	}
}

// SetTitleBarButtonsForDarwin ...
func (f *QFramelessWindow) SetTitleBarButtonsForDarwin() {
	btnSizePolicy := widgets.NewQSizePolicy2(widgets.QSizePolicy__Fixed, widgets.QSizePolicy__Fixed, widgets.QSizePolicy__ToolButton)
	f.BtnMinimize = widgets.NewQToolButton(f.TitleBar)
	f.BtnMinimize.SetObjectName("BtnMinimize")
	f.BtnMinimize.SetSizePolicy(btnSizePolicy)

	f.BtnRestore = widgets.NewQToolButton(f.TitleBar)
	f.BtnRestore.SetObjectName("BtnRestore")
	f.BtnRestore.SetSizePolicy(btnSizePolicy)
	f.BtnRestore.SetVisible(false)

	f.BtnClose = widgets.NewQToolButton(f.TitleBar)
	f.BtnClose.SetObjectName("BtnClose")
	f.BtnClose.SetSizePolicy(btnSizePolicy)

	f.TitleBarLayout.SetSpacing(0)
	f.TitleBarLayout.SetAlignment(f.TitleBarBtnWidget, core.Qt__AlignLeft)
	f.TitleBarLayout.AddWidget(f.BtnClose, 0, 0)
	f.TitleBarLayout.AddWidget(f.BtnMinimize, 0, 0)
	f.TitleBarLayout.AddWidget(f.BtnRestore, 0, 0)
	f.TitleBarLayout.AddWidget(f.TitleLabel, 0, 0)
}

// SetupAttributes ...
func (f *QFramelessWindow) SetupAttributes() {
	f.SetAttribute(core.Qt__WA_TranslucentBackground, true)
	f.SetAttribute(core.Qt__WA_NoSystemBackground, true)
	f.SetAttribute(core.Qt__WA_Hover, true)
	f.SetMouseTracking(true)
}

// SetupWindowFlags ...
func (f *QFramelessWindow) SetupWindowFlags() {
	f.SetWindowFlag(core.Qt__Window, true)
	f.SetWindowFlag(core.Qt__FramelessWindowHint, true)
	f.SetWindowFlag(core.Qt__NoDropShadowWindowHint, true)
	f.SetWindowFlag(core.Qt__MSWindowsFixedSizeDialogHint, true)
}

// SetupTitle ...
func (f *QFramelessWindow) SetupTitle(title string) {
	f.TitleLabel.SetText(title)
}

// SetupTitleColor ...
func (f *QFramelessWindow) SetupTitleColor(red uint16, green uint16, blue uint16) {
	f.TitleColor = &RGB{
		R: red,
		G: green,
		B: blue,
	}
	f.SetupTitleBarColor()
}

// SetupTitleBarColor ...
func (f *QFramelessWindow) SetupTitleBarColor() {
	var color, labelColor *RGB
	if f.IsActiveWindow() {
		color = f.TitleColor
	} else {
		color = nil
	}
	labelColor = color
	if labelColor == nil {
		labelColor = &RGB{
			R: 128,
			G: 128,
			B: 128,
		}
	}
	if runtime.GOOS != "darwin" {
		f.TitleLabel.SetStyleSheet(fmt.Sprintf(" *{padding-left: 60px; color: rgb(%d, %d, %d); }", labelColor.R, labelColor.G, labelColor.B))
		f.SetupTitleBarColorForNotDarwin(color)
	} else {
		f.TitleLabel.SetStyleSheet(fmt.Sprintf(" *{padding-right: 60px; color: rgb(%d, %d, %d); }", labelColor.R, labelColor.G, labelColor.B))
		f.SetupTitleBarColorForDarwin(color)
	}
}

func (c *RGB) fade() *RGB {
	r := (float64)(c.R)
	g := (float64)(c.G)
	b := (float64)(c.B)
	disp := (math.Abs(128-r) + math.Abs(128-g) + math.Abs(128-b)) / 3 * 1 / 4
	var newColor [3]float64
	for i, color := range []float64{
		r, g, b,
	} {
		if color > 128 {
			newColor[i] = color - disp
		} else {
			newColor[i] = color + disp
		}
		if newColor[i] < 0 {
			newColor[i] = 0
		} else if newColor[i] > 255 {
			newColor[i] = 255
		}
	}

	return &RGB{
		R: (uint16)(newColor[0]),
		G: (uint16)(newColor[1]),
		B: (uint16)(newColor[2]),
	}
}

// SetupTitleBarColorForNotDarwin ...
func (f *QFramelessWindow) SetupTitleBarColorForNotDarwin(color *RGB) {
	if color == nil {
		color = &RGB{
			R: 128,
			G: 128,
			B: 128,
		}
	} else {
		color = color.fade()
	}
	var SvgMinimize, SvgRestore, SvgClose string

	if runtime.GOOS == "windows" {
		SvgMinimize = fmt.Sprintf(`
		<svg style="width:24px;height:24px" viewBox="0 0 24 24">
		<path fill="%s" d="M20,14H4V10H20" />
		</svg>
		`, color.Hex())

		SvgRestore = fmt.Sprintf(`
		<svg style="width:24px;height:24px" viewBox="0 0 24 24">
		<path fill="%s" d="M4,8H8V4H20V16H16V20H4V8M16,8V14H18V6H10V8H16M6,12V18H14V12H6Z" />
		</svg>
		`, color.Hex())

		SvgClose = fmt.Sprintf(`
		<svg style="width:24px;height:24px" viewBox="0 0 24 24">
		<path fill="%s" d="M13.46,12L19,17.54V19H17.54L12,13.46L6.46,19H5V17.54L10.54,12L5,6.46V5H6.46L12,10.54L17.54,5H19V6.46L13.46,12Z" />
		</svg>
		`, color.Hex())
	} else {
		SvgMinimize = fmt.Sprintf(`
		<svg style="width:24px;height:24px" viewBox="0 0 24 24">
		<path fill="%s" d="M17,13H7V11H17M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2Z" />
		</svg>
		`, color.Hex())

		SvgRestore = fmt.Sprintf(`
		<svg style="width:24px;height:24px" viewBox="0 0 24 24">
		<path fill="%s" d="M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2M9,9H15V15H9" />
		</svg>
		`, color.Hex())

		SvgClose = fmt.Sprintf(`
		<svg style="width:24px;height:24px" viewBox="0 0 24 24">
		<g transform="translate(0,1)">
		<path fill="%s" d="M12 2C6.47 2 2 6.47 2 12s4.47 10 10 10 10-4.47 10-10S17.53 2 12 2zm5 13.59L15.59 17 12 13.41 8.41 17 7 15.59 10.59 12 7 8.41 8.41 7 12 10.59 15.59 7 17 8.41 13.41 12 17 15.59z"/><path d="M0 0h24v24H0z" fill="none"/></g></svg>
		`, "#e86032")
	}

	f.IconMinimize.IconBtn.Load2(core.NewQByteArray2(SvgMinimize, len(SvgMinimize)))
	f.IconRestore.IconBtn.Load2(core.NewQByteArray2(SvgRestore, len(SvgRestore)))
	f.IconClose.IconBtn.Load2(core.NewQByteArray2(SvgClose, len(SvgClose)))

	f.IconMinimize.Show()
	f.IconRestore.Show()
	f.IconRestore.Widget.SetVisible(false)
	f.IconClose.Show()
}

func (f *QFramelessWindow) SetupTitleBarColorForDarwin(color *RGB) {
	var baseStyle, restoreAndMaximizeColor, minimizeColor, closeColor string
	baseStyle = ` #BtnMinimize, #BtnMaximize, #BtnRestore, #BtnClose {
		min-width: 10px;
		min-height: 10px;
		max-width: 10px;
		max-height: 10px;
		border-radius: 6px;
		border-width: 1px;
		border-style: solid;
		margin: 4px;
	}`
	if color != nil {
		restoreAndMaximizeColor = `
			#BtnRestore, #BtnMaximize {
				background-color: rgb(53, 202, 74);
				border-color: rgb(34, 182, 52);
			}
		`
		minimizeColor = `
			#BtnMinimize {
				background-color: rgb(253, 190, 65);
				border-color: rgb(239, 170, 47);
			}
		`
		closeColor = `
			#BtnClose {
				background-color: rgb(252, 98, 93);
				border-color: rgb(239, 75, 71);
			}
		`
	} else {
		restoreAndMaximizeColor = `
			#BtnRestore, #BtnMaximize {
				background-color: rgba(128, 128, 128, 0.3);
				border-color: rgb(128, 128, 128, 0.2);
			}
		`
		minimizeColor = `
			#BtnMinimize {
				background-color: rgba(128, 128, 128, 0.3);
				border-color: rgb(128, 128, 128, 0.2);
			}
		`
		closeColor = `
			#BtnClose {
				background-color: rgba(128, 128, 128, 0.3);
				border-color: rgb(128, 128, 128, 0.2);
			}
		`
	}
	RestoreColorHover := `
		#BtnRestore:hover {
			background-color: rgb(53, 202, 74);
			border-color: rgb(34, 182, 52);
			background-image: url(":/icons/RestoreHoverDarwin.png");
			background-repeat: no-repeat;
			background-position: center center; 
		}
	`
	minimizeColorHover := `
		#BtnMinimize:hover {
			background-color: rgb(253, 190, 65);
			border-color: rgb(239, 170, 47);
			background-image: url(":/icons/MinimizeHoverDarwin.png");
			background-repeat: no-repeat;
			background-position: center center; 
		}
	`
	closeColorHover := `
		#BtnClose:hover {
			background-color: rgb(252, 98, 93);
			border-color: rgb(239, 75, 71);
			background-image: url(":/icons/CloseHoverDarwin.png");
			background-repeat: no-repeat;
			background-position: center center; 
		}
	`
	f.BtnMinimize.SetStyleSheet(baseStyle + minimizeColor + minimizeColorHover)
	f.BtnRestore.SetStyleSheet(baseStyle + restoreAndMaximizeColor + RestoreColorHover)
	f.BtnClose.SetStyleSheet(baseStyle + closeColor + closeColorHover)
}

// SetupContent ...
func (f *QFramelessWindow) SetupContent(layout widgets.QLayout_ITF) {
	f.Content.SetLayout(layout)
}

// UpdateWidget ...
func (f *QFramelessWindow) UpdateWidget() {
	f.Widget.Update()
	f.Update()
}

// SetupWindowActions ...
func (f *QFramelessWindow) SetupWindowActions() {
	f.ConnectEventFilter(func(watched *core.QObject, event *core.QEvent) bool {
		e := gui.NewQMouseEventFromPointer(core.PointerFromQEvent(event))
		switch event.Type() {
		case core.QEvent__ActivationChange:
			f.SetupTitleBarColor()

		case core.QEvent__HoverMove:
			f.updateCursorShape(e.GlobalPos())

		case core.QEvent__Leave:
			cursor := gui.NewQCursor()
			cursor.SetShape(core.Qt__ArrowCursor)
			f.SetCursor(cursor)

		case core.QEvent__MouseMove:
			f.mouseMove(e)

		case core.QEvent__MouseButtonPress:
			f.mouseButtonPressed(e)

		case core.QEvent__MouseButtonRelease:
			f.isDragStart = false
			f.isLeftButtonPressed = false
			f.hoverEdge = None

		default:
		}

		return f.Widget.EventFilter(watched, event)
	})
}

func (f *QFramelessWindow) mouseMove(e *gui.QMouseEvent) {
	window := f
	margin := f.shadowMargin

	if f.isLeftButtonPressed {

		if f.hoverEdge != None {
			X := e.GlobalPos().X()
			Y := e.GlobalPos().Y()

			if f.MousePos[0] == X && f.MousePos[1] == Y {
				return
			}

			f.MousePos[0] = X
			f.MousePos[1] = Y
			left := window.FrameGeometry().Left() + margin
			top := window.FrameGeometry().Top() + margin
			right := window.FrameGeometry().Right() - margin
			bottom := window.FrameGeometry().Bottom() - margin

			switch f.hoverEdge {
			case Top:
				top = Y
			case Bottom:
				bottom = Y
			case Left:
				left = X
			case Right:
				right = X
			case TopLeft:
				top = Y
				left = X
			case TopRight:
				top = Y
				right = X
			case BottomLeft:
				bottom = Y
				left = X
			case BottomRight:
				bottom = Y
				right = X
			default:
			}

			topLeftPoint := core.NewQPoint2(left, top)
			rightBottomPoint := core.NewQPoint2(right, bottom)
			rect := core.NewQRect2(topLeftPoint, rightBottomPoint)

			// minimum size
			minimumWidth := f.minimumWidth
			minimumHeight := f.minimumHeight
			if rect.Width() <= minimumWidth {
				switch f.hoverEdge {
				case Left:
					left = right - minimumWidth
				case Right:
					right = left + minimumWidth
				case TopLeft:
					left = right - minimumWidth
				case TopRight:
					right = left + minimumWidth
				case BottomLeft:
					left = right - minimumWidth
				case BottomRight:
					right = left + minimumWidth
				default:
				}
			}
			if rect.Height() <= minimumHeight {
				switch f.hoverEdge {
				case Top:
					top = bottom - minimumHeight
				case Bottom:
					bottom = top + minimumHeight
				case TopLeft:
					top = bottom - minimumHeight
				case TopRight:
					top = bottom - minimumHeight
				case BottomLeft:
					bottom = top + minimumHeight
				case BottomRight:
					bottom = top + minimumHeight
				default:
				}
			}
			if rect.Width() <= minimumWidth || rect.Height() <= minimumHeight {
				right = right - 1
				bottom = bottom - 1
			}

			topLeftPoint = core.NewQPoint2(left-margin, top-margin)
			rightBottomPoint = core.NewQPoint2(right+margin, bottom+margin)
			newRect := core.NewQRect2(topLeftPoint, rightBottomPoint)

			window.SetGeometry(newRect)
		}
	}
}

func (f *QFramelessWindow) mouseButtonPressed(e *gui.QMouseEvent) {
	f.hoverEdge = f.calcCursorPos(e.GlobalPos(), f.FrameGeometry())
	if f.hoverEdge != None {
		f.isLeftButtonPressed = true
	}
}

func (f *QFramelessWindow) mouseButtonPressedForWin(e *gui.QMouseEvent) {
	if f.hoverEdge != None {
		f.isLeftButtonPressed = true
	}
}

func (f *QFramelessWindow) updateCursorShape(pos *core.QPoint) {
	if f.isLeftButtonPressed {
		return
	}
	window := f
	cursor := gui.NewQCursor()
	if window.IsFullScreen() || window.IsMaximized() {
		if f.isCursorChanged {
			cursor.SetShape(core.Qt__ArrowCursor)
			window.SetCursor(cursor)
		}
	}
	hoverEdge := f.calcCursorPos(pos, window.FrameGeometry())
	f.isCursorChanged = true
	switch hoverEdge {
	case Top, Bottom:
		cursor.SetShape(core.Qt__SizeVerCursor)
		window.SetCursor(cursor)
	case Left, Right:
		cursor.SetShape(core.Qt__SizeHorCursor)
		window.SetCursor(cursor)
	case TopLeft, BottomRight:
		cursor.SetShape(core.Qt__SizeFDiagCursor)
		window.SetCursor(cursor)
	case TopRight, BottomLeft:
		cursor.SetShape(core.Qt__SizeBDiagCursor)
		window.SetCursor(cursor)
	default:
		cursor.SetShape(core.Qt__ArrowCursor)
		window.SetCursor(cursor)
		f.isCursorChanged = false
	}
}

func (f *QFramelessWindow) calcCursorPos(pos *core.QPoint, rect *core.QRect) Edge {
	rectX := rect.X()
	rectY := rect.Y()
	rectWidth := rect.Width()
	rectHeight := rect.Height()
	posX := pos.X()
	posY := pos.Y()

	edge := f.detectEdgeOnCursor(posX, posY, rectX, rectY, rectWidth, rectHeight)
	return edge
}

func (f *QFramelessWindow) detectEdgeOnCursor(posX, posY, rectX, rectY, rectWidth, rectHeight int) Edge {
	doubleBorderSize := f.borderSize * 2
	octupleBorderSize := f.borderSize * 8
	topBorderSize := 2 - 1

	margin := f.shadowMargin
	rectX = rectX + margin
	rectY = rectY + margin
	rectWidth = rectWidth - (2 * margin)
	rectHeight = rectHeight - (2 * margin)

	var onLeft, onRight, onBottom, onTop, onBottomLeft, onBottomRight, onTopRight, onTopLeft bool

	onBottomLeft = (((posX <= (rectX + octupleBorderSize)) && posX >= rectX &&
		(posY <= (rectY + rectHeight)) && (posY >= (rectY + rectHeight - doubleBorderSize))) ||
		((posX <= (rectX + doubleBorderSize)) && posX >= rectX &&
			(posY <= (rectY + rectHeight)) && (posY >= (rectY + rectHeight - octupleBorderSize))))

	if onBottomLeft {
		return BottomLeft
	}

	onBottomRight = (((posX >= (rectX + rectWidth - octupleBorderSize)) && (posX <= (rectX + rectWidth)) &&
		(posY >= (rectY + rectHeight - doubleBorderSize)) && (posY <= (rectY + rectHeight))) ||
		((posX >= (rectX + rectWidth - doubleBorderSize)) && (posX <= (rectX + rectWidth)) &&
			(posY >= (rectY + rectHeight - octupleBorderSize)) && (posY <= (rectY + rectHeight))))

	if onBottomRight {
		return BottomRight
	}

	onTopRight = (posX >= (rectX + rectWidth - doubleBorderSize)) && (posX <= (rectX + rectWidth)) &&
		(posY >= rectY) && (posY <= (rectY + doubleBorderSize))
	if onTopRight {
		return TopRight
	}

	onTopLeft = posX >= rectX && (posX <= (rectX + doubleBorderSize)) &&
		posY >= rectY && (posY <= (rectY + doubleBorderSize))
	if onTopLeft {
		return TopLeft
	}

	onLeft = (posX >= (rectX - doubleBorderSize)) && (posX <= (rectX + doubleBorderSize)) &&
		(posY <= (rectY + rectHeight - doubleBorderSize)) &&
		(posY >= rectY+doubleBorderSize)
	if onLeft {
		return Left
	}

	onRight = (posX >= (rectX + rectWidth - doubleBorderSize)) &&
		(posX <= (rectX + rectWidth)) &&
		(posY >= (rectY + doubleBorderSize)) && (posY <= (rectY + rectHeight - doubleBorderSize))
	if onRight {
		return Right
	}

	onBottom = (posX >= (rectX + doubleBorderSize)) && (posX <= (rectX + rectWidth - doubleBorderSize)) &&
		(posY >= (rectY + rectHeight - doubleBorderSize)) && (posY <= (rectY + rectHeight))
	if onBottom {
		return Bottom
	}

	onTop = (posX >= (rectX + doubleBorderSize)) && (posX <= (rectX + rectWidth - doubleBorderSize)) &&
		(posY >= rectY) && (posY <= (rectY + topBorderSize))
	if onTop {
		return Top
	}

	return None
}

// Hex ...
func (c *RGB) Hex() string {
	return fmt.Sprintf("#%02x%02x%02x", uint8(c.R), uint8(c.G), uint8(c.B))
}

// Brend ...
func (c *RGB) Brend(color *RGB, alpha float64) *RGB {
	if color == nil {
		return &RGB{0, 0, 0}
	}
	return &RGB{
		R: uint16((float64(c.R) * float64(1-alpha)) + (float64(color.R) * float64(alpha))),
		G: uint16((float64(c.G) * float64(1-alpha)) + (float64(color.G) * float64(alpha))),
		B: uint16((float64(c.B) * float64(1-alpha)) + (float64(color.B) * float64(alpha))),
	}
}
