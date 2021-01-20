// +build unit

package parser

import (
	"testing"

	"github.com/lenguti/STLParser/lexer"
	"github.com/stretchr/testify/require"
)

type mockScanner struct {
	tokens []struct {
		t lexer.Token
		v string
	}
}

func (m *mockScanner) Scan() (lexer.Token, string) {
	if len(m.tokens) == 0 {
		return lexer.EOF, ""
	}

	// Pop last element of tokens.
	temp := m.tokens[len(m.tokens)-1]
	// Resize tokens with one less.
	m.tokens = m.tokens[:len(m.tokens)-1]
	return temp.t, temp.v
}

func TestParse(t *testing.T) {
	// Arrange
	var (
		p = &Parser{
			s: &mockScanner{
				tokens: []struct {
					t lexer.Token
					v string
				}{
					{
						t: lexer.WORD,
						v: "foo",
					},
					{
						t: lexer.ENDSOLID,
						v: "endsolid",
					},
					{
						t: lexer.NEWLINE,
						v: "\n",
					},
					{
						t: lexer.ENDFACET,
						v: "endfacet",
					},
					{
						t: lexer.NEWLINE,
						v: "\n",
					},
					{
						t: lexer.ENDLOOP,
						v: "endloop",
					},
					{
						t: lexer.NEWLINE,
						v: "\n",
					},
					{
						t: lexer.INTEGER,
						v: "0",
					},
					{
						t: lexer.INTEGER,
						v: "0",
					},
					{
						t: lexer.INTEGER,
						v: "0",
					},
					{
						t: lexer.VERTEX,
						v: "vertex",
					},
					{
						t: lexer.NEWLINE,
						v: "\n",
					},
					{
						t: lexer.INTEGER,
						v: "0",
					},
					{
						t: lexer.INTEGER,
						v: "0",
					},
					{
						t: lexer.INTEGER,
						v: "0",
					},
					{
						t: lexer.VERTEX,
						v: "vertex",
					},
					{
						t: lexer.NEWLINE,
						v: "\n",
					},
					{
						t: lexer.INTEGER,
						v: "0",
					},
					{
						t: lexer.INTEGER,
						v: "0",
					},
					{
						t: lexer.INTEGER,
						v: "0",
					},
					{
						t: lexer.VERTEX,
						v: "vertex",
					},
					{
						t: lexer.NEWLINE,
						v: "\n",
					},
					{
						t: lexer.LOOP,
						v: "loop",
					},
					{
						t: lexer.OUTER,
						v: "outer",
					},
					{
						t: lexer.NEWLINE,
						v: "\n",
					},
					{
						t: lexer.INTEGER,
						v: "0",
					},
					{
						t: lexer.INTEGER,
						v: "0",
					},
					{
						t: lexer.INTEGER,
						v: "0",
					},
					{
						t: lexer.NORMAL,
						v: "normal",
					},
					{
						t: lexer.FACET,
						v: "facet",
					},
					{
						t: lexer.NEWLINE,
						v: "\n",
					},
					{
						t: lexer.WORD,
						v: "foo",
					},
					{
						t: lexer.SOLID,
						v: "solid",
					},
				},
			},
		}
		expected = Solid{
			Name: "foo",
			Facets: []Facet{
				{
					Normal: Vector{X: 0, Y: 0, Z: 0},
					Vertices: []Vector{
						{X: 0, Y: 0, Z: 0},
						{X: 0, Y: 0, Z: 0},
						{X: 0, Y: 0, Z: 0},
					},
				},
			},
		}
	)

	// Act
	s, err := p.Parse()

	// Assert
	require.NoError(t, err)
	require.Equal(t, expected, s)
}
