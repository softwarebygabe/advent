package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseIntCodeInput(filepath string) []int {
	result := []int{}
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numberStrings := strings.Split(line, ",")
		for _, num := range numberStrings {
			integer, err := strconv.ParseInt(num, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			result = append(result, int(integer))
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func add(intcodeIn []int, cursorIn int, paramModes []int) (intcode []int, cursor int, err error) {
	// fmt.Println(
	// 	"addOp ->",
	// 	intcodeIn[cursorIn],
	// 	intcodeIn[cursorIn+1],
	// 	intcodeIn[cursorIn+2],
	// 	intcodeIn[cursorIn+3],
	// 	paramModes,
	// )
	oplength := 4
	cursorEnd := cursorIn + oplength
	transformedParams := []int{}
	paramCur := 0
	for i := cursorIn + 1; i < cursorEnd; i++ {
		argVal := intcodeIn[i]
		mode := 0
		if paramCur < len(paramModes) {
			mode = paramModes[paramCur]
		}
		if i < cursorEnd-1 {
			switch mode {
			case 0:
				// position
				if argVal > len(intcodeIn)-1 {
					return nil, 0, errors.New("position outside of intcode len")
				}
				transformedParams = append(transformedParams, intcodeIn[argVal])
			case 1:
				// value
				transformedParams = append(transformedParams, argVal)
			}
		} else {
			transformedParams = append(transformedParams, argVal)
		}
		paramCur++
	}

	sum := transformedParams[0] + transformedParams[1]

	// store sum
	intcodeIn[transformedParams[2]] = sum

	intcode = intcodeIn
	cursor = cursorEnd
	return
}

func multiply(intcodeIn []int, cursorIn int, paramModes []int) (intcode []int, cursor int, err error) {
	oplength := 4
	cursorEnd := cursorIn + oplength
	transformedParams := []int{}
	paramCur := 0
	for i := cursorIn + 1; i < cursorEnd; i++ {
		argVal := intcodeIn[i]
		mode := 0
		if paramCur < len(paramModes) {
			mode = paramModes[paramCur]
		}
		if i < cursorEnd-1 {
			switch mode {
			case 0:
				// position
				if argVal >= len(intcodeIn) {
					return nil, 0, errors.New("position outside of intcode len")
				}
				transformedParams = append(transformedParams, intcodeIn[argVal])
			case 1:
				// value
				transformedParams = append(transformedParams, argVal)
			}
		} else {
			transformedParams = append(transformedParams, argVal)
		}
		paramCur++
	}

	product := transformedParams[0] * transformedParams[1]

	// store product
	intcodeIn[transformedParams[2]] = product

	intcode = intcodeIn
	cursor = cursorEnd
	return
}

func input(intcodeIn []int, cursorIn, inputVal int) (intcode []int, cursor int, err error) {
	oplength := 2

	// save at proper place
	intcode = intcodeIn
	intcode[intcode[cursorIn+1]] = inputVal
	cursor = cursorIn + oplength
	return
}

func output(intcodeIn []int, cursorIn int, paramModes []int) (intcode []int, cursor int, err error) {
	oplength := 2

	mode := 0
	if len(paramModes) == 1 {
		mode = paramModes[0]
	}

	outputVal := 0
	switch mode {
	case 0:
		outputVal = intcodeIn[intcodeIn[cursorIn+1]]
	case 1:
		outputVal = intcodeIn[cursorIn+1]
	}

	// just print the right value
	fmt.Printf("output -> %d\n", outputVal)

	intcode = intcodeIn
	cursor = cursorIn + oplength
	return
}

func parseOpCode(val int) (opcode int, paramModes []int, err error) {
	fullS := strconv.Itoa(val)
	fullSList := strings.Split(fullS, "")
	opcodeS := ""
	paramModes = []int{}
	for i := len(fullSList) - 1; i >= 0; i-- {
		s := fullSList[i]
		if i > len(fullSList)-3 {
			opcodeS = s + opcodeS
		} else {
			sI, _ := strconv.Atoi(s)
			paramModes = append(paramModes, sI)
		}
	}
	opcode, err = strconv.Atoi(opcodeS)
	// fmt.Println(val, opcode, paramModes)
	return
}

func computer(intcode []int, startingInput int) ([]int, error) {
	cursor := 0
	stop := false
	for !stop {

		opcode, paramModes, err := parseOpCode(intcode[cursor])
		if err != nil {
			return intcode, err
		}

		switch opcode {
		case 1:
			intcode, cursor, err = add(intcode, cursor, paramModes)
			if err != nil {
				return intcode, err
			}
		case 2:
			intcode, cursor, err = multiply(intcode, cursor, paramModes)
			if err != nil {
				return intcode, err
			}
		case 3:
			intcode, cursor, err = input(intcode, cursor, startingInput)
			if err != nil {
				return intcode, err
			}
		case 4:
			intcode, cursor, err = output(intcode, cursor, paramModes)
			if err != nil {
				return intcode, err
			}
		case 99:
			stop = true
		}
		if cursor >= len(intcode) {
			stop = true
		}
	}
	return intcode, nil
}

func main() {
	fmt.Println("Hello World")

	intcodeInput := parseIntCodeInput("../input.txt")

	startingInputVal := 1

	_, err := computer(intcodeInput, startingInputVal)

	if err != nil {
		panic(err)
	}

}
