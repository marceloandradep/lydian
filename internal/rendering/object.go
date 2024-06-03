package rendering

import (
	"lydian/internal/math"
	"lydian/internal/rendering/camera"
	"lydian/internal/utlis/plgparser"
	"lydian/internal/utlis/plgparser/plgpolygon"
	"os"
)

type Object3D struct {
	culled                bool
	maxRadius, avgRadius  float64
	worldPos              math.Vector
	direction             math.Vector
	ux, uy, uz            math.Vector
	vertexList            []math.Vector
	transformedVertexList []math.Vector
	polygonList           []Polygon3D
}

func (o *Object3D) IsVisible() bool {
	return !o.culled
}

func (o *Object3D) SetWorldPos(x float64, y float64, z float64) {
	o.worldPos.X = x
	o.worldPos.Y = y
	o.worldPos.Z = z
}

func (o *Object3D) GetWorldPos() *math.Vector {
	return &o.worldPos
}

func (o *Object3D) Reset() {
	o.culled = false
	for i := 0; i < len(o.polygonList); i++ {
		if !o.polygonList[i].active {
			continue
		}
		o.polygonList[i].clipped = false
		o.polygonList[i].backface = false
	}
}

func (o *Object3D) Update(cam camera.Camera) {
	o.modelToWorldTransform(TransformedOnly)
	o.cull(cam)
	o.removeBackFaces(cam)
	o.worldToCameraTransform(cam)
	o.cameraToPerspectiveTransform(cam)
	o.perspectiveToScreenTransform(cam)
}

func (o *Object3D) Transform(mt *math.Matrix, transformType TransformType, transformBasis bool) {
	switch transformType {
	case LocalOnly:
		for i := 0; i < len(o.vertexList); i++ {
			o.vertexList[i] = *mt.MultiplyVertex(o.vertexList[i])
		}
	case TransformedOnly:
		for i := 0; i < len(o.transformedVertexList); i++ {
			o.transformedVertexList[i] = *mt.MultiplyVertex(o.transformedVertexList[i])
		}
	case LocalToTransformed:
		for i := 0; i < len(o.vertexList); i++ {
			o.transformedVertexList[i] = *mt.MultiplyVertex(o.vertexList[i])
		}
	}
	if transformBasis {
		o.ux = *mt.MultiplyVertex(o.ux)
		o.uy = *mt.MultiplyVertex(o.uy)
		o.uz = *mt.MultiplyVertex(o.uz)
	}
}

func (o *Object3D) GetPolygonList() []Polygon3D {
	return o.polygonList
}

func (o *Object3D) computeRadius() {
	totalRadius := .0
	for _, vector := range o.vertexList {
		length := vector.Length()
		o.maxRadius = max(o.maxRadius, length)
		totalRadius += length
	}
	o.avgRadius = totalRadius / float64(len(o.vertexList))
}

func (o *Object3D) modelToWorldTransform(transformType TransformType) {
	switch transformType {
	case LocalOnly:
		for i := 0; i < len(o.vertexList); i++ {
			o.vertexList[i] = *o.vertexList[i].Add(o.worldPos)
		}
	case TransformedOnly:
		for i := 0; i < len(o.transformedVertexList); i++ {
			o.transformedVertexList[i] = *o.transformedVertexList[i].Add(o.worldPos)
		}
	case LocalToTransformed:
		for i := 0; i < len(o.vertexList); i++ {
			o.transformedVertexList[i] = *o.vertexList[i].Add(o.worldPos)
		}
	}
}

func (o *Object3D) worldToCameraTransform(cam camera.Camera) {
	if o.culled {
		return
	}
	for i := 0; i < len(o.transformedVertexList); i++ {
		o.transformedVertexList[i] = *cam.WorldToCameraMatrix().MultiplyVertex(o.transformedVertexList[i])
	}
}

func (o *Object3D) cameraToPerspectiveTransform(cam camera.Camera) {
	if o.culled {
		return
	}
	for i := 0; i < len(o.transformedVertexList); i++ {
		o.transformedVertexList[i] = *cam.CameraToPerspectiveMatrix().MultiplyVertex(o.transformedVertexList[i]).DeHomogenize()
	}
}

