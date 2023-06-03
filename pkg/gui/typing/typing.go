package typing

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
	"image/color"
)

type Align int

const (
	Left Align = iota
	Center
	Right
)

type Size float64

const (
	Size13 = iota + 1
	Size26
	Size39
	Size52
	Size65
	Size78
)

var basicAtlas = text.NewAtlas(basicfont.Face7x13, text.ASCII)

func TypeOnCanvas(canvas *pixelgl.Canvas, txt string, align Align, padding pixel.Vec, size Size, color color.Color) {
	txtObj := text.New(padding, basicAtlas)
	txtWidth := txtObj.BoundsOf(txt).W()
	switch align {
	case Left:
	case Center:
		newX := (canvas.Bounds().W() - (float64(size) * txtWidth)) / 2
		txtObj.Orig.X, txtObj.Dot.X = newX, newX
	case Right:
		newX := canvas.Bounds().W() - (float64(size) * txtWidth) - padding.X
		txtObj.Orig.X, txtObj.Dot.X = newX, newX
	default:
		panic(fmt.Errorf("unreachable"))
	}
	txtObj.Color = color
	_, _ = fmt.Fprintln(txtObj, txt)
	txtObj.Draw(canvas, pixel.IM.Scaled(txtObj.Orig, float64(size)))
}

func BoundsOf(txt string, padding pixel.Vec, size Size) pixel.Rect {
	txtObj := text.New(padding, basicAtlas)
	bounds := txtObj.BoundsOf(txt)
	return pixel.Rect{
		Min: pixel.ZV,
		Max: pixel.V(bounds.W(), bounds.H()).Scaled(float64(size)),
	}
}
