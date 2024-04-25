package dsa

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func IsPalindromeLinkedList(head *ListNode) bool {

	var values []int
	current := head
	for current != nil {
		values = append(values, current.Val)
		current = current.Next
	}

	left, right := 0, len(values)-1
	for left < right {
		if values[left] != values[right] {
			return false
		}
		left++
		right--
	}
	return true
}
