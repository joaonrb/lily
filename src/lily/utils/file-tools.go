package utils

import (
	"bufio"
	"os"
	"io"
)

type lineIterator struct {
	file   *os.File
	buffer *bufio.Reader
}

func NewLineIterator(filename string) *lineIterator {

	// Open file
	file, err := os.Open(filename)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	fileReader := bufio.NewReader(file)

	return &lineIterator{file, fileReader}
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

/**
 * Return the next line and if is last.
 * return line string: The next line in the file.
 * return isLast bool: If this line is the last in the file.
 */
func (self *lineIterator) Next() (line string, isLast bool) {
	line, err := readLine(self.buffer)
	if err == io.EOF {
		return line, true
	}
	return line, false
}

func (self *lineIterator) Close() {
	self.file.Close()
}
