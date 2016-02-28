//
// Author João Nuno.
//
// joaonrb@gmail.com
//
package utils

import (
	"bufio"
	"os"
	"io"
)

type lineIterator struct {
	file    *os.File
	buffer  *bufio.Reader
	hasNext bool
}

func NewLineIterator(filename string) (*lineIterator, error) {

	// Open file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	fileReader := bufio.NewReader(file)

	return &lineIterator{file, fileReader, true}, nil
}

func readLine(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool = true
		err error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func (self *lineIterator) Next() string {
	line, err := readLine(self.buffer)
	if err == io.EOF { self.hasNext = false }
	return line
}

func (self *lineIterator) HasNext() bool {
	return self.hasNext
}

func (self *lineIterator) Close() {
	self.file.Close()
}