func (o *Object3D) perspectiveToScreenTransform(cam camera.Camera) {
	if o.culled {
		return
	}
	for i := 0; i < len(o.transformedVertexList); i++ {
		o.transformedVertexList[i] = *cam.PerspectiveToScreenMatrix().MultiplyVertex(o.transformedVertexList[i]).DeHomogenize()
	}
}

func (o *Object3D) cull(cam camera.Camera) {
	sphere := math.Sphere{
		Center: *cam.WorldToCameraMatrix().MultiplyVertex(o.worldPos),
		Radius: o.maxRadius,
	}
	if cam.IsCulled(sphere) {
		o.culled = true
	}
}

func (o *Object3D) removeBackFaces(cam camera.Camera) {
	if o.culled {
		return
	}
	for i := 0; i < len(o.polygonList); i++ {
		polygon := o.polygonList[i]

		if !polygon.shouldTestBackFaceRemoval() {
			continue
		}

		p0 := o.transformedVertexList[polygon.indices[0]]
		p1 := o.transformedVertexList[polygon.indices[1]]
		p2 := o.transformedVertexList[polygon.indices[2]]

		u := p1.Sub(p0)
		v := p2.Sub(p0)

		normal := u.Cross(v)
		view := cam.GetPosition().Sub(p0)

		dot := normal.Dot(view)

		if dot <= 0 {
			polygon.backface = true
		}

		o.polygonList[i] = polygon
	}
}

func Load(filename string, pos, scale, rotation math.Vector) (*Object3D, error) {
	object := &Object3D{
		worldPos:  pos,
		direction: *math.NewVector(0, 0, 1),
		ux:        *math.NewVector(1, 0, 0),
		uy:        *math.NewVector(0, 1, 0),
		uz:        *math.NewVector(0, 0, 1),
	}

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	plg, err := plgparser.Parse(f)
	if err != nil {
		return nil, err
	}

	sclMatrix := math.ScaleMatrix(scale.X, scale.Y, scale.Z)
	rotMatrix := math.RotationSequenceMatrix(rotation.X, rotation.Y, rotation.Z, math.RotationZYX)
	localTransform := sclMatrix.MultiplyMatrix(*rotMatrix)

	object.vertexList = make([]math.Vector, 0)
	object.transformedVertexList = make([]math.Vector, 0)
	for _, vertex := range plg.VertexList {
		vector := localTransform.MultiplyVertex(*math.NewVector(vertex.X, vertex.Y, vertex.Z))
		object.vertexList = append(object.vertexList, *vector)
		object.transformedVertexList = append(object.transformedVertexList, *vector)
	}

	object.computeRadius()

	object.polygonList = make([]Polygon3D, 0)
	for _, plgPolygon := range plg.PolygonList {
		polygon := toPolygon3D(plgPolygon, object.transformedVertexList)
		object.polygonList = append(object.polygonList, *polygon)
	}

	return object, nil
}

func toPolygon3D(plgPolygon plgpolygon.PLGPolygon, vertexList []math.Vector) *Polygon3D {
	var color uint32
	var colorMode ColorMode
	var shadeMode ShadeMode

	if plgPolygon.IsColorModeRGB() {
		colorMode = ColorRBG16
		color = plgPolygon.GetRGB32Color()
	} else {
		colorMode = Color8Bit
		color = uint32(plgPolygon.Get8BitColorIndex())
	}

	switch plgPolygon.ShadeMode() {
	case plgpolygon.Pure:
		shadeMode = ShadeModePure
	case plgpolygon.Flat:
		shadeMode = ShadeModeFlat
	case plgpolygon.Gouraud:
		shadeMode = ShadeModeGouraud
	case plgpolygon.Phong:
		shadeMode = ShadeModePhong
	default:
		shadeMode = ShadeModePure
	}

	return &Polygon3D{
		active:     true,
		is2Sided:   plgPolygon.Is2Sided(),
		colorMode:  colorMode,
		color:      color,
		shadeMode:  shadeMode,
		vertexList: vertexList,
		indices:    [3]int{plgPolygon.V0, plgPolygon.V1, plgPolygon.V2},
	}
}
