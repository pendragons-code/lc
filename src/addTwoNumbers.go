package main

// Definition for singly-linked list. - not required, just to make the compiler happy
type ListNode struct {
	Val  int
	Next *ListNode
}

// Function to add two numbers represented by linked lists
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// Initialize a dummy node to simplify the code
	dummy := &ListNode{}
	current := dummy
	carry := 0

	// Traverse both lists
	for l1 != nil || l2 != nil || carry != 0 {
		// Get the values from both lists, defaulting to 0 if the list is exhausted
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		val2 := 0
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		// Sum the values along with the carry
		sum := val1 + val2 + carry

		// Calculate the new carry and the digit to store in the current node
		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}

		// Move to the next node
		current = current.Next
	}

	// Return the result starting from the node after dummy
	return dummy.Next
}
