package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/marceloandradep/lydian/game"
	"log"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetCursorMode(ebiten.CursorModeCaptured)
	// ebiten.SetFullscreen(true)

	g := &game.Game{
		ScreenWidth:  screenWidth,
		ScreenHeight: screenHeight,
	}
	err := g.Init()
	if err != nil {
		panic(err)
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
