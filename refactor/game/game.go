package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"lydian/refactor/camera"
	"lydian/refactor/geometry"
	"lydian/refactor/graphics"
	"lydian/refactor/math"
	"lydian/refactor/rendering"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

type Game struct {
	cameraPos    *math.Vector
	cameraRot    *math.Vector
	camera       *camera.Camera
	forward      *math.Vector
	left         *math.Vector
	lastMouseX   int
	lastMouseY   int
	triangleList []*rendering.Triangle3D
	clipper      graphics.Clipper
}

func (g *Game) Init() error {
	g.clipper = graphics.Clipper{MinX: 0, MinY: 0, MaxX: screenWidth - 1, MaxY: screenHeight - 1}

	g.cameraPos = math.NewVector3(0, 0, 0)
	g.cameraRot = math.NewVector3(0, 0, 0)
	g.forward = math.NewVector3(0, 0, 1)
	g.left = math.NewVector3(-1, 0, 0)

	viewPortSize := &geometry.Dimension{
		Width:  screenWidth - 1,
		Height: screenHeight - 1,
	}

	euler := camera.NewEuler(g.cameraPos, g.cameraRot, math.RotationZYX, 20, 500, 90, viewPortSize)
	g.camera = euler

	g.triangleList = createMap(1000)

	return nil
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	x, y := ebiten.CursorPosition()

	if g.lastMouseX == 0 {
		g.lastMouseX = x
	}

	if g.lastMouseY == 0 {
		g.lastMouseY = y
	}

	dx := (x - g.lastMouseX) / 10
	dy := (y - g.lastMouseY) / 10

	g.lastMouseX = x
	g.lastMouseY = y

	g.cameraRot.X -= float64(dy)
	g.cameraRot.Y -= float64(dx)

	rot := math.RotationMatrix(float64(dx), math.YAxis)
	g.forward = rot.MultiplyVertex(g.forward)
	g.left = rot.MultiplyVertex(g.left)

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.cameraPos = g.cameraPos.Add(g.forward)
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.cameraPos = g.cameraPos.Sub(g.forward)
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.cameraPos = g.cameraPos.Add(g.left)
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.cameraPos = g.cameraPos.Sub(g.left)
	}

	g.camera.SetPos(g.cameraPos)
	g.camera.Rotate(g.cameraRot)

	g.camera.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.camera.Clear()

	for _, t := range g.triangleList {
		g.camera.AddToCamera(t)
	}

	graphics.Rasterize(screen, g.clipper, g.camera)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
