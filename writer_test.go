package csv

import (
	"strings"
	"testing"
)

func TestWriteAll(t *testing.T) {
	cases := []struct {
		input  [][]string
		output string
	}{
		{[][]string{}, ""},
		{[][]string{{}}, "\r\n"},
		{[][]string{{``}}, "\"\"\r\n"},
		{[][]string{{`a`, `b`}}, "\"a\",\"b\"\r\n"},
		{[][]string{{`a`, `b`}, {`c`}}, "\"a\",\"b\"\r\n\"c\"\r\n"},
		{[][]string{{`a`, `b`}, {}}, "\"a\",\"b\"\r\n\r\n"},
		{[][]string{{`"`}}, "\"\"\"\"\r\n"},
		{[][]string{{`"a"`}}, "\"\"\"a\"\"\"\r\n"},
		{[][]string{{"\n"}}, "\"\n\"\r\n"},
		{[][]string{{"\r"}}, "\"\r\"\r\n"},
	}

	for i := 0; i < len(cases); i++ {

		b := strings.Builder{}
		w := NewWriter(&b)

		err := w.WriteAll(cases[i].input)
		if err != nil {
			t.Fatal(err)
		}

		if b.String() != cases[i].output {
			t.Fatalf("case %v expected output %#v, got %#v", i, cases[i].output, b.String())
		}
	}
}

func TestWriteAllErrorPassthrough(t *testing.T) {
	w := NewWriter(writerStub{})
	err := w.WriteAll([][]string{{""}})

	if err != stubError {
		t.Fatalf("expected %v got %v", stubError, err)
	}
}
