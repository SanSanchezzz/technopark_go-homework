package main

import (
	"errors"
	"flag"
)

// Options - структура для передачи флагов
type Options struct {
	c bool
	d bool
	u bool
	f int
	s int
	i bool
}

// InitFlags - инициализирует флаги для uniq
func InitFlags() (Options, error) {
	flagCPtr := flag.Bool("c", false, "for number of occurrences of lines in the input")
	flagDPtr := flag.Bool("d", false, "print only those lines that were repeated in the input data")
	flagUPtr := flag.Bool("u", false, "print only those lines that have not been repeated in the input data")

	flagIPtr := flag.Bool("i", false, "case-insensitive")
	flagFPtr := flag.Int("f", 0, "ignore the first num_fields fields in the line")
	flagSPtr := flag.Int("s", 0, "ignore the first num_chars characters in the string")

	flag.Parse()

	opt := Options{
		c: *flagCPtr,
		d: *flagDPtr,
		u: *flagUPtr,
		f: *flagFPtr,
		s: *flagSPtr,
		i: *flagIPtr,
	}
	if !opt.c && opt.d && opt.u || opt.c && !opt.d && opt.u || opt.c && opt.d && !opt.u {
		return opt, errors.New("invalid arguments passed")
	}

	if !opt.c && !opt.u && !opt.d {
		opt.u = true
		opt.d = true
	}

	return opt, nil
}
