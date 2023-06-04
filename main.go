package main

import (
	"battleship/pkg/gui/sprites"
	"battleship/pkg/p2p"
	"battleship/pkg/run"
	"embed"
	"github.com/faiface/pixel/pixelgl"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

//go:embed assets
var assetsDir embed.FS

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	sprites.Init(assetsDir)
	pixelgl.Run(run.Run)
	p2p.Close()
}
