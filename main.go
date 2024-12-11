package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func modifyWord(word string) string {
	vowels := "aeiouAEIOU"
	vowelCount := 0
	firstVowelIndex := -1
	secondVowelIndex := -1

	for i, char := range word {
		if strings.ContainsRune(vowels, char) {
			vowelCount++
			if vowelCount == 1 {
				firstVowelIndex = i
			}
			if vowelCount == 2 {
				secondVowelIndex = i
				break // Exit the loop after finding the second vowel
			}
		}
	}
	
	if vowelCount >= 2 {
		return word[:secondVowelIndex] + "uzz"
	} else if vowelCount == 1 {
		return word[:firstVowelIndex] + "uzz"
	}
	
	return word // Return the original word if there are no vowels
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter a word (or press Enter to exit): ")
		scanner.Scan()
		input := scanner.Text()

		if input == "" {
			break
		}

		modified := modifyWord(input)
		fmt.Println("Modified word:", modified)
	}

	os.Exit(0)
}