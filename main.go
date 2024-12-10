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
	secondVowelIndex := -1

	for i, char := range word {
		if strings.ContainsRune(vowels, char) {
			if secondVowelIndex == -1 || i-secondVowelIndex > 1 {
				vowelCount++
				if vowelCount == 2 {
					secondVowelIndex = i
					break // Exit the loop after finding the second distinct vowel
				}
			}
		}
	}

	if secondVowelIndex != -1 {
		return word[:secondVowelIndex] + "uzz"
	}

	return word // Return the original word if there are fewer than 2 vowels
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
