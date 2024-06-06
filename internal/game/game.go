package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"lydian/internal/camera"
	"lydian/internal/geometry"
	"lydian/internal/graphics"
	"lydian/internal/math"
	"lydian/internal/rendering"
)

type Game struct {
	ScreenWidth  int
	ScreenHeight int

	cameraPos    *math.Vector
	cameraRot    *math.Vector
	camera       *camera.Camera
	forward      *math.Vector
	left         *math.Vector
	upward       *math.Vector
	lastMouseX   int
	lastMouseY   int
	triangleList []*rendering.Triangle3D
	clipper      graphics.Clipper
	cubes        []*rendering.Object3D
	scene        *rendering.Scene
}

func (g *Game) Init() error {
	g.clipper = graphics.Clipper{MinX: 0, MinY: 0, MaxX: g.ScreenWidth - 1, MaxY: g.ScreenHeight - 1}

	g.cameraPos = math.NewVector3(0, 0, 0)
	g.cameraRot = math.NewVector3(0, 0, 0)
	g.forward = math.NewVector3(0, 0, 1)
	g.left = math.NewVector3(-1, 0, 0)
	g.upward = math.NewVector3(0, 1, 0)

	viewPortSize := &geometry.Dimension{
		Width:  float64(g.ScreenWidth - 1),
		Height: float64(g.ScreenHeight - 1),
	}

	cam := camera.NewEuler(g.cameraPos, g.cameraRot, math.RotationZYX, 0, 1000, 90, viewPortSize)
	// cam := camera.NewUVN(g.cameraPos, math.NewVector3(0, 0, 1), 0, 500, 90, viewPortSize)
	g.camera = cam

	g.triangleList = createMap(1000)

	cubes, err := Cubes(5)
	if err != nil {
		return err
	}
	g.cubes = cubes

	g.scene = rendering.NewScene()
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
	g.upward = rot.MultiplyVertex(g.upward)

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

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.cameraPos = g.cameraPos.Add(g.upward)
	}

	if ebiten.IsKeyPressed(ebiten.KeyControlLeft) {
		g.cameraPos = g.cameraPos.Sub(g.upward)
	}

	g.camera.SetPos(g.cameraPos)
	g.camera.Rotate(g.cameraRot)
	// g.camera.SetPolarCoordinates(g.cameraRot.X, g.cameraRot.Y)

	g.camera.Clear()
	g.camera.Update()

	for _, t := range g.triangleList {
		g.camera.AddTriangleToCamera(t)
	}

	g.scene.Clear()

	for _, cube := range g.cubes {
		g.scene.AddToScene(cube)
	}

	g.camera.AddSceneToCamera(g.scene)
	g.camera.ProjectTriangles()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	graphics.Rasterize(screen, g.clipper, g.camera)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}
