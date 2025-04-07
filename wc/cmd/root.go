package cmd

import (
	"bufio"
	"os"
)

func GetFileSizeInBytes(filepath string) (int, error) {
	file, err := os.Open(filepath)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	fileReader := bufio.NewReader(file)
	totalBytesRead := 0
	buffer := make([]byte, 4096)

	for {
		n, err := fileReader.Read(buffer)

		if n > 0 {
			totalBytesRead += n
		}

		if err != nil {
			break
		}
	}

	return totalBytesRead, nil
}

func GetNumLinesInFile(filepath string) (int, error) {
	file, err := os.Open(filepath)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	totalLines := 0

	for fileScanner.Scan() {
		totalLines++
	}

	if err := fileScanner.Err(); err != nil {
		return 0, err
	}

	return totalLines, nil
}

func GetTotalWordsInFile(filepath string) (int, error) {
	file, err := os.Open(filepath)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanWords)

	totalWords := 0

	for fileScanner.Scan() {
		totalWords++
	}

	if err := fileScanner.Err(); err != nil {
		return 0, err
	}

	return totalWords, nil
}

func GetTotalCharsInFile(filepath string) (int, error) {
	file, err := os.Open(filepath)

	if err != nil {
		return 0, err
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanRunes)

	totalCharacters := 0

	for fileScanner.Scan() {
		totalCharacters++
	}

	if err := fileScanner.Err(); err != nil {
		return 0, err
	}

	return totalCharacters, nil
}
