package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	byteFlag := flag.Bool("c", false, "Count number of characters")
	lineFlag := flag.Bool("l", false, "Count number of lines")
	wordFlag := flag.Bool("w", false, "Count number of words")
	charFlag := flag.Bool("m", false, "Count nunber of characters")

	flag.Parse()

	var fileData string
	var filepath string

	fi, _ := os.Stdin.Stat()

	if fi.Mode()&os.ModeCharDevice == 0 {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		fileData = string(data)
	} else {
		args := flag.Args()

		if len(args) > 0 {
			filepath = args[0]
		} else {
			panic("Enter atleast 1 file")
		}

		fileReadData, err := os.ReadFile(filepath)

		if err != nil {
			panic(err)
		}

		fileData = string(fileReadData)
	}

	numBytes := len(fileData)
	numLines := len(strings.Split(string(fileData), "\n"))
	numWords := len(strings.Fields(string(fileData)))
	numChars := utf8.RuneCountInString(string(fileData))

	if *byteFlag {
		fmt.Println(numBytes)
	} else if *lineFlag {
		fmt.Println(numLines)
	} else if *wordFlag {
		fmt.Println(numWords)
	} else if *charFlag {
		fmt.Println(numChars)
	} else {
		fmt.Println(numBytes, numWords, numChars)
	}
}
