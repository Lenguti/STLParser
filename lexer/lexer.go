package lexer

import (
	"bufio"
	"bytes"
	"io"
	"log"
)

// runeReader represents behavior for reading and unreading runes.
type runeReader interface {
	ReadRune() (r rune, size int, err error)
	UnreadRune() error
}

// Scanner represents our main "reading" implementation for reading and storing runes.
type Scanner struct {
	r runeReader
	b *bytes.Buffer
}

// NewScanner returns a pointer to a 'Scanner' with an instantiated
// reader and buffer.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		r: bufio.NewReader(r),
		b: new(bytes.Buffer),
	}
}

// Scan will read from our input and return the corresponding token and string value.
func (s *Scanner) Scan() (Token, string) {
	// Read the next rune.
	r := s.read()

	// If we catch any whitespace then consume it until we find next token.
	if isWhitespace(r) {
		s.unread()
		return s.scanWhitespace()
	}

	// If we catch any letter then consume as word or keyword.
	if isLetter(r) {
		s.unread()
		return s.scanString()
	}

	// If we catch any integer then consume as integer or number (floating point values).
	if isInteger(r) {
		s.unread()
		return s.scanNumber()
	}

	// If we reach this point consume rune for what it is.
	var (
		tok Token
		val string
	)
	switch r {
	case 0:
		tok = EOF
	case '\n':
		tok = NEWLINE
		val = string(r)
	default:
		val = string(r)
	}

	return tok, val
}

func (s *Scanner) read() rune {
	r, _, err := s.r.ReadRune()
	if err != nil {
		log.Printf("lexer: error reading rune [%s]", err)
		return rune(EOF)
	}
	return r
}

func (s *Scanner) unread() {
	_ = s.r.UnreadRune() // Ignore error as we know we will ever only read rune.
}

// scanWhitespace will consume any whitespace or tab rune until
// a different rune is found and return 'WS' token and its value.
func (s *Scanner) scanWhitespace() (Token, string) {
	s.clearBuffer()
	s.b.WriteRune(s.read())

	for {
		r := s.read()
		if isWhitespace(r) {
			s.b.WriteRune(r)
		} else {
			if r != rune(EOF) {
				s.unread()
			}
			break
		}
	}

	return WS, s.b.String()
}

// scanInteger will consume any integer rune until
// a different rune is found and return 'INTEGER' token and its value.
func (s *Scanner) scanInteger() (Token, string) {
	s.clearBuffer()
	s.b.WriteRune(s.read())

	for {
		r := s.read()
		if isInteger(r) {
			s.b.WriteRune(r)
		} else {
			if r != rune(EOF) {
				s.unread()
			}
			break
		}
	}

	return INTEGER, s.b.String()
}

// scanNumber will consume any interger rune until a different rune other than
// 'PERIOD' is found. If 'PERIOD' is found we know we have encounterd a floating point
// number and we will continue to parse integers after the found 'PERIOD'.
func (s *Scanner) scanNumber() (Token, string) {
	var (
		tok                   Token
		val                   string
		firstHalf, secondHalf string
	)
	tok, firstHalf = s.scanInteger()
	r := s.read()
	if isPeriod(r) {
		tok, secondHalf = s.scanInteger()
	} else {
		s.unread()
	}

	val = firstHalf
	if secondHalf != "" {
		val += "." + secondHalf
	}

	return tok, val
}

// scanString will consume any character rune until
// a different rune is found. If the value of the string
// is a keyword token then we return said token and value
// otherwise we return 'WORD' token.
func (s *Scanner) scanString() (Token, string) {
	s.clearBuffer()
	s.b.WriteRune(s.read())

	for {
		r := s.read()
		if isLetter(r) {
			s.b.WriteRune(r)
		} else {
			if r != rune(EOF) {
				s.unread()
			}
			break
		}
	}

	var (
		tok = WORD
		val = s.b.String()
	)
	switch val {
	case "solid":
		tok = SOLID
	case "facet":
		tok = FACET
	case "normal":
		tok = NORMAL
	case "outer":
		tok = OUTER
	case "loop":
		tok = LOOP
	case "vertex":
		tok = VERTEX
	case "endloop":
		tok = ENDLOOP
	case "endfacet":
		tok = ENDFACET
	case "endsolid":
		tok = ENDSOLID
	}

	return tok, val
}

func (s *Scanner) clearBuffer() {
	if s.b.Len() != 0 {
		s.b.Reset()
	}
}

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\t'
}

func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isInteger(r rune) bool {
	return r >= '0' && r <= '9'
}

func isPeriod(r rune) bool {
	return r == '.'
}
