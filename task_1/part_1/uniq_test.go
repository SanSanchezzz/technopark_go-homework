package main

import (
	"reflect"
	"testing"
)

func TestDefaultSuccess(t *testing.T) {
	data := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}
	opt := Options{
		c: false,
		d: true,
		u: true,
		f: 0,
		s: 0,
		i: false,
	}
	expected := []string{
		"I love music.",
		"",
		"I love music of Kartik.",
		"Thanks.",
	}
	result := Uniq(opt, data)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Check default behaviour failed")
	}
}

func TestFlagCSuccess(t *testing.T) {
	data := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}
	opt := Options{
		c: true,
		d: true,
		u: true,
		f: 0,
		s: 0,
		i: false,
	}
	expected := []string{
		"3 I love music.",
		"1 ",
		"2 I love music of Kartik.",
		"1 Thanks.",
	}
	result := Uniq(opt, data)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Check test with -c failed")
	}
}

func TestFlagDSuccess(t *testing.T) {
	data := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}
	opt := Options{
		c: false,
		d: true,
		u: false,
		f: 0,
		s: 0,
		i: false,
	}
	expected := []string{
		"I love music.",
		"I love music of Kartik.",
	}
	result := Uniq(opt, data)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Check test with -d failed")
	}
}

func TestFlagISuccess(t *testing.T) {
	data := []string{
		"I love music.",
		"i love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}
	opt := Options{
		c: false,
		d: true,
		u: true,
		f: 0,
		s: 0,
		i: true,
	}
	expected := []string{
		"I love music.",
		"",
		"I love music of Kartik.",
		"Thanks.",
	}
	result := Uniq(opt, data)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Check test with -d failed")
	}
}

func TestFlagFSuccess(t *testing.T) {
	data := []string{
		"We love music.",
		"I love music.",
		"They love music.",
		"",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}
	opt := Options{
		c: false,
		d: true,
		u: true,
		f: 1,
		s: 0,
		i: false,
	}
	expected := []string{
		"We love music.",
		"",
		"I love music of Kartik.",
		"Thanks.",
	}
	result := Uniq(opt, data)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Check test with -f failed")
	}
}

func TestFlagSSuccess(t *testing.T) {
	data := []string{
		"I love music.",
		"A love music.",
		"C love music.",
		"",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}
	opt := Options{
		c: false,
		d: true,
		u: true,
		f: 0,
		s: 1,
		i: false,
	}
	expected := []string{
		"I love music.",
		"",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}
	result := Uniq(opt, data)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Check test with -s failed")
	}
}