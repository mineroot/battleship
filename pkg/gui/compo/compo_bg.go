package compo

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image/color"
)

const cellSize = 50

type Bg struct {
	base
}

func NewBg(bounds pixel.Rect) *Bg {
	canvas := pixelgl.NewCanvas(bounds)
	canvas.Clear(colornames.White)

	imd := imdraw.New(nil)
	imd.Color = color.RGBA{R: 0x02, G: 0xAF, B: 0xEF, A: 0xFF}
	for i := 0; i < 64; i++ {
		x := float64(i * cellSize)
		imd.Push(pixel.V(x, 0), pixel.V(x, bounds.H()))
		imd.Line(2)

		y := float64(i * cellSize)
		imd.Push(pixel.V(0, y), pixel.V(bounds.W(), y))
		imd.Line(2)
	}
	imd.Draw(canvas)

	bg := &Bg{
		base: newBase(canvas),
	}
	bg.Size = bounds.Size()
	bg.Pos = bounds.Size().Scaled(0.5)
	return bg
}

func (bg *Bg) Update(*pixelgl.Window, float64) {
	// nop
}
