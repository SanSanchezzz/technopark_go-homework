package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getExample() (string, error) {
	if len(os.Args) != 2 {
		return "", errors.New("invalid argument")
	}

	return os.Args[1], nil
}

func getPriorityOp(op rune) int {
	switch op {
	case '(':
		return 0
	case '+':
		fallthrough
	case '-':
		return 1
	case '*':
		fallthrough
	case '/':
		return 2
	default:
		return -1
	}
}

func getPostfixNotation(example string) (string, error) {
	stackOps := &Stack{buff: make([]interface{}, 0)} // constructor
	numFlag := false
	result := ""

	for _, char := range example {
		if char > '0' && char < '9' {

			if numFlag == false {
				result += " "
			}
			result += string(char)
			numFlag = true

			continue
		}

		numFlag = false

		switch char {
		case '(':
			stackOps.Push(char)
		case ')':
			for stackOps.Len() != 0 {
				popEl, _ := stackOps.Pop()
				if popEl == '(' {
					continue
				}
				result += " " + string(popEl.(rune))
			}
		case '+':
			fallthrough
		case '-':
			fallthrough
		case '*':
			fallthrough
		case '/':
			if stackOps.Len() == 0 {
				stackOps.Push(char)
				continue
			}
			peekEl, err := stackOps.Peek()
			if err != nil {
				return "", errors.New("invalid char")
			}

			if getPriorityOp(peekEl.(rune)) < getPriorityOp(char) {
				stackOps.Push(char)
				continue
			} else {
				popEl, _ := stackOps.Pop()
				result += " " + string(popEl.(rune))
				stackOps.Push(char)
			}
		default:
			return "", errors.New("invalid char")
		}
	}

	for stackOps.Len() != 0 {
		popEl, _ := stackOps.Pop()
		result += " " + string(popEl.(rune))
	}

	return result[1:], nil
}

func expression(lValue, rValue float64, opr string) (float64, error) {
	var res float64
	switch opr {
	case "+":
		res = lValue + rValue
	case "-":
		res = lValue - rValue
	case "*":
		res = lValue * rValue
	case "/":
		if rValue == 0 {
			return 0, errors.New("division by zero")
		}
		res = lValue / rValue
	}

	return res, nil
}

func calculate(exampleString string) (float64, error) {
	example := strings.Split(exampleString, " ")
	stack := &Stack{buff: make([]interface{}, 0)} // constructor

	for _, val := range example {
		num, numFlag := strconv.ParseFloat(val, 64)
		if numFlag == nil {
			stack.Push(num)
			continue
		}

		rNum, rErr := stack.Pop()
		lNum, lErr := stack.Pop()
		if lErr != nil || rErr != nil {
			return 0, errors.New("error in example")
		}

		res, err := expression(lNum.(float64), rNum.(float64), val)
		if err != nil {
			return 0, err
		}
		stack.Push(res)
	}

	result, _ := stack.Pop()
	return result.(float64), nil
}

func main() {
	example, err := getExample()
	if err != nil {
		fmt.Println("error")
		return
	}

	postfixNotation, err := getPostfixNotation(example)
	if err != nil {
		fmt.Println(err)
	}

	result, err := calculate(postfixNotation)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
