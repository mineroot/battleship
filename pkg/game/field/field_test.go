package field

import (
	"battleship/pkg/game/position"
	"battleship/pkg/game/ship"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	variantOneIndex  = 0
	variantFourIndex = 9
)

func (f *Field) getSankShips() ship.Ships {
	var sankShips ship.Ships
	for _, s := range f.ships {
		if s.IsSank() {
			sankShips = append(sankShips, s)
		}
	}
	return sankShips
}

/*
	ships layout
	h - head (pos), t - tail:

| h |   | h |   | h |   | h |   | h | t |
|   |   |   |   |   |   |   |   |   |   |
| h | t |   | h | t |   | h | t | t |   |
|   |   |   |   |   |   |   |   |   |   |
| h | t | t |   | h | t | t | t |   |   |
|   |   |   |   |   |   |   |   |   |   |
|   |   |   |   |   |   |   |   |   |   |
|   |   |   |   |   |   |   |   |   |   |
|   |   |   |   |   |   |   |   |   |   |
|   |   |   |   |   |   |   |   |   |   |
*/
func getValidShips() ship.Ships {
	return ship.Ships{
		ship.New(position.New(0, 0), ship.One, ship.Horizontal),
		ship.New(position.New(2, 0), ship.One, ship.Horizontal),
		ship.New(position.New(4, 0), ship.One, ship.Horizontal),
		ship.New(position.New(6, 0), ship.One, ship.Horizontal),

		ship.New(position.New(8, 0), ship.Two, ship.Horizontal),
		ship.New(position.New(0, 2), ship.Two, ship.Horizontal),
		ship.New(position.New(3, 2), ship.Two, ship.Horizontal),

		ship.New(position.New(6, 2), ship.Three, ship.Horizontal),
		ship.New(position.New(0, 4), ship.Three, ship.Horizontal),

		ship.New(position.New(4, 4), ship.Four, ship.Horizontal),
	}
}

func TestNewField(t *testing.T) {
	t.Run("empty ships", func(t *testing.T) {
		var ships ship.Ships
		field, err := NewField(ships)
		assert.Nil(t, field)
		assert.ErrorIs(t, err, ErrVariantCount)
	})

	t.Run("invalid ships count", func(t *testing.T) {
		ships := append(getValidShips(), ship.New(position.Zero, ship.Four, ship.Vertical)) // add extra 4-deck ship
		field, err := NewField(ships)
		assert.Nil(t, field)
		assert.ErrorIs(t, err, ErrVariantCount)
	})

	t.Run("ship out of field #1", func(t *testing.T) {
		ships := getValidShips()
		// setting 4-deck ship pos to (0, 7)
		// so the part of ship will be out of field
		ships[variantFourIndex] = ship.New(position.New(0, 7), ship.Four, ship.Vertical)
		field, err := NewField(ships)
		assert.Nil(t, field)
		assert.ErrorIs(t, err, ErrShipOutOfField)
	})

	t.Run("ship out of field #2", func(t *testing.T) {
		ships := getValidShips()
		// setting first 1-deck ship pos to (-1, 5)
		// so ship will be out of field
		ships[variantOneIndex] = ship.New(position.New(-1, 5), ship.One, ship.Horizontal)
		field, err := NewField(ships)
		assert.Nil(t, field)
		assert.ErrorIs(t, err, ErrShipOutOfField)
	})

	t.Run("ships overlap #1", func(t *testing.T) {
		ships := getValidShips()
		// place 4-deck ship over 3-deck ship,
		// so there will be an overlap
		ships[variantFourIndex] = ship.New(position.New(0, 4), ship.Four, ship.Vertical)
		field, err := NewField(ships)
		assert.Nil(t, field)
		assert.ErrorIs(t, err, ErrShipsOverlap)
	})
	t.Run("ships overlap #2", func(t *testing.T) {
		ships := getValidShips()
		// place 4-deck ship near 3-deck ship,
		// so there will be an overlap
		ships[variantFourIndex] = ship.New(position.New(3, 4), ship.Four, ship.Horizontal)
		field, err := NewField(ships)
		assert.Nil(t, field)
		assert.ErrorIs(t, err, ErrShipsOverlap)
	})

	t.Run("success", func(t *testing.T) {
		validShips := getValidShips()
		field, err := NewField(validShips)
		assert.NoError(t, err)
		require.NotNil(t, field)
		assert.Equal(t, field.ships, validShips)
	})
}

func TestField_Shoot(t *testing.T) {
	t.Run("out of field", func(t *testing.T) {
		f, _ := NewField(getValidShips())
		outOfFieldPos := position.New(-1, 0)
		isShot, err := f.Shoot(outOfFieldPos)
		assert.False(t, isShot)
		assert.ErrorIs(t, err, ErrPosOutOfField)
		outOfFieldPos = position.New(Size, 0)
		isShot, err = f.Shoot(outOfFieldPos)
		assert.False(t, isShot)
		assert.ErrorIs(t, err, ErrPosOutOfField)
	})
	t.Run("empty pos & already shot", func(t *testing.T) {
		f, _ := NewField(getValidShips())
		emptyPos := position.New(1, 0)
		isShot, err := f.Shoot(emptyPos)
		assert.False(t, isShot)
		assert.NoError(t, err)
		assert.True(t, f.shots[emptyPos.X][emptyPos.Y])
		assert.Empty(t, f.getSankShips())
		isShot, err = f.Shoot(emptyPos)
		assert.False(t, isShot)
		assert.ErrorIs(t, err, ErrAlreadyShot)
	})
	t.Run("sink ship", func(t *testing.T) {
		f, _ := NewField(getValidShips())
		shipPos := position.Zero
		isShot, err := f.Shoot(shipPos)
		assert.True(t, isShot)
		assert.NoError(t, err)
		assert.True(t, f.shots[shipPos.X][shipPos.Y])
		assert.True(t, f.shots[shipPos.X+1][shipPos.Y])
		assert.True(t, f.shots[shipPos.X+1][shipPos.Y+1])
		assert.True(t, f.shots[shipPos.X][shipPos.Y+1])
		require.Len(t, f.getSankShips(), 1)
		assert.Equal(t, f.getSankShips()[0].Head(), shipPos)
	})
}
