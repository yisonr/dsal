package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 反转链表
func ReverseList(pHead *ListNode) *ListNode {
	// 为 nil
	if pHead == nil {
		return pHead
	}

	var newNode = new(ListNode)
	var index int
	for {
		// 结束循环条件
		if pHead == nil {
			break
		}
		if index == 0 {
			newNode.Next = nil
			newNode.Val = pHead.Val
		} else {
			node := new(ListNode)
			node.Next = newNode
			node.Val = pHead.Val
			newNode = node
		}
		pHead = pHead.Next
		index++
	}
	return newNode
}
