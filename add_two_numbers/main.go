package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	//var head, tail *ListNode
	var head, prev *ListNode
	c1, c2 := l1, l2
	carry := 0
	for c1 != nil || c2 != nil {
		v1, v2 := 0, 0
		if c1 != nil {
			v1 = c1.Val
			c1 = c1.Next
		}
		if c2 != nil {
			v2 = c2.Val
			c2 = c2.Next
		}

		s := (v1 + v2 + carry) % 10
		carry = (v1 + v2 + carry) / 10

		if head == nil {
			head = &ListNode{Val: s}
			prev = head
		} else {
			prev.Next = &ListNode{Val: s}
			prev = prev.Next
		}

	}

	if carry > 0 {
		prev.Next = &ListNode{Val: carry}
	}

	return head
}
