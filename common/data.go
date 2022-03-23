package main

// TreeNode 二叉树节点类型
type TreeNode struct {
	Val   int       // 节点存储的值
	Left  *TreeNode // 左侧子节点的指针
	Right *TreeNode // 右侧子节点的指针
}

// ListNode 单链表节点类型
type ListNode struct {
	Val  int       // 节点存储的值
	next *ListNode // 指向下一个节点的指针
}
