package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSurfaceArea(t *testing.T) {
	// Arrange
	var (
		s = Solid{
			Facets: []Facet{
				{
					Vertices: []Vector{
						{
							X: 1, Y: 2, Z: 3,
						},
						{
							X: -11, Y: -3, Z: -6,
						},
						{
							X: 5, Y: 2, Z: 8,
						},
					},
				},
				{
					Vertices: []Vector{
						{
							X: 4, Y: 4, Z: 4,
						},
						{
							X: 1, Y: -3, Z: 6,
						},
						{
							X: 7, Y: 2, Z: -8,
						},
					},
				},
			},
		}
		expected float64 = 68.4133765980988
	)

	// Act
	out := s.SurfaceArea()

	// Assert
	require.Equal(t, expected, out)
}

func TestBoundingBox(t *testing.T) {
	// Arrange
	var (
		s = Solid{
			Facets: []Facet{
				{
					Vertices: []Vector{
						{
							X: 1, Y: 2, Z: 3,
						},
						{
							X: -11, Y: -3, Z: -6,
						},
						{
							X: 5, Y: 2, Z: 8,
						},
					},
				},
				{
					Vertices: []Vector{
						{
							X: 4, Y: 4, Z: 4,
						},
						{
							X: 1, Y: -3, Z: 6,
						},
						{
							X: 7, Y: 2, Z: -8,
						},
					},
				},
			},
		}
		expectedMin = Vector{X: -11, Y: -3, Z: -8}
		expectedMax = Vector{X: 7, Y: 4, Z: 8}
	)

	// Act
	outMin, outMax := s.BoundingBox()

	// Assert
	require.Equal(t, expectedMin, outMin)
	require.Equal(t, expectedMax, outMax)
}

func TestArea(t *testing.T) {
	// Arrange
	var (
		f = Facet{
			Vertices: []Vector{
				{
					X: 1, Y: 2, Z: 3,
				},
				{
					X: -11, Y: -3, Z: -6,
				},
				{
					X: 5, Y: 2, Z: 8,
				},
			},
		}
		expected float64 = 20.006249023742555
	)

	// Act
	out := f.Area()

	// Assert
	require.Equal(t, expected, out)
}
