package graphics

type position int

const (
	center    position = 0x0
	north     position = 0x8
	south     position = 0x4
	east      position = 0x2
	west      position = 0x1
	northEast position = north | east
	southEast position = south | east
	northWest position = north | west
	southWest position = south | west
)

type Clipper struct {
	MinX, MinY, MaxX, MaxY int
}

func (c *Clipper) ClipLine(x0, y0, x1, y1 int) (int, int, int, int, bool) {
	xc0 := x0
	yc0 := y0
	xc1 := x1
	yc1 := y1

	p0clip := c.getPosition(xc0, yc0)
	p1clip := c.getPosition(xc1, yc1)

	// not visible
	if (p0clip & p1clip) != 0 {
		return xc0, yc0, xc1, yc1, false
	}

	// totally visible
	if p0clip == center && p1clip == center {
		return xc0, yc0, xc1, yc1, true
	}

	xc0, yc0 = c.clipEndpoint(p0clip, xc0, yc0, x0, y0, x1, y1)
	xc1, yc1 = c.clipEndpoint(p1clip, xc1, yc1, x0, y0, x1, y1)

	if !c.checkEndpoints(xc0, yc0, xc1, yc1) {
		return xc0, yc0, xc1, yc1, false
	}

	return xc0, yc0, xc1, yc1, true
}

func (c *Clipper) getPosition(x, y int) position {
	clip := center

	if y < c.MinY {
		clip |= north
	} else if y > c.MaxY {
		clip |= south
	}

	if x < c.MinX {
		clip |= west
	} else if x > c.MaxX {
		clip |= east
	}

	return clip
}

func (c *Clipper) clipEndpoint(p position, endx, endy, x0, y0, x1, y1 int) (int, int) {
	switch p {
	case center:
		return endx, endy
	case north:
		return horizontalIntersection(x0, y0, x1, y1, c.MinY), c.MinY
	case south:
		return horizontalIntersection(x0, y0, x1, y1, c.MaxY), c.MaxY
	case west:
		return c.MinX, verticalIntersection(x0, y0, x1, y1, c.MinX)
	case east:
		return c.MaxX, verticalIntersection(x0, y0, x1, y1, c.MaxX)
	case northEast:
		x := horizontalIntersection(x0, y0, x1, y1, c.MinY)
		if x > c.MaxX {
			return c.MaxX, verticalIntersection(x0, y0, x1, y1, c.MaxX)
		}
		return x, c.MinY
	case southEast:
		x := horizontalIntersection(x0, y0, x1, y1, c.MaxY)
		if x > c.MaxX {
			return c.MaxX, verticalIntersection(x0, y0, x1, y1, c.MaxX)
		}
		return x, c.MaxY
	case northWest:
		x := horizontalIntersection(x0, y0, x1, y1, c.MinY)
		if x < c.MinX {
			return c.MinX, verticalIntersection(x0, y0, x1, y1, c.MinX)
		}
		return x, c.MinY
	case southWest:
		x := horizontalIntersection(x0, y0, x1, y1, c.MaxY)
		if x < c.MinX {
			return c.MinX, verticalIntersection(x0, y0, x1, y1, c.MinX)
		}
		return x, c.MaxY
	default:
		return endx, endy
	}
}

func (c *Clipper) checkEndpoints(x0, y0, x1, y1 int) bool {
	if x0 < c.MinX || x0 > c.MaxX ||
		y0 < c.MinY || y0 > c.MaxY ||
		x1 < c.MinX || x1 > c.MaxX ||
		y1 < c.MinY || y1 > c.MaxY {
		return false
	}
	return true
}

func horizontalIntersection(x0, y0, x1, y1, yHoriz int) int {
	return int(float64(x0) + 0.5 + float64(yHoriz-y0)*float64(x1-x0)/float64(y1-y0))
}

func verticalIntersection(x0, y0, x1, y1, xVert int) int {
	return int(float64(y0) + 0.5 + float64(xVert-x0)*float64(y1-y0)/float64(x1-x0))
}
