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

func transformInstructionParams(intcode, paramModes []int, cursor, oplength int) ([]int, error) {
	// fmt.Println(
	// 	intcode[cursor],
	// 	intcode[cursor+1],
	// 	intcode[cursor+2],
	// 	intcode[cursor+3],
	// 	paramModes,
	// )
	cursorEnd := cursor + oplength
	transformedParams := []int{}
	paramCur := 0
	for i := cursor + 1; i < cursorEnd; i++ {
		argVal := intcode[i]
		mode := 0
		if paramCur < len(paramModes) {
			mode = paramModes[paramCur]
		}
		if i < cursorEnd-1 {
			switch mode {
			case 0:
				// position
				if argVal > len(intcode)-1 {
					return nil, errors.New("position outside of intcode len")
				}
				transformedParams = append(transformedParams, intcode[argVal])
			case 1:
				// value
				transformedParams = append(transformedParams, argVal)
			}
		} else {
			transformedParams = append(transformedParams, argVal)
		}
		paramCur++
	}
	return transformedParams, nil
}

func add(intcodeIn []int, cursorIn int, paramModes []int) (intcode []int, cursor int, err error) {
	oplength := 4
	transformedParams, err := transformInstructionParams(
		intcodeIn,
		paramModes,
		cursorIn,
		oplength,
	)
	if err != nil {
		return nil, 0, err
	}

	sum := transformedParams[0] + transformedParams[1]

	// store sum
	intcodeIn[transformedParams[2]] = sum

	intcode = intcodeIn
	cursor = cursorIn + oplength
	return
}

func multiply(intcodeIn []int, cursorIn int, paramModes []int) (intcode []int, cursor int, err error) {
	oplength := 4
	transformedParams, err := transformInstructionParams(
		intcodeIn,
		paramModes,
		cursorIn,
		oplength,
	)
	if err != nil {
		return nil, 0, err
	}

	product := transformedParams[0] * transformedParams[1]

	// store product
	intcodeIn[transformedParams[2]] = product

	intcode = intcodeIn
	cursor = cursorIn + oplength
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

func jumpIfTrue(intcodeIn []int, cursorIn int, paramModes []int) (intcode []int, cursor int, err error) {
	oplength := 3
	transformedParams, err := transformInstructionParams(
		intcodeIn,
		paramModes,
		cursorIn,
		oplength,
	)
	if err != nil {
		return nil, 0, err
	}

	intcode = intcodeIn
	// should jump ?
	if transformedParams[0] != 0 {
		// jump!
		if len(paramModes) == 2 {
			if paramModes[1] == 1 {
				cursor = transformedParams[1]
			} else {
				cursor = intcode[transformedParams[1]]
			}
		} else {
			cursor = intcode[transformedParams[1]]
		}
		return
	}
	cursor = cursorIn + oplength
	return
}

func jumpIfFalse(intcodeIn []int, cursorIn int, paramModes []int) (intcode []int, cursor int, err error) {
	oplength := 3
	transformedParams, err := transformInstructionParams(
		intcodeIn,
		paramModes,
		cursorIn,
		oplength,
	)
	if err != nil {
		return nil, 0, err
	}

	intcode = intcodeIn
	// should jump ?
	if transformedParams[0] == 0 {
		// jump!
		if len(paramModes) == 2 {
			if paramModes[1] == 1 {
				cursor = transformedParams[1]
			} else {
				cursor = intcode[transformedParams[1]]
			}
		} else {
			cursor = intcode[transformedParams[1]]
		}
		return
	}
	cursor = cursorIn + oplength
	return
}

func lessThan(intcodeIn []int, cursorIn int, paramModes []int) (intcode []int, cursor int, err error) {
	oplength := 4
	transformedParams, err := transformInstructionParams(
		intcodeIn,
		paramModes,
		cursorIn,
		oplength,
	)
	if err != nil {
		return nil, 0, err
	}

	if transformedParams[0] < transformedParams[1] {
		intcodeIn[transformedParams[2]] = 1
	} else {
		intcodeIn[transformedParams[2]] = 0
	}

	intcode = intcodeIn
	cursor = cursorIn + oplength
	return
}

func equals(intcodeIn []int, cursorIn int, paramModes []int) (intcode []int, cursor int, err error) {
	oplength := 4
	transformedParams, err := transformInstructionParams(
		intcodeIn,
		paramModes,
		cursorIn,
		oplength,
	)
	if err != nil {
		return nil, 0, err
	}

	if transformedParams[0] == transformedParams[1] {
		intcodeIn[transformedParams[2]] = 1
	} else {
		intcodeIn[transformedParams[2]] = 0
	}

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
		case 5:
			intcode, cursor, err = jumpIfTrue(intcode, cursor, paramModes)
			if err != nil {
				return intcode, err
			}
		case 6:
			intcode, cursor, err = jumpIfFalse(intcode, cursor, paramModes)
			if err != nil {
				return intcode, err
			}
		case 7:
			intcode, cursor, err = lessThan(intcode, cursor, paramModes)
			if err != nil {
				return intcode, err
			}
		case 8:
			intcode, cursor, err = equals(intcode, cursor, paramModes)
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

	startingInputVal := 5

	_, err := computer(intcodeInput, startingInputVal)

	if err != nil {
		panic(err)
	}

}
