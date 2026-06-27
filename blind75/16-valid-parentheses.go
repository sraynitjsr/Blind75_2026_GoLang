package blind75

import "fmt"

type Stack struct {
	data []byte
}

func (s *Stack) Push(c byte) {
	s.data = append(s.data, c)
}

func (s *Stack) Pop() byte {
	temp := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return temp
}

func (s *Stack) Peek() byte {
	return s.data[len(s.data)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func isValid(s string) bool {
	myStack := Stack{
		data: make([]byte, 0),
	}

	for _, val := range s {
		if val == '(' || val == '{' || val == '[' {
			myStack.Push(byte(val))
		} else {
			if myStack.IsEmpty() {
				return false
			}
			if val == ')' && myStack.Peek() != '(' {
				return false
			}

			if val == '}' && myStack.Peek() != '{' {
				return false
			}

			if val == ']' && myStack.Peek() != '[' {
				return false
			}
			myStack.Pop()
		}
	}
	return myStack.IsEmpty()
}

func SixteenValidParantheses() {
	expression := "()[]{}"
	fmt.Printf("%v Expression is Valid Parentheses or Not => %v\n", expression, isValid(expression))
}
