package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"lydian/internal/geometry"
	"lydian/internal/graphics"
	"lydian/internal/math"
	"lydian/internal/rendering"
	"lydian/internal/rendering/camera"
)

const (
	screenWidth  = 1600
	screenHeight = 1200
)

type Game struct {
	cube       rendering.Object3D
	cubePos    *math.Vector
	cubeRot    *math.Vector
	camera     camera.Camera
	cameraPos  *math.Vector
	cameraRot  *math.Vector
	forward    *math.Vector
	left       *math.Vector
	lastMouseX int
	lastMouseY int
	clipper    graphics.Clipper
}

func (g *Game) Init() error {
	g.clipper = graphics.Clipper{MinX: 0, MinY: 0, MaxX: screenWidth - 1, MaxY: screenHeight - 1}

	pos := math.NewVector(0, 0, 0)
	scale := math.NewVector(5, 5, 5)
	rot := math.NewVector(0, 0, 0)

	cube, err := rendering.Load("/Users/marcelopereira/GolandProjects/lydian/internal/game/resources/cube.plg", *pos, *scale, *rot)
	if err != nil {
		return err
	}
	g.cube = *cube

	g.cameraPos = math.NewVector(0, 0, 0)
	g.cameraRot = math.NewVector(0, 0, 0)

	viewPortSize := geometry.Dimension{
		Width:  screenWidth - 1,
		Height: screenHeight - 1,
	}

	euler := camera.NewEuler(*g.cameraPos, *g.cameraRot, math.RotationZYX, 50, 500, 90, viewPortSize)
	g.camera = euler

	g.cubePos = math.NewVector(0, 0, 100)
	g.cubeRot = math.NewVector(0, 0, 0)

	g.forward = math.NewVector(0, 0, 1)
	g.left = math.NewVector(-1, 0, 0)

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
	g.forward = rot.MultiplyVertex(*g.forward)
	g.left = rot.MultiplyVertex(*g.left)

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.cameraPos = g.cameraPos.Add(*g.forward)
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.cameraPos = g.cameraPos.Sub(*g.forward)
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.cameraPos = g.cameraPos.Add(*g.left)
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.cameraPos = g.cameraPos.Sub(*g.left)
	}

	g.cube.Reset()
	g.cube.SetWorldPos(g.cubePos.X, g.cubePos.Y, g.cubePos.Z)

	g.cubeRot.X += 1
	g.cubeRot.Y += 1
	g.cubeRot.Z += 1

	rotation := math.RotationSequenceMatrix(g.cubeRot.X, g.cubeRot.Y, g.cubeRot.Z, math.RotationZYX)
	g.cube.Transform(rotation, rendering.LocalToTransformed, true)

	g.camera.SetPosition(g.cameraPos.X, g.cameraPos.Y, g.cameraPos.Z)
	g.camera.SetRotation(g.cameraRot.X, g.cameraRot.Y, g.cameraRot.Z)
	g.camera.Update()
	g.cube.Update(g.camera)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	graphics.DrawObject(screen, g.clipper, g.cube)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
