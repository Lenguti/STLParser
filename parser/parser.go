package parser

import (
	"io"
	"strconv"

	"github.com/lenguti/STLParser/lexer"
	"github.com/pkg/errors"
)

type scanner interface {
	Scan() (lexer.Token, string)
}

// Parser represnts our main parsing object for reading contents of an STL file.
type Parser struct {
	s scanner
	b struct {
		tok lexer.Token // Last read token.
		val string      // Last read value.
		n   int         // Buffer size, max of 1.
	}
}

// New Returns a pointer to a 'Parser' with an
// instantiated 'lexer.Scanner'.
func New(r io.Reader) *Parser {
	return &Parser{
		s: lexer.NewScanner(r),
	}
}

// parseVector will parse a vector of the form:
// '0 0 0' representing the X, Y, and Z of a vector.
// Will return an error if more than 3 points are found.
func (p *Parser) parseVector() (Vector, error) {
	var (
		v      Vector
		i      int
		points = make([]float64, 3)
	)
	// Since we know we are parsing a vector, loop through x, y, and z values until we reach end of line.
	for tok, val := p.scanIgnoreWhitespace(); tok != lexer.NEWLINE; tok, val = p.scanIgnoreWhitespace() {
		if tok != lexer.INTEGER {
			return v, errors.Errorf("parse vector: found [%v], expected 'integer'", val)
		}

		// If we recieve more than three points then this is an invalid vector.
		if i > 2 {
			return v, errors.New("parse vector: too many points in vector")
		}

		floatVal, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return v, errors.WithMessage(err, "parse vector: unable to parse into float")
		}
		points[i] = floatVal
		i++
	}
	p.unscan() // Put back last read token and value into buffer to be read again.
	v.X = points[0]
	v.Y = points[1]
	v.Z = points[2]
	return v, nil
}

// parseVertices will parse vertices of the form:
// 'vertex 0 0 0' representing the 3 vertices of a triangle.
// Will return an error if more than 3 vertices are found.
func (p *Parser) parseVertices() ([]Vector, error) {
	var (
		i  int
		vs = make([]Vector, 3)
	)
	// Since we know we are parsing a the vertices of a facet, loop through each vertice until we reach end loop.
	for tok, val := p.scanIgnoreWhitespace(); tok != lexer.ENDLOOP; tok, val = p.scanIgnoreWhitespace() {
		if tok != lexer.VERTEX {
			return vs, errors.Errorf("parse vertices: found [%v], expected 'vertex'", val)
		}

		// If we recieve more than three vertices then we know this is an invalid shape.
		if i > 2 {
			return vs, errors.New("parse verticies: too many vertices in facet")
		}

		v, err := p.parseVector()
		if err != nil {
			return vs, errors.WithMessage(err, "parse vertices: unable to parse vector")
		}
		vs[i] = v
		i++

		tok, val = p.scanIgnoreWhitespace()
		if tok != lexer.NEWLINE {
			return vs, errors.Errorf("parse verticies: found [%v], expected 'newline'", val)
		}
	}
	p.unscan() // Put back last read token and value into buffer to be read again.
	return vs, nil
}

