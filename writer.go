package csv

import (
	"bufio"
	"io"
)

// Writer can be used to write CSV formatted data to an io.Writer
type Writer struct {
	w *bufio.Writer
}

// NewWriter returns a writer ready to write CSV formatted data to the destination
func NewWriter(destination io.Writer) *Writer {
	return &Writer{
		w: bufio.NewWriter(destination),
	}
}

// Write writes a single record as CSV formatted data to the destination
func (w *Writer) Write(record []string) error {
	var err error

	for i := 0; i < len(record); i++ {
		if i > 0 {
			_, err = w.w.Write([]byte(`,`))
			if err != nil {
				return err
			}
		}

		_, err = w.w.Write([]byte(`"`))
		if err != nil {
			return err
		}

		for j := 0; j < len(record[i]); j++ {
			if record[i][j] == '"' {
				_, err = w.w.Write([]byte(`""`))
			} else {
				_, err = w.w.Write([]byte(record[i][j : j+1]))
			}

			if err != nil {
				return err
			}
		}

		_, err = w.w.Write([]byte(`"`))
		if err != nil {
			return err
		}
	}

	_, err = w.w.Write([]byte("\r\n"))
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

// Flush flushes the internal write buffer
func (w *Writer) Flush() error {
	return w.w.Flush()
}
