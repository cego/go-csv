package csv

import (
	"io"
	"strings"
	"testing"
)

func TestReadAll(t *testing.T) {
	cases := []struct {
		input   string
		records [][]string
	}{
		{``, [][]string{{``}}},
		{`a`, [][]string{{`a`}}},
		{`a,bc`, [][]string{{`a`, `bc`}}},
		{`,abc`, [][]string{{``, `abc`}}},
		{"a\nbc", [][]string{{`a`}, {`bc`}}},
		{"a\r\nbc", [][]string{{`a`}, {`bc`}}},
		{"a\r\nbc\nd", [][]string{{`a`}, {`bc`}, {`d`}}},
		{"\n", [][]string{{``}}},
		{"\na,b\n", [][]string{{``}, {`a`, `b`}}},
		{"\r,\n", [][]string{{"\r", ``}}},
		{"\"a\nb\",c\n", [][]string{{"a\nb", `c`}}},
		{`a,"b,c"`, [][]string{{`a`, `b,c`}}},
		{`a,""""`, [][]string{{`a`, `"`}}},
	}

	for i := 0; i < len(cases); i++ {
		r := NewReader(strings.NewReader(cases[i].input))
		readRecords, err := r.ReadAll()
		if err != io.EOF {
			t.Fatalf("case %v expected %v got %v", i, io.EOF, err)
		}

		if len(cases[i].records) != len(readRecords) {
			t.Fatalf("case %v expected records %#v, got %#v", i, cases[i].records, readRecords)
		}

		for j := 0; j < len(cases[i].records); j++ {
			if len(cases[i].records[j]) != len(readRecords[j]) {
				t.Fatalf("case %v expected records %#v, got %#v", i, cases[i].records, readRecords)
			}

			for k := 0; k < len(cases[i].records[j]); k++ {
				if cases[i].records[j][k] != readRecords[j][k] {
					t.Fatalf("case %v expected records %#v, got %#v", i, cases[i].records, readRecords)
				}
			}
		}
	}
}