// parseFacet will parse a facet of the form:
// 'facet normal 0 0 0
//   outer loop
//     vertex 0 0 0
//     vertex 0 1 1
//     vertex 1 1 1
//   endloop
// endfacet'
// Representing a triangle and its given normal and vertices.
func (p *Parser) parseFacet() (Facet, error) {
	f := NewFacet()
	tok, val := p.scanIgnoreWhitespace()
	if tok != lexer.FACET {
		return f, errors.Errorf("parse facet: found [%v], expected 'facet'", val)
	}

	tok, val = p.scanIgnoreWhitespace()
	if tok != lexer.NORMAL {
		return f, errors.Errorf("parse facet: found [%v], expected 'normal'", val)
	}

	normal, err := p.parseVector()
	if err != nil {
		return f, errors.WithMessage(err, "parse facet: unable to parse normal vector")
	}
	f.Normal = normal

	tok, val = p.scanIgnoreWhitespace()
	if tok != lexer.NEWLINE {
		return f, errors.Errorf("parse facet: found [%v], expected 'newline'", val)
	}

	tok, val = p.scanIgnoreWhitespace()
	if tok != lexer.OUTER {
		return f, errors.Errorf("parse facet: found [%v], expected 'outer'", val)
	}

	tok, val = p.scanIgnoreWhitespace()
	if tok != lexer.LOOP {
		return f, errors.Errorf("parse facet: found [%v], expected 'loop'", val)
	}

	tok, val = p.scanIgnoreWhitespace()
	if tok != lexer.NEWLINE {
		return f, errors.Errorf("parse facet: found [%v], expected 'newline'", val)
	}

	vertices, err := p.parseVertices()
	if err != nil {
		return f, errors.WithMessage(err, "parse facet: unable to parse vertices")
	}
	f.Vertices = vertices

	tok, val = p.scanIgnoreWhitespace()
	if tok != lexer.ENDLOOP {
		return f, errors.Errorf("parse facet: found [%v], expected 'endloop'", val)
	}

	tok, val = p.scanIgnoreWhitespace()
	if tok != lexer.NEWLINE {
		return f, errors.Errorf("parse facet: found [%v], expected 'newline'", val)
	}

	tok, val = p.scanIgnoreWhitespace()
	if tok != lexer.ENDFACET {
		return f, errors.Errorf("parse facet: found [%v], expected 'endfacet'", val)
	}

	return f, nil
}

// Parse is our main entry point into parsing an STL file.
// Parse will read and evaluate tokens to build and determine
// a proper STL object.
func (p *Parser) Parse() (Solid, error) {
	var s Solid
	tok, val := p.scanIgnoreWhitespace()
	if tok != lexer.SOLID {
		return s, errors.Errorf("parse: found [%v], expected 'solid'", val)
	}

	tok, val = p.scanIgnoreWhitespace()
	if tok != lexer.WORD {
		return s, errors.Errorf("parse: found [%v], expected name of solid", val)
	}
	s.Name = val

	tok, val = p.scanIgnoreWhitespace()
	if tok != lexer.NEWLINE {
		return s, errors.Errorf("parse: found [%v], expected 'newline'", val)
	}

	for {
		f, err := p.parseFacet()
		if err != nil {
			return s, errors.WithMessage(err, "parse: unable to parse facet")
		}
		s.Facets = append(s.Facets, f)

		tok, val = p.scanIgnoreWhitespace()
		if tok != lexer.NEWLINE {
			return s, errors.Errorf("parse: found [%v], expected 'newline'", val)
		}

		tok, val = p.scanIgnoreWhitespace()
		if tok == lexer.ENDSOLID {
			break
		} else {
			// If current token is not 'ENDSOLID' it should be 'FACET' in which case we must put it back
			// on the buffer so the next facet can be parsed.
			p.unscan()
		}
	}

	tok, val = p.scanIgnoreWhitespace()
	if tok != lexer.WORD {
		return s, errors.Errorf("parse: found [%v], expected name of solid", val)
	}

	if s.Name != val {
		return s, errors.Errorf("parse: solid names do not match [%s] and [%s]", s.Name, val)
	}

	return s, nil
}

// scan will read off the buffer, if buffer is currently empty then we read from the underlying scanner.
func (p *Parser) scan() (lexer.Token, string) {
	// If we have a token on the buffer, then return it.
	if p.b.n != 0 {
		p.b.n = 0
		return p.b.tok, p.b.val
	}
	p.b.tok, p.b.val = p.s.Scan()
	return p.b.tok, p.b.val
}

func (p *Parser) unscan() {
	p.b.n = 1
}

// scanIgnoreWhitespace will return the token and value of the next non whitespace token.
func (p *Parser) scanIgnoreWhitespace() (lexer.Token, string) {
	tok, val := p.scan()
	if tok == lexer.WS {
		tok, val = p.scan()
	}
	return tok, val
}
