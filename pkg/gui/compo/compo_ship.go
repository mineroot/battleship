package compo

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"math"
	"time"
)

type Ship struct {
	base
	originPos pixel.Vec
	decks     int
	dragDelta pixel.Vec
}

func NewShip(decks int, originPos pixel.Vec) *Ship {
	canvas := pixelgl.NewCanvas(pixel.R(0, 0, float64(cellSize*decks), cellSize))
	s := &Ship{
		base:      newBase(canvas),
		originPos: originPos,
		decks:     decks,
	}

	s.Pos = originPos
	s.Size = s.Canvas.Bounds().Size()
	color := pixel.RGB(0, 0, 1).Mul(pixel.Alpha(0.05))
	s.Canvas.Clear(color)
	border := imdraw.New(nil)
	border.Color = colornames.Blue
	border.Push(s.Canvas.Bounds().Min, s.Canvas.Bounds().Max)
	border.Rectangle(4)
	border.Draw(s.Canvas)
	s.On(DragStart, func(data ...any) {
		pos := data[0].(pixel.Vec)
		s.dragDelta = pos.Sub(s.Pos)
	})
	s.On(Dragging, func(data ...any) {
		mouseMos := data[0].(pixel.Vec)
		s.Pos = mouseMos.Sub(s.dragDelta)
	})
	s.On(DragEnd, func(data ...any) {
		s.roundPos()
	})
	s.On(Click, func(data ...any) {
		clickDuration := data[1].(time.Duration)
		if clickDuration < time.Millisecond*150 {
			s.switchOrientation()
			s.alignAfterOrientationSwitch()
		}
	})
	return s
}

func (s *Ship) roundPos() {
	isDecksEven := math.Mod(float64(s.decks), 2) == 0
	shiftX, shiftY := 25.0, 25.0
	if isDecksEven {
		if s.Orientation == Horizontal {
			shiftX = 0
		} else {
			shiftY = 0
		}
	}
	// round to nearest cellSize
	s.Pos.X = math.Round((s.Pos.X-shiftX)/cellSize)*cellSize + shiftX
	s.Pos.Y = math.Round((s.Pos.Y-shiftY)/cellSize)*cellSize + shiftY
}

func (s *Ship) alignAfterOrientationSwitch() {
	isDecksEven := math.Mod(float64(s.decks), 2) == 0
	if !isDecksEven {
		return
	}

	shift := 25.0
	if s.Orientation == Horizontal {
		s.Pos.X += shift
		s.Pos.Y -= shift
	} else {
		s.Pos.X -= shift
		s.Pos.Y += shift
	}
}
