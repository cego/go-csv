package csv

import (
	"fmt"
	"io"
)

// tokenizer is used to read CSV tokens from a source
type tokenizer struct {
	source        io.Reader
	buffer        []byte
	betweenQuotes bool
	onFirstLine   bool
}

// newTokenizer returns a tokenizer ready to start reading CSV formatted data from the given source
func newTokenizer(source io.Reader) tokenizer {
	return tokenizer{
		source:      source,
		buffer:      make([]byte, 0, 3),
		onFirstLine: true,
	}
}

// readToken reads and returns the next token from the source
// when readToken returns an error, the token can be ignored (the error shouldn't be)
func (t *tokenizer) readToken() (token, error) {
	err := t.fillBuffer()
	if err != nil && err != io.EOF {
		return token{}, err
	}

	if len(t.buffer) == 0 && t.betweenQuotes {
		return token{}, fmt.Errorf("syntax error, reached EOF while scanning quoted field")
	}

	// eof
	if len(t.buffer) == 0 {
		return token{eof, ``}, io.EOF
	}

	// escaped quotes (text)
	if t.betweenQuotes && t.peekBuffer(2) == `""` {
		t.popBuffer(2)
		return token{text, `"`}, nil
	}

	// quote
	if !t.betweenQuotes && t.peekBuffer(1) == `"` {
		t.betweenQuotes = true
		return token{quote, t.popBuffer(1)}, nil
	}
	if t.betweenQuotes && t.peekBuffer(1) == `"` {
		t.betweenQuotes = false
		return token{quote, t.popBuffer(1)}, nil
	}

	// comma
	if !t.betweenQuotes && t.peekBuffer(1) == `,` {
		return token{comma, t.popBuffer(1)}, nil
	}

	// linebreak
	if !t.betweenQuotes && t.peekBuffer(2) == "\r\n" {
		t.onFirstLine = false
		return token{lineBreak, t.popBuffer(2)}, nil
	}
	if !t.betweenQuotes && t.peekBuffer(1) == "\n" {
		t.onFirstLine = false
		return token{lineBreak, t.popBuffer(1)}, nil
	}

	// text
	return token{text, t.popBuffer(1)}, nil
}

// allTokens reads and returns all unread tokens from the source
func (t *tokenizer) allTokens() ([]token, error) {
	var (
		err    error
		tkn    token
		result []token
	)

	for err == nil {
		tkn, err = t.readToken()
		result = append(result, tkn)
	}

	return result, err
}

// fillBuffer fills the internal buffer to capacity by reading from the source
func (t *tokenizer) fillBuffer() error {
	n, err := t.source.Read(t.buffer[len(t.buffer):cap(t.buffer)])
	t.buffer = t.buffer[:len(t.buffer)+n]
	return err
}

// peekBuffer returns the first (at most n) bytes of the internal buffer as a string without modifiying the buffer
func (t *tokenizer) peekBuffer(n int) string {
	l := n
	if l > len(t.buffer) {
		l = len(t.buffer)
	}

	return string(t.buffer[:l])
}

// popBuffer removes the first (at most n) bytes of the internal buffer and then return them as a string
func (t *tokenizer) popBuffer(n int) string {
	pop := t.peekBuffer(n)
	l := len(pop)

	copy(t.buffer[:], t.buffer[l:])
	t.buffer = t.buffer[:len(t.buffer)-l]

	return pop
}
