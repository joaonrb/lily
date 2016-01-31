package utils

import (
	"bufio"
	"os"
	"io"
)

type fiState int

const (
	FI_SUCCESS fiState = iota
	FI_FAILED
	FI_IGNORED  // In case of lines that not are meant to be executed
)

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

func FileIterator(filename string, method func(line string) fiState) (totalLines, okLines, errLines int, err error) {

	// Open file
	file, err := os.Open(filename)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	fileReader := bufio.NewReader(file)


	for line, lerr := readLine(fileReader); lerr != io.EOF; line, lerr = readLine(fileReader) {
		totalLines ++
		if err != nil {
			errLines ++
		} else {
			switch method(line) {
			case FI_SUCCESS: okLines ++
			case FI_FAILED: errLines ++
			}
		}
	}
	return totalLines, okLines, errLines, nil
}