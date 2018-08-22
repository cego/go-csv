package csv

import "fmt"

var stubError = fmt.Errorf("stub error")

type readerStub struct{}

func (r readerStub) Read(b []byte) (int, error) {
	return 0, stubError
}

type writerStub struct {
	errorCountdown int
}

func (r *writerStub) Write(b []byte) (int, error) {
	r.errorCountdown -= len(b)
	if r.errorCountdown <= 0 {
		return 0, stubError
	}

	return len(b), nil
}

func (r *writerStub) Flush() error {
	if r.errorCountdown <= 0 {
		return stubError
	}
	return nil
}
