package mainn

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Основной пример:
func main() {
	// Create the first linked list: 2 -> 4 -> 3
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}

	// Create the second linked list: 5 -> 6 -> 4
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	// Add the two linked lists
	result := addTwoNumbers(l1, l2)

	fmt.Println(result)

	// Print the result
	for result != nil {
		fmt.Print(result.Val)
		if result.Next != nil {
			fmt.Print(" -> ")
		}
		result = result.Next
	}
	fmt.Println()
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{0, nil}
	p, q, curr := l1, l2, dummyHead
	carry := 0
	for p != nil || q != nil {
		x := 0
		if p != nil {
			x = p.Val
			p = p.Next
		}
		y := 0
		if q != nil {
			y = q.Val
			q = q.Next
		}
		sum := carry + x + y
		carry = sum / 10
		curr.Next = &ListNode{sum % 10, nil}
		curr = curr.Next
	}
	if carry > 0 {
		curr.Next = &ListNode{carry, nil}
	}
	return dummyHead.Next
}
