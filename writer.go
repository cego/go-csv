package csv

import (
	"bytes"
	"io"
)

// Writer can be used to write CSV formatted data to an io.Writer
type Writer struct {
	destination io.Writer
	buffer      *bytes.Buffer
}

// NewWriter returns a writer ready to write CSV formatted data to the destination
func NewWriter(destination io.Writer) *Writer {
	return &Writer{
		destination: destination,
		buffer:      &bytes.Buffer{},
	}
}

// Write writes a single record as CSV formatted data to the destination
func (w *Writer) Write(record []string) error {
	for i := 0; i < len(record); i++ {
		if i > 0 {
			w.buffer.Write([]byte(`,`))
		}

		w.buffer.Write([]byte(`"`))

		for j := 0; j < len(record[i]); j++ {
			if record[i][j] == '"' {
				w.buffer.Write([]byte(`""`))
			} else {
				w.buffer.Write([]byte(record[i][j : j+1]))
			}
		}

		w.buffer.Write([]byte(`"`))
	}

	w.buffer.Write([]byte("\r\n"))

	_, err := io.Copy(w.destination, w.buffer)
	return err
}

// WriteAll writes records as CSV formatted data to the destination
func (w *Writer) WriteAll(records [][]string) error {
	for i := 0; i < len(records); i++ {
		err := w.Write(records[i])
		if err != nil {
			return err
		}
	}

	return nil
}
