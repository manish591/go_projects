package main

import (
	"flag"
	"fmt"
	"wc/cmd"
)

const errorMessage = "Error occured while reading file"

func main() {
	countBytesFlag := flag.String("c", "", "Count number of characters")
	countLinesFlag := flag.String("l", "", "Count number of lines")
	countWordsFlag := flag.String("w", "", "Count number of words")
	countCharsFlag := flag.String("m", "", "Count nunber of characters")

	isFlag := false

	flag.Parse()

	if *countBytesFlag != "" {
		isFlag = true
		totatBytes, err := cmd.GetFileSizeInBytes(*countBytesFlag)

		if err != nil {
			fmt.Println(errorMessage)
		} else {
			fmt.Println("Total Bytes:", totatBytes)
		}
	}

	if *countLinesFlag != "" {
		isFlag = true
		totatLines, err := cmd.GetNumLinesInFile(*countLinesFlag)

		if err != nil {
			fmt.Println(errorMessage)
		} else {
			fmt.Println("Total Lines:", totatLines)
		}
	}

	if *countWordsFlag != "" {
		isFlag = true
		totalWords, err := cmd.GetTotalWordsInFile(*countWordsFlag)

		if err != nil {
			fmt.Println(errorMessage)
		} else {
			fmt.Println("Total Words:", totalWords)
		}
	}

	if *countCharsFlag != "" {
		isFlag = true
		totalChars, err := cmd.GetTotalCharsInFile(*countCharsFlag)

		if err != nil {
			fmt.Println(errorMessage)
		} else {
			fmt.Println("Total Characters:", totalChars)
		}
	}

	if !isFlag {
		files := flag.Args()

		if len(files) > 0 {
			filepath := files[0]

			if totatBytes, err := cmd.GetFileSizeInBytes(filepath); err != nil {
				fmt.Println(errorMessage)
			} else {
				fmt.Println("Total Bytes:", totatBytes)
			}

			if totatWords, err := cmd.GetTotalWordsInFile(filepath); err != nil {
				fmt.Println(errorMessage)
			} else {
				fmt.Println("Total Words:", totatWords)
			}

			if totalLines, err := cmd.GetNumLinesInFile(filepath); err != nil {
				fmt.Println(errorMessage)
			} else {
				fmt.Println("Total Lines:", totalLines)
			}
		}
	}
}
