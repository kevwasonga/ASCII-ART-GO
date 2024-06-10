package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fileName := "standard.txt"

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error occured while opening file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	// initialize  a map to map the provided bannner file

	bannerMap := make(map[int]string)
	key := 32
	lineCount := 0
	chunk := ""
	// scan line by line using scanner.txt

	for scanner.Scan() {
		lines := scanner.Text()
		if lines != "" {
			chunk += lines
			lineCount++
		}
		// allocate a key to every 8 non empty lines
		if lineCount == 8 {
			bannerMap[key] = chunk
			key++
			chunk = ""

		}
		// handling special inputs via os
		if os.Args[1] == "" {
			fmt.Println()
			return

		}
		if os.Args[1] == "\\n" {
			fmt.Println()
			return

		}
		// accept os the rest of os args

		args := os.Args[1:]

		// printing bannerFile Equivalents
		for _, arg := range args {
			arg = strings.ReplaceAll(arg, "\\n", "\n")

			lines := strings.Split(arg, "\n")

			for _, line := range lines {
				if line == "" {
					fmt.Println()
				}
			}
		}

	}
}
