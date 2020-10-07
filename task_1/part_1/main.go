package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
)

func readStrings(stdin *os.File) (strings []string, err error) {
	stringReader := bufio.NewReader(stdin)
	for {
		dataString, err := stringReader.ReadString('\n')
		if err != nil && len(dataString) == 0 {
			break
		} else if err != nil {
			err = errors.New("file reading error")
			return nil, err
		}
		strings = append(strings, dataString[:len(dataString)-1])
	}

	return strings, nil
}

func main() {
	fileIn := os.Stdin
	fileOut := os.Stdout
	var err error

	opt, err := InitFlags()
	if err != nil {
		Usage()
		return
	}

	fileCount := len(flag.Args())
	switch fileCount {
	case 0:
	case 1:
		fileIn, err = os.Open(flag.Args()[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		defer fileIn.Close()

	case 2:
		fileIn, err = os.Open(flag.Args()[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		defer fileIn.Close()
		fileOut, err = os.Create(flag.Args()[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		defer fileOut.Close()

	default:
		Usage()
	}

	text, err := readStrings(fileIn)
	if err != nil || len(text) == 0 {
		fmt.Println(err)
		return
	}

	for _, str := range Uniq(opt, text) {
		_, err = fileOut.WriteString(str + "\n")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
