package field

import (
	"battleship/pkg/game/position"
	"battleship/pkg/game/ship"
	"errors"
	"fmt"
)

const size = 10

var (
	ErrVariantCount   = errors.New("wrong variant count")
	ErrShipOutOfField = errors.New("ship out of field")
	ErrPosOutOfField  = errors.New("position out of field")
	ErrShipsOverlap   = errors.New("ships overlap")
	ErrAlreadyShot    = errors.New("already shot")
)

type Field struct {
	ships ship.Ships
	shots [size][size]bool
}

func NewField(ships ship.Ships) (*Field, error) {
	wrap := func(err error) error {
		return fmt.Errorf("unable to create new field: %w", err)
	}
	err := validateVariantsCount(ships)
	if err != nil {
		return nil, wrap(err)
	}

	err = validatePositions(ships)
	if err != nil {
		return nil, wrap(err)
	}

	return &Field{ships: ships}, nil
}

func (f *Field) Shoot(pos position.Pos) (bool, error) {
	if !pos.InField(size, size) {
		return false, fmt.Errorf("%w: %s", ErrPosOutOfField, pos)
	}
	if f.shots[pos.X][pos.Y] {
		return false, fmt.Errorf("%w: %s", ErrAlreadyShot, pos)
	}
	f.shots[pos.X][pos.Y] = true
	var shipsHeadsAndTails position.Positions
	for _, s := range f.ships {
		if s.IsSank() {
			continue
		}
		f.trySink(s)
		shipsHeadsAndTails = append(shipsHeadsAndTails, s.HeadAndTails()...)
	}
	return shipsHeadsAndTails.Has(pos), nil
}

func (f *Field) trySink(s *ship.Ship) {
	for _, pos := range s.HeadAndTails() {
		if !f.shots[pos.X][pos.Y] {
			return
		}
	}
	s.Sink()
	for _, p := range s.Neighborhoods() {
		if p.InField(size, size) {
			f.shots[p.X][p.Y] = true
		}
	}
}

func validateVariantsCount(ships ship.Ships) error {
	countByVariant := make(map[ship.Variant]int)
	for _, s := range ships {
		countByVariant[s.Variant()]++
	}

	for _, variant := range ship.Variants() {
		gotCount := countByVariant[variant]
		expectedCount := 4 - int(variant)
		if gotCount != expectedCount {
			return fmt.Errorf("%w: %s-deck: expected %d, got %d", ErrVariantCount, variant, expectedCount, gotCount)
		}
	}
	return nil
}

func validatePositions(ships ship.Ships) error {
	if len(ships) < 2 {
		panic("ships len must be at least 2")
	}
	// check ships for out of field, overlap and intersection with neighborhoods
	for i := 0; i < ships.Count(); i++ {
		shipIHeadAndTails := ships[i].HeadAndTails()
		if !shipIHeadAndTails.InField(size, size) {
			return ErrShipOutOfField
		}
		for j := i + 1; j < len(ships); j++ {
			shipJHeadAndTails := ships[j].HeadAndTails()
			if !shipIHeadAndTails.Intersect(shipJHeadAndTails).Empty() {
				return ErrShipsOverlap
			}
			shipJNeighborhoods := ships[j].Neighborhoods()
			if !shipIHeadAndTails.Intersect(shipJNeighborhoods).Empty() {
				return ErrShipsOverlap
			}
		}
	}
	return nil
}
