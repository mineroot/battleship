package ship

import (
	"battleship/pkg/game/position"
)

type Ship struct {
	head        position.Pos
	variant     Variant
	orientation Orientation
	isSank      bool
}

func New(head position.Pos, variant Variant, orientation Orientation) *Ship {
	return &Ship{head: head, variant: variant, orientation: orientation}
}

func (s *Ship) Head() position.Pos {
	return s.head
}

func (s *Ship) HeadAndTails() position.Positions {
	var headAndTails position.Positions
	decks := s.variant.Decks()
	for i := 0; i < decks; i++ {
		pos := s.Head()
		if s.IsHorizontal() {
			pos.X = pos.X + i
		} else {
			pos.Y = pos.Y + i
		}
		headAndTails = append(headAndTails, pos)
	}
	return headAndTails
}

func (s *Ship) Neighborhoods() position.Positions {
	delta := position.Positions{
		position.New(-1, -1), position.New(0, -1), position.New(1, -1),
		position.New(-1, 0), position.New(1, 0),
		position.New(-1, 1), position.New(0, 1), position.New(1, 1),
	}
	headAndTails := s.HeadAndTails()
	var neighborhoods position.Positions
	for _, pos := range headAndTails {
		for _, d := range delta {
			neighborhood := pos.Add(d)
			if !neighborhoods.Has(neighborhood) && !headAndTails.Has(neighborhood) {
				neighborhoods = append(neighborhoods, neighborhood)
			}
		}
	}
	return neighborhoods
}

func (s *Ship) Variant() Variant {
	return s.variant
}

func (s *Ship) Orientation() Orientation {
	return s.orientation
}

func (s *Ship) IsHorizontal() bool {
	return s.orientation == Horizontal
}

func (s *Ship) IsVertical() bool {
	return s.orientation == Vertical
}

func (s *Ship) Sink() {
	s.isSank = true
}

func (s *Ship) IsSank() bool {
	return s.isSank
}

type Ships []*Ship

func (s Ships) Count() int {
	return len(s)
}
