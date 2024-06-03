package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"lydian/internal/geometry"
	"lydian/internal/graphics"
	"lydian/internal/math"
	"lydian/internal/rendering"
	"lydian/internal/rendering/camera"
)

type GameUVN struct {
	cube       rendering.Object3D
	cubePos    *math.Vector
	cubeRot    *math.Vector
	camera     camera.Camera
	cameraPos  *math.Vector
	elevation  float64
	heading    float64
	lastMouseX int
	lastMouseY int
	clipper    graphics.Clipper
}

func (g *GameUVN) Init() error {
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

	viewPortSize := geometry.Dimension{
		Width:  screenWidth - 1,
		Height: screenHeight - 1,
	}

	uvn := camera.NewUVN(*g.cameraPos, *math.NewVector(0, 0, 0), 50, 500, 90, viewPortSize)
	g.camera = uvn

	g.cubePos = math.NewVector(0, 0, 100)
	g.cubeRot = math.NewVector(0, 0, 0)

	return nil
}

func (g *GameUVN) Update() error {
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

	g.elevation -= float64(dy)
	g.heading -= float64(dx)

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.cameraPos.Z += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.cameraPos.Z -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.cameraPos.X -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.cameraPos.X += 1
	}

	g.cube.Reset()
	g.cube.SetWorldPos(g.cubePos.X, g.cubePos.Y, g.cubePos.Z)

	g.cubeRot.X += 1
	g.cubeRot.Y += 1
	g.cubeRot.Z += 1

	// rotation := math.RotationSequenceMatrix(g.cubeRot.X, g.cubeRot.Y, g.cubeRot.Z, math.RotationZYX)
	rotation := math.IdentityMatrix()
	g.cube.Transform(rotation, rendering.LocalToTransformed, true)

	g.camera.SetPosition(g.cameraPos.X, g.cameraPos.Y, g.cameraPos.Z)
	g.camera.(*camera.UVNCamera).SetPolarCoordinates(g.elevation, g.heading)
	// g.camera.(*camera.UVNCamera).SetTarget(*g.cube.GetWorldPos())
	g.camera.Update()
	g.cube.Update(g.camera)

	return nil
}

func (g *GameUVN) Draw(screen *ebiten.Image) {
	graphics.DrawObject(screen, g.clipper, g.cube)
}

func (g *GameUVN) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
