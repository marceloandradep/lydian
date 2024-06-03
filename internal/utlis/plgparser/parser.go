package plgparser

import (
	"bufio"
	"io"
	"lydian/internal/utlis/plgparser/plgpolygon"
	"strconv"
	"strings"
)

type PLGObject struct {
	Header      *PLGHeader
	VertexList  []plgpolygon.PLGVertex
	PolygonList []plgpolygon.PLGPolygon
}

type PLGHeader struct {
	ObjectName  string
	NumVertices int
	NumPolygons int
}

func Parse(r io.ReadCloser) (*PLGObject, error) {
	s := bufio.NewScanner(r)

	header, err := readHeader(s)
	if err != nil {
		return nil, err
	}

	vertexList, err := readVertexList(s, header.NumVertices)
	if err != nil {
		return nil, err
	}

	polygonList, err := readPolygonList(s, header.NumPolygons)
	if err != nil {
		return nil, err
	}

	return &PLGObject{header, vertexList, polygonList}, nil
}

func readHeader(s *bufio.Scanner) (*PLGHeader, error) {
	line := readLine(s)
	if line == nil {
		return nil, HeaderNotFound
	}

	fields := strings.Fields(*line)
	if len(fields) != 3 {
		return nil, InvalidSyntax
	}

	name := fields[0]

	numVertices, err := strconv.Atoi(fields[1])
	if err != nil {
		return nil, InvalidSyntax
	}

	numPolygons, err := strconv.Atoi(fields[2])
	if err != nil {
		return nil, InvalidSyntax
	}

	return &PLGHeader{
		ObjectName:  name,
		NumVertices: numVertices,
		NumPolygons: numPolygons,
	}, nil
}

func readVertexList(s *bufio.Scanner, n int) ([]plgpolygon.PLGVertex, error) {
	vertices := make([]plgpolygon.PLGVertex, 0)

	for i := 0; i < n; i++ {
		line := readLine(s)

		fields := strings.Fields(*line)
		if len(fields) != 3 {
			return nil, InvalidSyntax
		}

		x, err := strconv.ParseFloat(fields[0], 64)
		if err != nil {
			return nil, InvalidSyntax
		}

		y, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			return nil, InvalidSyntax
		}

		z, err := strconv.ParseFloat(fields[2], 64)
		if err != nil {
			return nil, InvalidSyntax
		}

		vertices = append(vertices, plgpolygon.PLGVertex{X: x, Y: y, Z: z})
	}

	return vertices, nil
}

func readPolygonList(s *bufio.Scanner, n int) ([]plgpolygon.PLGPolygon, error) {
	polygons := make([]plgpolygon.PLGPolygon, 0)

	for i := 0; i < n; i++ {
		line := readLine(s)

		fields := strings.Fields(*line)
		if len(fields) != 5 {
			return nil, InvalidSyntax
		}

		surfaceDescriptor := fields[0]

		v0, err := strconv.Atoi(fields[2])
		if err != nil {
			return nil, InvalidSyntax
		}

		v1, err := strconv.Atoi(fields[3])
		if err != nil {
			return nil, InvalidSyntax
		}

		v2, err := strconv.Atoi(fields[4])
		if err != nil {
			return nil, InvalidSyntax
		}

		poly, err := plgpolygon.New(surfaceDescriptor, v0, v1, v2)
		if err != nil {
			return nil, InvalidSyntax
		}

		polygons = append(polygons, *poly)
	}

	return polygons, nil
}

func readLine(s *bufio.Scanner) *string {
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" || line[0] == '#' {
			continue
		}
		return &line
	}
	return nil
}
