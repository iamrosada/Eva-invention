package dsa

func ValidPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			// Try deleting either the character at the left pointer or the right pointer
			return isPalindrome(s, left+1, right) || isPalindrome(s, left, right-1)
		}
		left++
		right--
	}
	return true
}

// Function to check if a string is a palindrome within the specified range
func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

/**

We have two options: we can either remove the character at the left position or the right position.

Let's say we have the string "abca" as an example:

Initially, left = 0 and right = len(s)-1 = 3.
We compare s[0] and s[3]. Since they are different (s[0] != s[3]),
we have to try deleting one of the characters to see if the resulting string is a palindrome.
We try two possibilities:

Deleting the character at the left position:
We call isPalindrome(s, left+1, right).

This effectively means we are skipping the character at position left.

So, we compare s[1] and s[3].
If they match, we continue checking if the substring "bca" is a palindrome.

Deleting the character at the right position: We call isPalindrome(s, left, right-1).
This effectively means we are skipping the character at position right.

So, we compare s[0] and s[2].
If they match, we continue checking if the substring "abc" is a palindrome.
If any of these substrings are palindromes, we return true,
indicating that the original string s can be a palindrome after deleting at most one character.

**/

/**
680. Valid Palindrome II

Given a string s,
return true if the s can be palindrome after
deleting at most one character from it.



Example 1:

Input: s = "aba"
Output: true
Example 2:

Input: s = "abca"
Output: true
Explanation: You could delete the character 'c'.
Example 3:

Input: s = "abc"
Output: false


Constraints:

1 <= s.length <= 105
s consists of lowercase English letters.

**/
