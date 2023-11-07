package main

import (
	"fmt"
)

type Stack []string

var OpsList = []string{"+", "-", "*", "/", "^"}

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func ComposeRPNFromText(text string) (finalOutput string) {
	i := 1
	stack := Stack{}
	for {
		//check first index of remaining text
		if text[i-1:i] == "(" {
			//push the "(" into stack
			stack.Push(text[i-1 : i])
		} else if text[i-1:i] == ")" {
			ops, err := stack.Pop()
			if !err {
				fmt.Println("error! stack is empty")
				return
			}
			//pop and print the ops
			finalOutput += ops
			//pop but dont print the ")"
			stack.Pop()
		} else if contains(OpsList, text[i-1:i]) {
			stack.Push(text[i-1 : i])
		} else {
			finalOutput += text[i-1 : i]
		}
		if i >= len(text) {
			break
		}
		i++
	}
	fmt.Println("result:", finalOutput)
	return
}
