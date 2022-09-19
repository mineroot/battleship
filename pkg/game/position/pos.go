package position

import "fmt"

var Zero = Pos{}

type Pos struct {
	X, Y int
}

func (p Pos) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

func New(x, y int) Pos {
	return Pos{X: x, Y: y}
}

func (p Pos) InField(width, height int) bool {
	return 0 <= p.X && p.X < width &&
		0 <= p.Y && p.Y < height
}

func (p Pos) Add(delta Pos) Pos {
	p.X += delta.X
	p.Y += delta.Y
	return p
}

type Positions []Pos

func (ps Positions) Has(pos Pos) bool {
	for _, p := range ps {
		if p == pos {
			return true
		}
	}
	return false
}

func (ps Positions) Empty() bool {
	return len(ps) == 0
}

func (ps Positions) Intersect(ps2 Positions) Positions {
	var intersection Positions
	for _, pos := range ps {
		for _, pos2 := range ps2 {
			if pos == pos2 {
				intersection = append(intersection, pos)
			}
		}
	}
	return intersection
}

func (ps Positions) InField(width, height int) bool {
	for _, pos := range ps {
		if !pos.InField(width, height) {
			return false
		}
	}
	return true
}
