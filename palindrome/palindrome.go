package main

import (
	"fmt"
	"sort"
)

type Pallindrome struct {
	start int
	end   int
}

func main() {

	fmt.Println("Enter a word")
	var word string
	fmt.Scanln(&word)

	fmt.Println("Enter a minimum length palindrome")
	var min int
	fmt.Scanln(&min)

	var pallindromes []Pallindrome

	for i := 0; i+min <= len(word); i++ {

		var currWord string = word[i : i+min-1]

		for end := i + min - 1; end < len(word); end++ {
			currWord += word[end : end+1]

			if isPalindrome(currWord) {
				pallindromes = append(pallindromes, Pallindrome{start: i, end: end})
				end = len(word)
			}
		}
	}

	sort.SliceStable(pallindromes, func(i, j int) bool {
		return pallindromes[i].end < pallindromes[j].end
	})

	var max int = 0

	if len(pallindromes) > 0 {
		var end int = -1

		for _, curr := range pallindromes {
			if curr.start > end {
				end = curr.end
				max++
			}
		}

	}

	fmt.Println(max)

}

func isPalindrome(str string) bool {
	for i := 0; i < len(str); i++ {
		j := len(str) - 1 - i
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
