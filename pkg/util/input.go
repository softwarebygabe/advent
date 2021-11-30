package util

import (
	"bufio"
	"log"
	"os"
)

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
