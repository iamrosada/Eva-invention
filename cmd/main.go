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

	//Palindrome
	head1 := &dsa.ListNode{Val: 1, Next: &dsa.ListNode{Val: 2, Next: &dsa.ListNode{Val: 2, Next: &dsa.ListNode{Val: 1}}}}
	fmt.Println("Example 1:", dsa.IsPalindromeLinkedList(head1)) // Output

	//680
	fmt.Println(dsa.ValidPalindrome("aba"))  // Output: true
	fmt.Println(dsa.ValidPalindrome("abca")) // Output: true
	fmt.Println(dsa.ValidPalindrome("abc"))  // Output: false
}
