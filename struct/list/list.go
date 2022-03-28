package list

/*
	链表由一个个数据节点组成的, 它是一个递归结构,要么它是空的, 要么它存在一个
	指向另外一个数据节点的引用;

	单链表,就是链表是单向的, 可以一直往下找到下一个数据节点, 它只有一个方向;
	双链表,每个节点既可以找到它之前的节点,也可以找到之后的节点, 是双向的;
	循环链表, 一直往下找数据节点，最后回到了自己那个节点, 形成了一个回路, 循环
		单链表和循环双链表的区别就是, 一个只能一个方向走, 一个两个方向都可以走;
*/

type SingleListNode struct {
	NodeElem interface{}
	next     *SingleListNode
}

// 带表头的空表
func InitSingleListNode() *SingleListNode {
	return new(SingleListNode)
}

// 以下方法的接收者皆为链表的表头节点(哨兵)
func (head *SingleListNode) IsEmpty() bool {
	if head.next == nil {
		return true
	}
	return false
}

func (head *SingleListNode) Len() int {
	// 表头不算入
	var length int
	for head.next != nil {
		length++
		head = head.next
	}
	return length
}

func (head *SingleListNode) Find(v interface{}) *SingleListNode {
	return nil
}

// 只删除
func (head *SingleListNode) Delete(v interface{}) *SingleListNode {
	return nil
}

// 单链表默认后插, 任意位置插
func (head *SingleListNode) Insert(v interface{}, node *SingleListNode) *SingleListNode {
	return nil
}

func (head *SingleListNode) Next(v interface{}) *SingleListNode {
	return head.next
}

func (head *SingleListNode) HeadInsert(v interface{}) *SingleListNode {
	return head.next
}

func (head *SingleListNode) TailInsert(v interface{}) *SingleListNode {
	return head.next
}
