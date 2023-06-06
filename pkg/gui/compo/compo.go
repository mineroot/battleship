package compo

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"math"
	"time"
)

type Orientation int

const (
	Horizontal Orientation = iota
	Vertical
)

type Compo interface {
	Update(win *pixelgl.Window, dt float64)
	Draw(target pixel.ComposeTarget, dt float64)
	On(event Event, handler Handler)
}

type base struct {
	Pos         pixel.Vec
	Size        pixel.Vec
	Orientation Orientation
	Canvas      *pixelgl.Canvas

	eventHandlers         map[Event][]Handler
	isMouseEnter          bool
	isMouseDown           bool
	mouseLJustPressedTime time.Time
	isFocus               bool
}

func newBase(canvas *pixelgl.Canvas) base {
	return base{
		Canvas:        canvas,
		eventHandlers: make(map[Event][]Handler),
	}
}

func (b *base) Update(win *pixelgl.Window, _ float64) {
	mousePos := win.MousePosition()
	clickOutOfCompo := win.JustPressed(pixelgl.MouseButtonLeft) && !b.Rect().Contains(mousePos)
	if (!win.Focused() || clickOutOfCompo) && b.isFocus {
		b.isFocus = false
		b.triggerEvent(FocusOut)
	}
	if win.JustReleased(pixelgl.MouseButtonLeft) && b.isMouseEnter {
		b.triggerEvent(MouseLUp, mousePos)
	}
	if win.JustReleased(pixelgl.MouseButtonLeft) && b.isMouseDown {
		b.isMouseDown = false
		b.triggerEvent(DragEnd, mousePos)
		if b.isMouseEnter {
			clickDuration := time.Since(b.mouseLJustPressedTime)
			relativeMousePos := pixel.V(mousePos.X-b.Pos.X+(b.Size.X/2), mousePos.Y-b.Pos.Y+(b.Size.Y/2))
			b.triggerEvent(Click, mousePos, relativeMousePos, clickDuration)
			if !b.isFocus {
				b.isFocus = true
				b.triggerEvent(FocusIn)
			}
		}
	}
	if !win.MouseInsideWindow() {
		return
	}
	if b.Rect().Contains(mousePos) && !b.isMouseEnter {
		b.isMouseEnter = true
		b.triggerEvent(MouseEnter, mousePos)
	} else if b.isMouseEnter && !b.Rect().Contains(mousePos) {
		b.isMouseEnter = false
		b.triggerEvent(MouseLeave, mousePos)
	}
	if b.isMouseEnter {
		b.triggerEvent(MouseMove, mousePos)
	}
	if b.isMouseEnter && win.JustPressed(pixelgl.MouseButtonLeft) {
		b.mouseLJustPressedTime = time.Now()
		b.isMouseDown = true
		b.triggerEvent(MouseLDown, mousePos)
		b.triggerEvent(DragStart, mousePos)
	}
	if b.isMouseDown {
		b.triggerEvent(Dragging, mousePos)
	}
}

func (b *base) Draw(target pixel.ComposeTarget, _ float64) {
	angle := 0.0
	if b.Orientation == Vertical {
		angle = math.Pi / 2
	}
	m := pixel.IM.Moved(b.Pos).Rotated(b.Pos, angle)
	b.Canvas.Draw(target, m)
}

func (b *base) On(event Event, handler Handler) {
	b.eventHandlers[event] = append(b.eventHandlers[event], handler)
}

func (b *base) Rect() pixel.Rect {
	var size pixel.Vec
	if b.Orientation == Horizontal {
		size = b.Size
	} else {
		size = pixel.V(b.Size.Y, b.Size.X)
	}
	halfSize := size.Scaled(0.5)

	return pixel.Rect{
		Min: b.Pos.Sub(halfSize),
		Max: b.Pos.Add(halfSize),
	}
}

func (b *base) triggerEvent(event Event, data ...any) {
	handlers, ok := b.eventHandlers[event]
	if !ok {
		return
	}
	for _, handler := range handlers {
		handler(data...)
	}
}

func (b *base) switchOrientation() {
	b.Orientation ^= 1
}
