package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/go-playground/assert.v1"
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

func TestToHash(t *testing.T) {
	// Arrange
	var (
		f = Facet{
			Vertices: []Vector{
				{X: 0, Y: 1, Z: 2},
				{X: 0, Y: 1, Z: 2},
				{X: 0, Y: 1, Z: 2},
			},
		}
		expected = "0.01.02.0:0.01.02.0:0.01.02.0"
	)

	// Assert
	h := f.toHash()

	// Act
	assert.Equal(t, expected, h)
}

func TestCheckDuplicates(t *testing.T) {
	// Arrange
	tcs := map[string]struct {
		solid    Solid
		expected bool
	}{
		"with duplicate scenario": {
			solid: Solid{
				Facets: []Facet{
					{
						Vertices: []Vector{
							{X: 1, Y: 2, Z: 3},
							{X: 4, Y: 5, Z: 6},
							{X: 7, Y: 8, Z: 9},
						},
					},
					{
						Vertices: []Vector{
							{X: 1, Y: 2, Z: 3},
							{X: 4, Y: 5, Z: 6},
							{X: 7, Y: 8, Z: 9},
						},
					},
				},
			},
			expected: true,
		},
		"without duplicate scenario": {
			solid: Solid{
				Facets: []Facet{
					{
						Vertices: []Vector{
							{X: 1, Y: 2, Z: 3},
							{X: 4, Y: 5, Z: 6},
							{X: 7, Y: 8, Z: 9},
						},
					},
					{
						Vertices: []Vector{
							{X: 1, Y: 1, Z: 3},
							{X: 4, Y: 5, Z: 6},
							{X: 7, Y: 8, Z: 9},
						},
					},
				},
			},
			expected: false,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			// Act
			hasDuplicates := tc.solid.CheckDuplicates()

			// Assert
			assert.Equal(t, tc.expected, hasDuplicates)
		})
	}
}
