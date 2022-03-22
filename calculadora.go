package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calc struct {
}

func (Calc) calculate(input string, operation string) int {

	values := strings.Split(input, operation)
	value1 := parse(values[0])
	value2 := parse(values[0])

	switch operation {
	case "+":
		fmt.Println(value1 + value2)
		return value1 + value2
	case "-":
		fmt.Println(value1 - value2)
		return value1 - value2
	case "*":
		fmt.Println(value1 * value2)
		return value1 * value2
	case "/":
		fmt.Println(value1 / value2)
		return value1 / value2
	default:
		fmt.Println(operation, "not supported")
		return 0
	}
}

func parse(input string) int {
	operator1, _ := strconv.Atoi(input)
	return operator1
}

func ReadInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// func main() {

// 	input := readInput()
// 	operation := readInput()
// 	fmt.Println(input)
// 	fmt.Println(operation)

// 	c := Calc{}
// 	fmt.Println(c.calculate(input, operation))

// }
