package main

import "errors"

type Stack struct {
	buff []interface{}
}

func (s *Stack) Len() int {
	return len(s.buff)
}

func (s *Stack) Pop() (interface{}, error) {
	len := len(s.buff)
	if len == 0 {
		return ' ', errors.New("empty stack")
	}

	popEl := s.buff[len-1]
	s.buff = s.buff[:len-1]
	return popEl, nil
}

func (s *Stack) Push(val interface{}) {
	s.buff = append(s.buff, val)
}

func (s *Stack) Peek() (interface{}, error) {
	len := len(s.buff)
	if len == 0 {
		return ' ', errors.New("empty stack")
	}

	return s.buff[len-1], nil
}
