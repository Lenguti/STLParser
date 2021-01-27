package parser

import (
	"fmt"
	"math"
	"strings"
)

// Solid represents the main object represented by the STL file.
type Solid struct {
	Name   string
	Facets []Facet
}

// SurfaceArea will calculate and return the total surface area of the solid.
func (s Solid) SurfaceArea() float64 {
	var surfaceArea float64
	for i := 0; i < len(s.Facets); i++ {
		surfaceArea += s.Facets[i].Area()
	}
	return surfaceArea
}

/*
	solid{
		Triangles: []Triangle{
			{
				v: a b c
			},
			{
				v: a b c
			},
		},
	}

	hasDuplicates := solid.CheckDuplicates() bool
	if hasDuplicates {
		return error
	}
*/

/*
	Solid{
		Facets: []Facet{
			{
				Vertices: []Vertex{
					{
						{X: 1, y: 1, Z: 1},
						{X: 2, y: 2, Z: 2},
						{X: 3, y: 3 Z: 3},
					}
					"111:222:333"
				},
			},
			{
				Vertices: []Vertex{
					{
						{X: 1, y: 1, Z: 1},
						{X: 1, y: 1, Z: 1},
						{X: 1, y: 1, Z: 1},
					}
				},
			},
		},
	}
*/
func (s Solid) CheckDuplicates() bool {
	duplicatesMap := map[string]bool{}
	for i := 0; i < len(s.Facets); i++ {
		f := s.Facets[i]
		facetHash := f.toHash()
		if _, ok := duplicatesMap[facetHash]; ok {
			return true
		} else {
			duplicatesMap[facetHash] = true
		}
	}
	return false
}

// BoundingBox will calculate and return the min and max vertices
// representing the bounding box of the solid.
func (s Solid) BoundingBox() (Vector, Vector) {
	var (
		minX, minY, minZ = math.Inf(1), math.Inf(1), math.Inf(1)
		maxX, maxY, maxZ = math.Inf(-1), math.Inf(-1), math.Inf(-1)
	)
	for i := 0; i < len(s.Facets); i++ {
		var (
			vs         = s.Facets[i].Vertices
			p1, p2, p3 = vs[0], vs[1], vs[2]
		)
		minX = min(min(minX, p1.X), min(p2.X, p3.X))
		minY = min(min(minY, p1.Y), min(p2.Y, p3.Y))
		minZ = min(min(minZ, p1.Z), min(p3.Z, p3.Z))
		maxX = max(max(maxX, p1.X), max(p2.X, p3.X))
		maxY = max(max(maxY, p1.Y), max(p2.Y, p3.Y))
		maxZ = max(max(maxZ, p1.Z), max(p2.Z, p3.Z))
	}

	return Vector{X: minX, Y: minY, Z: minZ}, Vector{X: maxX, Y: maxY, Z: maxZ}
}

// Facet represents a component of the solid, with a normal and vertices.
// Represented as a triangle.
type Facet struct {
	Normal   Vector
	Vertices []Vector
}

// toHash will convert the facets vertices into a single string
// seperated by ':'.
func (f Facet) toHash() string {
	var hash string
	for i := 0; i < len(f.Vertices); i++ {
		v := f.Vertices[i]
		hash += fmt.Sprintf("%.1f%.1f%.1f:", v.X, v.Y, v.Z)
	}

	return strings.TrimSuffix(hash, ":")
}

// Area will calculate and return the area of the facet.
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

// NewFacet will create a 'Facet' with instantiated vertices of length 3
// representing the three points of a triangle.
func NewFacet() Facet {
	return Facet{
		Vertices: make([]Vector, 3),
	}
}

// Vector represents a point in 3d space with an X, Y, and Z component.
type Vector struct {
	X, Y, Z float64
}

// min will calculate and return the min of two given values.
func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// max will calculate and return the max of two given values.
func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
