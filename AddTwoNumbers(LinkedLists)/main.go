package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// printList()
	addInList()

}

func printList() {

	// Создаем звенья списка:
	first := &ListNode{Val: 1}

	second := &ListNode{Val: 2}
	// Присваиваем к первому звену второе:
	first.Next = second
	// Присваиваем к третьему звену первое(со вторым):
	third := &ListNode{Val: 3, Next: first}

	// Печатаем результат *используя цикл*:
	for third != nil {
		fmt.Print(third.Val, " ")
		third = third.Next
	}
}

type List struct {
	Val  int
	Next *List
}

func addInList() {

	foot := &List{Val: 666}
	head := &List{Val: 2, Next: foot}

	ListADD(head, 3)
	ListADD(head, 4)
	ListADD(head, 5)

	// for head != nil {
	// 	fmt.Print(head.Val, " ")
	// 	head = head.Next
	// }

	ListDEL(head, 666)

	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
}

func ListADD(head *List, val int) {

	newNode := &List{Val: val}

	for head.Next != nil {
		head = head.Next
	}

	head.Next = newNode
}

func ListDEL(myHEAD *List, myVAL int) *List {
	if myHEAD == nil {
		return nil
	}
	if myHEAD.Val == myVAL {
		return myHEAD.Next
	}

	for myHEAD.Next != nil {
		if myHEAD.Next.Val == myVAL {
			myHEAD.Next = myHEAD.Next.Next
			return myHEAD
		}
		myHEAD = myHEAD.Next
	}
	return myHEAD

}

// // Основной пример:
// func main() {
// 	// Create the first linked list: 2 -> 4 -> 3
// 	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}

// 	// Create the second linked list: 5 -> 6 -> 4
// 	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}

// 	// Add the two linked lists
// 	result := addTwoNumbers(l1, l2)

// 	fmt.Println(result)

// 	// Print the result
// 	for result != nil {
// 		fmt.Print(result.Val)
// 		if result.Next != nil {
// 			fmt.Print(" -> ")
// 		}
// 		result = result.Next
// 	}
// 	fmt.Println()
// }

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	dummyHead := &ListNode{0, nil}
// 	p, q, curr := l1, l2, dummyHead
// 	carry := 0
// 	for p != nil || q != nil {
// 		x := 0
// 		if p != nil {
// 			x = p.Val
// 			p = p.Next
// 		}
// 		y := 0
// 		if q != nil {
// 			y = q.Val
// 			q = q.Next
// 		}
// 		sum := carry + x + y
// 		carry = sum / 10
// 		curr.Next = &ListNode{sum % 10, nil}
// 		curr = curr.Next
// 	}
// 	if carry > 0 {
// 		curr.Next = &ListNode{carry, nil}
// 	}
// 	return dummyHead.Next
// }
