package compo

import (
	"battleship/pkg/game/field"
	"battleship/pkg/game/position"
	"battleship/pkg/game/ship"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Field struct {
	base
	ships []*Ship
}

func NewField() *Field {
	size := pixel.V(10*cellSize, 10*cellSize)
	f := &Field{
		base: newBase(pixelgl.NewCanvas(pixel.Rect{
			Min: pixel.ZV,
			Max: size,
		})),
	}
	f.Pos = pixel.V(300, 400)
	f.Size = size
	f.initShipsSetupScene()
	imd := imdraw.New(nil)
	imd.Color = colornames.Blue
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(f.Canvas.Bounds().Min, f.Canvas.Bounds().Max)
	imd.Rectangle(4)
	imd.Draw(f.Canvas)
	return f
}

func (f *Field) Validate() error {
	const fieldMarginX, fieldMarginY float64 = 50, 150
	ships := ship.Ships{}
	for _, s := range f.ships {
		if f.Rect().Contains(s.Pos) {
			var shipLenX, shipLenY float64
			if s.Orientation == Horizontal {
				shipLenX = cellSize * float64(s.decks)
				shipLenY = cellSize
			} else {
				shipLenX = cellSize
				shipLenY = cellSize * float64(s.decks)
			}
			x := (s.Pos.X - fieldMarginX - (shipLenX / 2)) / cellSize
			y := (s.Pos.Y - fieldMarginY - (shipLenY / 2)) / cellSize
			// flip 'y' to be (0,0) top left corner, instead of (0,0) bottom left corner
			y = field.Size - y - (shipLenY / cellSize)

			ships = append(
				ships,
				ship.New(position.New(int(x), int(y)), ship.Variant(s.decks-1), ship.Orientation(s.Orientation)),
			)
		}
	}
	_, err := field.NewField(ships)
	return err
}

func (f *Field) Update(win *pixelgl.Window, dt float64) {
	for _, s := range f.ships {
		s.Update(win, dt)
		if s.isMouseDown {
			break // prevent dragging multiple ships
		}
	}
}

func (f *Field) Draw(target pixel.ComposeTarget, dt float64) {
	f.base.Draw(target, dt)
	for _, s := range f.ships {
		s.Draw(target, dt)
	}
}

func (f *Field) initShipsSetupScene() {
	f.ships = []*Ship{
		NewShip(1, pixel.V(775, 225)),
		NewShip(1, pixel.V(875, 225)),
		NewShip(1, pixel.V(975, 225)),
		NewShip(1, pixel.V(1075, 225)),
		NewShip(2, pixel.V(800, 325)),
		NewShip(2, pixel.V(950, 325)),
		NewShip(2, pixel.V(1100, 325)),
		NewShip(3, pixel.V(825, 425)),
		NewShip(3, pixel.V(1025, 425)),
		NewShip(4, pixel.V(850, 525)),
	}
}
