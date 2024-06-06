package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"lydian/internal/game"
)

const (
	screenWidth  = 2000
	screenHeight = 1500
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
