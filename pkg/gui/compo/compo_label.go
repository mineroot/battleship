package compo

import (
	"battleship/pkg/gui/typing"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
)

type Label struct {
	base
}

func NewLabel(pos pixel.Vec, caption string, size typing.Size, color color.Color) *Label {
	canvas := pixelgl.NewCanvas(typing.BoundsOf(caption, pixel.V(0, 20), size))
	label := &Label{
		base: newBase(canvas),
	}
	label.Pos = pos
	typing.TypeOnCanvas(label.Canvas, caption, typing.Center, pixel.V(0, 20), size, color)
	label.Size = label.Canvas.Bounds().Size()
	return label
}

func (l *Label) Update(*pixelgl.Window, float64) {
	// nop
}
