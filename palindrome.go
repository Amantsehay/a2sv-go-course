package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s) 
	for i, j := 0, len(s)-1; i < j; {
		if !unicode.IsLetter(rune(s[i])) { 
			i++
			continue
		}
		if !unicode.IsLetter(rune(s[j])) { 
			j--
			continue
		}
		if s[i] != s[j] { 
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	s := "A man, a plan, a canal, Panama!"
	
	fmt.Printf("Is the string \"%s\" a palindrome? %t\n", s, isPalindrome(s))
}
