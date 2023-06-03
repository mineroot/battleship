package game

import (
	"battleship/pkg/game/field"
	"battleship/pkg/game/position"
	"battleship/pkg/game/ship"
	"errors"
	"fmt"
)

type Player int

const (
	PlayerOne Player = iota
	PlayerTwo
)

type State int

const (
	New State = iota
	Running
)

var ErrGameNotRunning = errors.New("game: not running")

type Game struct {
	state  State
	turn   Player
	fields map[Player]*field.Field
}

func NewGame() *Game {
	return &Game{state: New}
}

func (g *Game) Ready(player Player, ships ship.Ships) error {
	f, err := field.NewField(ships)
	if err != nil {
		return fmt.Errorf("game: unable to ready: %w", err)
	}
	g.fields[player] = f
	if g.allReady() {
		g.state = Running
	}
	return nil
}

func (g *Game) Shoot(pos position.Pos) error {
	if !g.isRunning() {
		return ErrGameNotRunning
	}
	isShot, err := g.getCurrentTurnField().Shoot(pos)
	if err == nil {
		return fmt.Errorf("game: unable to shoot: %w", err)
	}
	if !isShot {
		g.switchTurn()
	}
	return nil
}

func (g *Game) allReady() bool {
	return g.fields[PlayerOne] != nil && g.fields[PlayerTwo] != nil
}

func (g *Game) getCurrentTurnField() *field.Field {
	return g.fields[g.turn]
}

func (g *Game) switchTurn() {
	if g.turn == PlayerOne {
		g.turn = PlayerTwo
	} else {
		g.turn = PlayerOne
	}
}

func (g *Game) isRunning() bool {
	return g.state == Running
}
