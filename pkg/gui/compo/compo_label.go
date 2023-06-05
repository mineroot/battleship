package compo

import (
	"battleship/pkg/gui/typing"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
)

type Label struct {
	base
	caption string
	align   typing.Align
	size    typing.Size
	color   color.Color
}

func NewLabel(pos pixel.Vec, caption string, align typing.Align, size typing.Size, color color.Color) *Label {
	label := &Label{
		base:    newBase(nil),
		caption: caption,
		align:   align,
		size:    size,
		color:   color,
	}
	label.Pos = pos
	label.typeOnCanvas()
	return label
}

func (l *Label) SetCaption(caption string) {
	l.caption = caption
	l.typeOnCanvas()
}

func (l *Label) Update(*pixelgl.Window, float64) {
	// nop
}

func (l *Label) typeOnCanvas() {
	bounds := typing.BoundsOf(l.caption, pixel.ZV, l.size)
	// canvas with zero bounds will panic
	if bounds == pixel.ZR {
		bounds = pixel.R(0, 0, 1, 1)
	}
	canvas := pixelgl.NewCanvas(bounds)
	l.Canvas = canvas
	typing.TypeOnCanvas(l.Canvas, l.caption, l.align, pixel.ZV, l.size, l.color)
	l.Size = l.Canvas.Bounds().Size()
}
