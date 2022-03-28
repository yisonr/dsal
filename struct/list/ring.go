package list

// 环形双向链表
type RingNode struct {
	NodeElem   interface{}
	next, prev *RingNode
}
