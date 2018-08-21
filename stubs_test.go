package csv

import "fmt"

var stubError = fmt.Errorf("stub error")

type readerStub struct{}

func (r readerStub) Read(b []byte) (int, error) {
	return 0, stubError
}

type writerStub struct{}

func (r writerStub) Write(b []byte) (int, error) {
	return 0, stubError
}
