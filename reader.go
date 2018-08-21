package csv

import (
	"io"
)

// Reader can be used to read CSV formatted data from an io.Reader
type Reader struct {
	t tokenizer
}

// NewReader returns a Reader ready to start reading CSV formatted data from the given source
func NewReader(r io.Reader) *Reader {
	return &Reader{newTokenizer(r)}
}

// Read reads and returns the next record from the source
func (r *Reader) Read() ([]string, error) {
	var (
		record []string
		field  string
		token  token
		err    error
	)

	for err == nil {
		token, err = r.t.readToken()

		// text
		if token.kind == text {
			field += token.value
		}

		// end of field
		if token.kind == comma || token.kind == lineBreak || token.kind == eof && (r.t.onFirstLine || field != "") {
			record = append(record, field)
			field = ""
		}

		// end of record
		if token.kind == lineBreak || token.kind == eof {
			break
		}

	}

	return record, err
}

// ReadAll reads and returns all remaining records from the source
func (r *Reader) ReadAll() ([][]string, error) {
	var (
		records [][]string
		record  []string
		err     error
	)

	for err == nil {
		record, err = r.Read()
		if len(record) > 0 {
			records = append(records, record)
		}
	}

	return records, err
}
