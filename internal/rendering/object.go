package rendering

import (
	"lydian/internal/math"
	"lydian/internal/utlis/plgparser"
	"lydian/internal/utlis/plgparser/plgpolygon"
	"os"
)

type TransformType int

const (
	LocalOnly TransformType = iota
	TransformedOnly
	LocalToTransformed
)

type Object3D struct {
	maxRadius, avgRadius  float64
	worldPos              *math.Vector
	direction             *math.Vector
	ux, uy, uz            *math.Vector
	vertexList            []*math.Vector
	transformedVertexList []*math.Vector
	triangleList          []*Triangle3D
}

func (o *Object3D) Copy(pos, scale, rotation *math.Vector) *Object3D {
	copied := &Object3D{
		maxRadius: o.maxRadius,
		avgRadius: o.avgRadius,
		worldPos:  pos,
		direction: o.direction,
		ux:        o.ux,
		uy:        o.uy,
		uz:        o.uz,
	}

	sclMatrix := math.ScaleMatrix(scale.X, scale.Y, scale.Z)
	rotMatrix := math.RotationSequenceMatrix(rotation.X, rotation.Y, rotation.Z, math.RotationZYX)
	localTransform := sclMatrix.MultiplyMatrix(rotMatrix)

	copied.vertexList = make([]*math.Vector, 0)
	copied.transformedVertexList = make([]*math.Vector, 0)
	for _, vertex := range o.vertexList {
		vector := localTransform.MultiplyVertex(math.NewVector3(vertex.X, vertex.Y, vertex.Z))
		copied.vertexList = append(copied.vertexList, vector)
		copied.transformedVertexList = append(copied.transformedVertexList, vector.Copy())
	}

	copied.triangleList = make([]*Triangle3D, 0)
	for _, triangle := range o.triangleList {
		indices := triangle.Indices
		copiedTriangle := NewTriangle3D(copied.transformedVertexList, indices[0], indices[1], indices[2], triangle.Is2Sided, triangle.Color)
		copied.triangleList = append(copied.triangleList, copiedTriangle)
	}

	return copied
}

func (o *Object3D) SetWorldPos(x float64, y float64, z float64) {
	o.worldPos.X = x
	o.worldPos.Y = y
	o.worldPos.Z = z
}

func (o *Object3D) GetWorldPos() *math.Vector {
	return o.worldPos
}

func (o *Object3D) Transform(mt *math.Matrix, transformType TransformType, transformBasis bool) {
	switch transformType {
	case LocalOnly:
		for i := 0; i < len(o.vertexList); i++ {
			o.vertexList[i].Set(mt.MultiplyVertex(o.vertexList[i]))
		}
	case TransformedOnly:
		for i := 0; i < len(o.transformedVertexList); i++ {
			o.transformedVertexList[i].Set(mt.MultiplyVertex(o.transformedVertexList[i]))
		}
	case LocalToTransformed:
		for i := 0; i < len(o.vertexList); i++ {
			o.transformedVertexList[i].Set(mt.MultiplyVertex(o.vertexList[i]))
		}
	}
	if transformBasis {
		o.ux.Set(mt.MultiplyVertex(o.ux))
		o.uy.Set(mt.MultiplyVertex(o.uy))
		o.uz.Set(mt.MultiplyVertex(o.uz))
	}
}

func (o *Object3D) TriangleList() []*Triangle3D {
	return o.triangleList
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

func Load(filename string, pos, scale, rotation *math.Vector) (*Object3D, error) {
	object := &Object3D{
		worldPos:  pos,
		direction: math.NewVector3(0, 0, 1),
		ux:        math.NewVector3(1, 0, 0),
		uy:        math.NewVector3(0, 1, 0),
		uz:        math.NewVector3(0, 0, 1),
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
	localTransform := sclMatrix.MultiplyMatrix(rotMatrix)

	object.vertexList = make([]*math.Vector, 0)
	object.transformedVertexList = make([]*math.Vector, 0)
	for _, vertex := range plg.VertexList {
		vector := localTransform.MultiplyVertex(math.NewVector3(vertex.X, vertex.Y, vertex.Z))
		object.vertexList = append(object.vertexList, vector)
		object.transformedVertexList = append(object.transformedVertexList, vector.Copy())
	}

	object.computeRadius()

	object.triangleList = make([]*Triangle3D, 0)
	for _, plgPolygon := range plg.PolygonList {
		triangle := toTriangle3D(plgPolygon, object.transformedVertexList)
		object.triangleList = append(object.triangleList, triangle)
	}

	return object, nil
}

func toTriangle3D(plgPolygon *plgpolygon.PLGPolygon, vertexList []*math.Vector) *Triangle3D {
	return NewTriangle3D(vertexList, plgPolygon.V0, plgPolygon.V1, plgPolygon.V2, plgPolygon.Is2Sided(), plgPolygon.GetRGB32Color())
}
