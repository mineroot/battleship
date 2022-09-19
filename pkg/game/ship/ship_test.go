package ship

import (
	"battleship/pkg/game/position"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestShip_HeadAndTails(t *testing.T) {
	ship := New(position.Zero, Three, Horizontal)
	headAndTails := ship.HeadAndTails()
	require.Len(t, headAndTails, 3)
	assert.Equal(t, ship.Head(), headAndTails[0])
	assert.Equal(t, position.Positions{position.Zero, position.New(1, 0), position.New(2, 0)}, headAndTails)

	ship = New(position.Zero, Two, Vertical)
	headAndTails = ship.HeadAndTails()
	require.Len(t, headAndTails, 2)
	assert.Equal(t, ship.Head(), headAndTails[0])
	assert.Equal(t, position.Positions{position.Zero, position.New(0, 1)}, headAndTails)
}

func TestShip_Neighborhoods(t *testing.T) {
	ship := New(position.New(1, 1), One, Horizontal)
	neighborhoods := ship.Neighborhoods()
	// expected len = 2*variant + 6,
	// e.g. for variant(1) = 2*1 + 6 = 8
	require.Len(t, neighborhoods, 8)
	assert.Equal(t, position.Positions{
		position.New(0, 0),
		position.New(1, 0),
		position.New(2, 0),
		position.New(0, 1),
		position.New(2, 1),
		position.New(0, 2),
		position.New(1, 2),
		position.New(2, 2),
	}, neighborhoods)
}
