package position

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	pos := New(5, 5)
	assert.Equal(t, Pos{X: 5, Y: 5}, pos)
}

func TestPos_Add(t *testing.T) {
	pos := New(5, 5).Add(New(3, 2))
	assert.Equal(t, Pos{X: 8, Y: 7}, pos)
}

func TestPos_InField(t *testing.T) {
	pos := New(5, 5)
	assert.True(t, pos.InField(10, 10))
	assert.False(t, pos.InField(3, 10))
	pos = New(-1, 5)
	assert.False(t, pos.InField(10, 10))
}

func TestPositions_Empty(t *testing.T) {
	var positions Positions
	assert.True(t, positions.Empty())
	positions = append(positions, Zero)
	assert.False(t, positions.Empty())
}

func TestPositions_Has(t *testing.T) {
	positions := Positions{
		Zero,
		New(0, 1),
		New(0, 2),
	}
	assert.True(t, positions.Has(Zero))
	assert.True(t, positions.Has(New(0, 1)))
	assert.True(t, positions.Has(New(0, 2)))
	assert.False(t, positions.Has(New(2, 0)))
}

func TestPositions_InField(t *testing.T) {
	positions := Positions{
		Zero,
		New(0, 1),
		New(0, 2),
	}
	assert.True(t, positions.InField(10, 10))
	assert.False(t, positions.InField(2, 1))
}

func TestPositions_Intersect(t *testing.T) {
	positions1 := Positions{
		New(0, 1),
		New(0, 2),
	}
	positions2 := Positions{
		New(1, 0),
		New(2, 0),
	}
	assert.Empty(t, positions1.Intersect(positions2))
	positions1 = Positions{
		Zero,
		New(0, 1),
		New(0, 2),
	}
	positions2 = Positions{
		New(1, 0),
		Zero,
		New(2, 0),
	}
	assert.Equal(t, Positions{Zero}, positions1.Intersect(positions2))
}
