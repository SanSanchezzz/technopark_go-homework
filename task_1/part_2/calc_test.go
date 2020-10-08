package main

import (
	"testing"
)

func TestAddSuccess(t *testing.T) {
	polishStr := "1 2 +"
	result, err := calculate(polishStr)
	if result != 3 && err != nil {
		t.Fatalf("Add test failed on %s", err)
	}
}

func TestSubSuccess(t *testing.T) {
	polishStr := "10 5 -"
	result, err := calculate(polishStr)
	if result != 5 && err != nil {
		t.Fatalf("Sub test failed on %s", err)
	}
}

func TestMulSuccess(t *testing.T) {
	polishStr := "2 2 *"
	result, err := calculate(polishStr)
	if result != 4 && err != nil {
		t.Fatalf("Mul test failed on %s", err)
	}
}

func TestDivSuccess(t *testing.T) {
	polishStr := "6 3 /"
	result, err := calculate(polishStr)
	if result != 2 && err != nil {
		t.Fatalf("Div test failed on %s", err)
	}
}

func TestDivFail(t *testing.T) {
	polishStr := "69 0 /"
	_, err := calculate(polishStr)
	if err == nil {
		t.Fatalf("Div test not failed on %s", err)
	}
}

func TestPolishNotationSuccess(t *testing.T) {
	str := "(8+2*5)/(1+3*2-4)"
	result, err := getPostfixNotation(str)
	if result != "8 2 5 * + 1 3 2 * + 4 - /" && err != nil {
		t.Fatalf("Polish notation test failed on %s, result is %s", result, err)
	}
}

func TestPolishNotationFail(t *testing.T) {
	str := "(8PLUS2*5)/(1+3*2-4)"
	_, err := getPostfixNotation(str)
	if err == nil {
		t.Fatalf("Polish notation test not failed on %s", str)
	}
}
