// +build unit

package lexer

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// tokenPair represents a datastructure just for testing purposes.
// Expected token with its value.
type tokenPair struct {
	tok Token
	val string
}

func TestScan(t *testing.T) {
	// Arrange
	var (
		r = strings.NewReader(`
solid foo
	facet normal 0 0 0
		outer loop
			vertex 0 0 0
			vertex 1 0 0
			vertex 1 1 1
		endloop
	endfacet
endsolid foo
`)
		expectedParse = []tokenPair{
			{NEWLINE, "\n"},
			{SOLID, "solid"},
			{WS, " "},
			{WORD, "foo"},
			{NEWLINE, "\n"},
			{WS, "\t"},
			{FACET, "facet"},
			{WS, " "},
			{NORMAL, "normal"},
			{WS, " "},
			{INTEGER, "0"},
			{WS, " "},
			{INTEGER, "0"},
			{WS, " "},
			{INTEGER, "0"},
			{NEWLINE, "\n"},
			{WS, "\t\t"},
			{OUTER, "outer"},
			{WS, " "},
			{LOOP, "loop"},
			{NEWLINE, "\n"},
			{WS, "\t\t\t"},
			{VERTEX, "vertex"},
			{WS, " "},
			{INTEGER, "0"},
			{WS, " "},
			{INTEGER, "0"},
			{WS, " "},
			{INTEGER, "0"},
			{NEWLINE, "\n"},
			{WS, "\t\t\t"},
			{VERTEX, "vertex"},
			{WS, " "},
			{INTEGER, "1"},
			{WS, " "},
			{INTEGER, "0"},
			{WS, " "},
			{INTEGER, "0"},
			{NEWLINE, "\n"},
			{WS, "\t\t\t"},
			{VERTEX, "vertex"},
			{WS, " "},
			{INTEGER, "1"},
			{WS, " "},
			{INTEGER, "1"},
			{WS, " "},
			{INTEGER, "1"},
			{NEWLINE, "\n"},
			{WS, "\t\t"},
			{ENDLOOP, "endloop"},
			{NEWLINE, "\n"},
			{WS, "\t"},
			{ENDFACET, "endfacet"},
			{NEWLINE, "\n"},
			{ENDSOLID, "endsolid"},
			{WS, " "},
			{WORD, "foo"},
		}
		s = NewScanner(r)
	)

	// Act
	for _, tp := range expectedParse {
		tok, literal := s.Scan()

		// Assert
		require.Equal(t, tp.tok, tok, "Unexpected token returned [%d]", tok)
		require.Equal(t, tp.val, literal, "Unexpected literal returned [%s]", literal)
	}
}

type mockRuneReader struct {
	valid rune
	size  int
	err   error
}

func (m *mockRuneReader) ReadRune() (rune, int, error) {
	return m.valid, m.size, m.err
}

func (m *mockRuneReader) UnreadRune() error {
	return m.err
}

func TestRead(t *testing.T) {
	// Arrange
	tcs := map[string]struct {
		rr       runeReader
		expected rune
	}{
		"invalid read": {
			rr: &mockRuneReader{
				err: errors.New("unable to read rune"),
			},
			expected: rune(EOF),
		},
		"valid read": {
			rr: &mockRuneReader{
				valid: 'f',
				size:  4,
			},
			expected: 'f',
		},
	}

	// Act
	for testName, testCase := range tcs {
		s := Scanner{
			r: testCase.rr,
			b: new(bytes.Buffer),
		}
		t.Run(testName, func(t *testing.T) {
			out := s.read()

			// Assert
			require.Equal(t, testCase.expected, out)
		})
	}
}

func TestScanWhitespace(t *testing.T) {
	// Arrange
	tcs := map[string]struct {
		input    io.Reader
		expected tokenPair
	}{
		"whitespace read": {
			input: strings.NewReader(`    `),
			expected: tokenPair{
				tok: WS,
				val: "    ",
			},
		},
		"empty scan": {
			input: strings.NewReader(``),
			expected: tokenPair{
				tok: WS,
				val: string(rune(EOF)),
			},
		},
	}

	// Act
	for testName, testCase := range tcs {
		s := NewScanner(testCase.input)
		t.Run(testName, func(t *testing.T) {
			outTok, outVal := s.scanWhitespace()

			// Assert
			require.Equal(t, testCase.expected.tok, outTok)
			require.Equal(t, testCase.expected.val, outVal)
		})
	}
}

func TestScanNumber(t *testing.T) {
	// Arrange
	tcs := map[string]struct {
		input    io.Reader
		expected tokenPair
	}{
		"whole number": {
			input: strings.NewReader(`1234`),
			expected: tokenPair{
				tok: INTEGER,
				val: "1234",
			},
		},
		"floating point": {
			input: strings.NewReader(`0.1234`),
			expected: tokenPair{
				tok: INTEGER,
				val: "0.1234",
			},
		},
	}

	// Act
	for testName, testCase := range tcs {
		s := NewScanner(testCase.input)
		t.Run(testName, func(t *testing.T) {
			outTok, outVal := s.scanNumber()

			// Assert
			require.Equal(t, testCase.expected.tok, outTok)
			require.Equal(t, testCase.expected.val, outVal)
		})
	}
}

func TestScanString(t *testing.T) {
	// Arrange
	tcs := map[string]struct {
		input    io.Reader
		expected tokenPair
	}{
		"keyword token found": {
			input: strings.NewReader(`facet`),
			expected: tokenPair{
				tok: FACET,
				val: "facet",
			},
		},
		"unknow keyword": {
			input: strings.NewReader(`foobar`),
			expected: tokenPair{
				tok: WORD,
				val: "foobar",
			},
		},
	}

	// Act
	for testName, testCase := range tcs {
		s := NewScanner(testCase.input)
		t.Run(testName, func(t *testing.T) {
			outTok, outVal := s.scanString()

			// Assert
			require.Equal(t, testCase.expected.tok, outTok)
			require.Equal(t, testCase.expected.val, outVal)
		})
	}
}
