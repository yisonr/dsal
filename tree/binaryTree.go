package main

/*
	BinaryTree

	树的任意节点至多包含两棵子树
	二叉树遍历:
	二叉树的遍历是指从二叉树的根结点出发,按照某种次序依次访问二叉树中的所有
	结点, 使得每个结点被访问一次,且仅被访问一次;
	遍历:
	- 前序遍历
	- 中序遍历
	- 后序遍历
	- 层次遍历
*/

// 二叉树
type BinaryTreeNode struct {
	Elem  interface{}
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func (btn *BinaryTreeNode) LeftInsert(tptr *BinaryTreeNode) error {
	return nil
}

func (btn *BinaryTreeNode) RightInsert(tptr *BinaryTreeNode) error {
	return nil
}

// 对于二叉查找树, 返回指向树 btn 中具有关键字 nodeElem 的节点的指针
func (btn *BinaryTreeNode) Find(nodeElem interface{}) *BinaryTreeNode {
	return nil
}

// 对于二叉查找树, 找到最小元的位置
func (btn *BinaryTreeNode) FindMin() *BinaryTreeNode {
	return nil
}

// 对于二叉查找树, 找到最大元的位置
func (btn *BinaryTreeNode) FindMax() *BinaryTreeNode {
	return nil
}
