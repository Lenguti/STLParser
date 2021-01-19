package lexer

type Token int

const (
	// Anything we cannot parse.
	ILLEGAL Token = iota

	// Special tokens.
	PERIOD
	WS
	NEWLINE
	EOF

	// Keywords.
	SOLID
	ENDSOLID
	FACET
	ENDFACET
	OUTER
	LOOP
	ENDLOOP
	NORMAL
	VERTEX
	WORD
	INTEGER
)
