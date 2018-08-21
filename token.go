package csv

type token struct {
	kind  tokenKind
	value string
}

type tokenKind int

const (
	quote     tokenKind = 1
	comma     tokenKind = 2
	lineBreak tokenKind = 3
	text      tokenKind = 4
	eof       tokenKind = 5
)
