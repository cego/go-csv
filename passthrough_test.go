package csv

import (
	"io"
	"strings"
	"testing"
)

// TestWriterToReader tests that feeding records to a Writer and then into a Reader should result in unchanged records
func TestWriterToReader(t *testing.T) {
	cases := [][][]string{
		{{""}},
		{{""}, {""}},
		{{"foo"}, {"bar", "baz"}},
	}

	for i := 0; i < len(cases); i++ {
		b := strings.Builder{}
		w := NewWriter(&b)

		err := w.WriteAll(cases[i])
		if err != nil {
			t.Fatal(err)
		}

		err = w.Flush()
		if err != nil {
			t.Fatal(err)
		}

		r := NewReader(strings.NewReader(b.String()))
		records, err := r.ReadAll()
		if err != io.EOF {
			t.Fatalf("case %v expected %v got %v", i, io.EOF, err)
		}

		if len(records) != len(cases[i]) {
			t.Fatalf("case %v expected result %v got %v", i, cases[i], records)
		}

		for j := 0; j < len(records); j++ {
			if len(records[j]) != len(cases[i][j]) {
				t.Fatalf("case %v expected result %v got %v", i, cases[i], records)
			}

			for k := 0; k < len(records[j]); k++ {
				if records[j][k] != cases[i][j][k] {
					t.Fatalf("case %v expected result %v got %v", i, cases[i], records)
				}
			}
		}
	}
}

// TestReaderToWriter tests cases where feeding CSV to a Reader and then into a Writer should result in unchanged CSV
func TestReaderToWriter(t *testing.T) {
	cases := []string{
		"\"\"\r\n",
		"\"foo\",\"bar\"\r\n\"baz\"\r\n",
	}

	for i := 0; i < len(cases); i++ {
		r := NewReader(strings.NewReader(cases[i]))
		records, err := r.ReadAll()
		if err != io.EOF {
			t.Fatalf("case %v expected %v got %v", i, io.EOF, err)
		}

		b := strings.Builder{}
		w := NewWriter(&b)

		err = w.WriteAll(records)
		if err != nil {
			t.Fatal(err)
		}

		err = w.Flush()
		if err != nil {
			t.Fatal(err)
		}

		if b.String() != cases[i] {
			t.Fatalf("case %v expected result '%s' got '%s'", i, cases[i], b.String())
		}
	}
}
