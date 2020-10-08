package main

import (
	"fmt"
	"strconv"
	"strings"
)

func show(opt Options, result *[]string, str string, repeats int) {
	switch {
	case opt.c:
		*result = append(*result, strconv.Itoa(repeats+1)+" "+str)

	case (opt.d && repeats != 0) || (opt.u && repeats == 0):
		*result = append(*result, str)
	}
}

func skip(opt Options, str string) string {
	fields := strings.Split(str, " ")
	if len(fields) < opt.f {
		return "\n"
	}

	str = strings.Join(fields[opt.f:], " ")
	if len(str) < opt.s {
		return "\n"
	}

	return str[opt.s:]
}

// Uniq - возвращает уникальные строки
func Uniq(opt Options, text []string) []string {
	var result []string
	var currLine string
	var repeats int
	prevLine := skip(opt, text[0])

	if !opt.c && opt.u && opt.d {
		show(opt, &result, text[0], repeats)
	}
	for idx := 1; idx < len(text); idx++ {
		currLine = text[idx]

		if opt.i {
			prevLine = strings.ToLower(prevLine)
			currLine = strings.ToLower(currLine)
		}
		currLine = skip(opt, currLine)

		if prevLine == currLine {
			repeats++
			continue
		}
		if opt.c || !opt.d || !opt.u {
			show(opt, &result, text[idx-1], repeats)
		}
		if !opt.c && opt.u && opt.d {
			show(opt, &result, text[idx], repeats)
		}

		prevLine = currLine
		repeats = 0
	}
	if opt.c || !opt.d || !opt.u {
		show(opt, &result, text[len(text)-1], repeats)
	}

	return result
}

// Usage - выдаёт правильный вызов функции
func Usage() {
	fmt.Println("usage: uniq [-c | -d | -u] [-i] [-f fields] [-s chars] [input [output]]")
}
