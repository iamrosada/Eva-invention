package main

import (
	dsa "dsa-golang/dsa"
	"fmt"
)

func main() {
	testCases := []string{"radar", "hello", "A man, a plan, a canal, Panama!", "12321", "not a palindrome"}
	for _, test := range testCases {
		if dsa.IsPalindrome(test) {
			fmt.Printf("%s is a palindrome\n", test)
		} else {
			fmt.Printf("%s is not a palindrome\n", test)
		}
	}
}
