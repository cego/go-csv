package csv

import (
	"io"
	"strings"
	"testing"
)

func TestAllTokens(t *testing.T) {
	cases := []struct {
		input      string
		tokenKinds []tokenKind
	}{
		{``, []tokenKind{eof}},
		{`a`, []tokenKind{text, eof}},
		{`,abcd`, []tokenKind{comma, text, text, text, text, eof}},
		{`,"a",b`, []tokenKind{comma, quote, text, quote, comma, text, eof}},
		{`,"",a`, []tokenKind{comma, quote, quote, comma, text, eof}},
		{`,"""ab"`, []tokenKind{comma, quote, text, text, text, quote, eof}},
		{`,"""",a`, []tokenKind{comma, quote, text, quote, comma, text, eof}},
		{`,"""""",a`, []tokenKind{comma, quote, text, text, quote, comma, text, eof}},
		{"a\nb", []tokenKind{text, lineBreak, text, eof}},
		{"a\n,b", []tokenKind{text, lineBreak, comma, text, eof}},
		{`a, b`, []tokenKind{text, comma, text, text, eof}},
	}

	for i := 0; i < len(cases); i++ {
		r := newTokenizer(strings.NewReader(cases[i].input))
		readTokens, err := r.allTokens()
		if err != io.EOF {
			t.Fatalf("case %v expected %v got %v", i, io.EOF, err)
		}

		if len(cases[i].tokenKinds) != len(readTokens) {
			t.Fatalf("case %v expected tokenKinds %#v, got %#v", i, cases[i].tokenKinds, readTokens)
		}

		for j := 0; j < len(cases[i].tokenKinds); j++ {
			if cases[i].tokenKinds[j] != readTokens[j].kind {
				t.Fatalf("case %v expected tokenKinds %#v, got %#v", i, cases[i].tokenKinds, readTokens)
			}
		}
	}
}

func TestSyntaxErrors(t *testing.T) {
	cases := []string{
		`"`,
		`"a""`,
		"\"\n",
	}

	for i := 0; i < len(cases); i++ {
		tr := newTokenizer(strings.NewReader(cases[i]))

		var err error
		for err == nil {
			_, err = tr.readToken()
		}

		if err == nil || err == io.EOF {
			t.Fatalf("case %v expected non-nil, non-eof error, got %v", i, err)
		}
	}
}

func TestReadTokenErrorPassthrough(t *testing.T) {
	tr := newTokenizer(readerStub{})
	_, err := tr.readToken()

	if err != stubError {
		t.Fatalf("expected %v got %v", stubError, err)
	}
}
