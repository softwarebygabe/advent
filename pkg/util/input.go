package util

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

type ReaderProcessor[T []string | []int] func(io.Reader) T

func ReadInput[T []string | []int](fname string, fn ReaderProcessor[T]) (T, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return fn(f), nil
}

type LineEvaluator = func(line string)

func EvalEachLine(filepath string, fn LineEvaluator) {
	// open file
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// scan through file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		// parse an input line
		line := scanner.Text()
		fn(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func MustParseInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}
