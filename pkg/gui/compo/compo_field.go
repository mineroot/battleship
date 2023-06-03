package compo

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Field struct {
	base
}

func NewField() *Field {
	size := pixel.V(10*cellSize, 10*cellSize)
	f := &Field{
		base: newBase(pixelgl.NewCanvas(pixel.Rect{
			Min: pixel.ZV,
			Max: size,
		})),
	}
	f.Pos = pixel.V(300, 300)
	f.Size = size
	imd := imdraw.New(nil)
	imd.Color = colornames.Blue
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(f.Canvas.Bounds().Min, f.Canvas.Bounds().Max)
	imd.Rectangle(4)
	imd.Draw(f.Canvas)
	return f
}

func (f *Field) Update(*pixelgl.Window, float64) {
	// nop
}
