package parser

import (
	"math"
)

type Solid struct {
	Name   string
	Facets []Facet
}

func (s Solid) SurfaceArea() float64 {
	var surfaceArea float64
	for i := 0; i < len(s.Facets); i++ {
		surfaceArea += s.Facets[i].Area()
	}
	return surfaceArea
}

type Facet struct {
	Normal   Vector
	Vertices []Vector
}

func (f Facet) Area() float64 {
	if len(f.Vertices) != 3 {
		return 0
	}

	triangle := struct {
		v1, v2, v3 Vector
	}{
		v1: f.Vertices[0],
		v2: f.Vertices[1],
		v3: f.Vertices[2],
	}

	// Subtract triangle vertices v2-v1 and v3-v1.
	rv1 := Vector{X: triangle.v2.X - triangle.v1.X, Y: triangle.v2.Y - triangle.v1.Y, Z: triangle.v2.Z - triangle.v1.Z}
	rv2 := Vector{X: triangle.v3.X - triangle.v1.X, Y: triangle.v3.Y - triangle.v1.Y, Z: triangle.v3.Z - triangle.v1.Z}

	// Take cross product	of resulting two vectors.
	// (rv1.y*rv2.z) - (rv1.z*rv2.y)
	// (rv1.z*rv2.x) - (rv1.x*rv2.z)
	// (rv1.x*rv2.y) - (rv1.y*rv2.x)
	cp := Vector{X: (rv1.Y * rv2.Z) - (rv1.Z * rv2.Y), Y: (rv1.Z * rv2.X) - (rv1.X * rv2.Z), Z: (rv1.X * rv2.Y) - (rv1.Y * rv2.X)}

	// Calculate Area.
	return .5 * math.Sqrt(cp.X*cp.X+cp.Y*cp.Y+cp.Z*cp.Z)
}

func NewFacet() Facet {
	return Facet{
		Vertices: make([]Vector, 3),
	}
}

type Vector struct {
	X, Y, Z float64
}
