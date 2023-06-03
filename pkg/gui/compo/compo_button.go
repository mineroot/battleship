package compo

import (
	"battleship/pkg/gui/sprites"
	"battleship/pkg/gui/typing"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Button struct {
	base
	canvasDefault   *pixelgl.Canvas
	canvasMouseDown *pixelgl.Canvas
	isPressed       bool
}

func NewButton(pos pixel.Vec, caption string) *Button {
	canvasDefault := sprites.CreateCanvas(sprites.ButtonBlueDefault)
	canvasMouseDown := sprites.CreateCanvas(sprites.ButtonBlueMouseDown)
	btn := &Button{
		base:            newBase(canvasDefault),
		canvasDefault:   canvasDefault,
		canvasMouseDown: canvasMouseDown,
	}
	btn.Pos = pos
	btn.Size = btn.Canvas.Bounds().Size()
	btn.On(MouseLDown, btn.mouseDownButton)
	btn.On(MouseLeave, btn.releaseButton)
	btn.On(MouseLUp, btn.releaseButton)
	typing.TypeOnCanvas(btn.canvasDefault, caption, typing.Center, pixel.V(0, 20), typing.Size39, colornames.Yellow)
	typing.TypeOnCanvas(btn.canvasMouseDown, caption, typing.Center, pixel.V(0, 15), typing.Size39, colornames.Gold)
	return btn
}

func (b *Button) Update(win *pixelgl.Window, dt float64) {
	b.base.Update(win, dt)
	if b.isPressed {
		b.Canvas = b.canvasMouseDown
		return
	}
	b.Canvas = b.canvasDefault
}

func (b *Button) mouseDownButton(...any) {
	b.isPressed = true
}

func (b *Button) releaseButton(...any) {
	b.isPressed = false
}
