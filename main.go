package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fileName := "standard.txt"

	// Load the banner map from the file
	banner, err := loadBannerMap(fileName)
	if err != nil {
		fmt.Println("an error occured while loading the banner file")
	}

	// acccepting os args and creating ascii values

	args := os.Args[1]
	args = strings.ReplaceAll(args, "\\n", "\n")

	lines := strings.Split(args, "\n")

	for _, line := range lines {
		printLineBanner(line, banner)
	}
}

func loadBannerMap(fileName string) (map[int][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("an error occured while attempting to open the file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	bannerMap := make(map[int][]string)
	key := 32
	lineCount := 0
	chunk := []string{}

	for scanner.Scan() {
		lines := scanner.Text()

		if lines != "" {
			chunk = append(chunk, lines)
			lineCount++

			if lineCount == 8 {
				bannerMap[key] = chunk
				key++
				chunk = []string{}
				lineCount = 0

			}
		}

	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return bannerMap, nil
}

// Print the banner for a line of text
func printLineBanner(line string, bannerMap map[int][]string) {
	outpput := make([]string, 8)

	if line == "" {
		// fmt.Println()
		return
	}

	for _, char := range line {

		banner, exist := bannerMap[int(char)]
		if !exist {
			fmt.Println("char does not exist within ascii limits")
		}

		for i := 0; i < 8; i++ {
			outpput[i] += banner[i]
		}

	}

	for _,char := range outpput {
		fmt.Println(char)
	}
}
