package dsa

import (
	"strings"
	"unicode"
)

func IsPalindrome(value string) bool {

	word := preprocess(value)

	left := 0

	right := len(word) - 1

	for left < right {

		if word[left] != word[right] {
			return false
		}

		left++
		right--

	}
	return true

}

func preprocess(value string) string {
	var result string

	for _, char := range value {

		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			result += strings.ToLower(string(char))
		}
	}
	return result
}

/**
 Other way is using this step

 -I need only to know if the return will be uppercase or lowercase

func preprocess(value string) string {
    var result string
    for _, char := range value {
        // Check if the character is alphanumeric
        if ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') || ('0' <= char && char <= '9') {
            result += string(char)
        }
    }
    return strings.ToLower(result)
}

**/
/**
125. Valid Palindrome
A phrase is a palindrome if, after converting all
uppercase letters into lowercase letters and
removing all non-alphanumeric characters,
it reads the same forward and backward.
Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.



Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
Example 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.


Constraints:

1 <= s.length <= 2 * 105
s consists only of printable ASCII characters.



**/
