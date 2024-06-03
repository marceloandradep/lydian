package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"lydian/internal/game"
)

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetCursorMode(ebiten.CursorModeCaptured)

	g := &game.GameUVN{}
	err := g.Init()
	if err != nil {
		panic(err)
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
