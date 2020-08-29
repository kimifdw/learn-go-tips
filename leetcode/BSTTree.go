package main

/**
Node 定义节点
*/
type BSTNode struct {
	value int
	left  *BSTNode // 左节点
	right *BSTNode // 右节点
}

//  BST是一个节点的值为int类型的二叉搜索树
type BST struct {
	root *BSTNode
}

// Insert:插入一个元素
func (bst *BST) Insert(value int) {
	newBSTNode := &BSTNode{value: value, left: nil, right: nil}

	// 如果二叉树为空，那个这个节点为根节点
	if bst.root == nil {
		bst.root = newBSTNode
	} else {
		insertNode(bst.root, newBSTNode)
	}
}

// Remove:删除一个元素
func (bst *BST) Remove(value int) bool {
	_, existed := remove(bst.root, value)
	return existed
}

func insertNode(root, newBSTNode *BSTNode) {
	// 放在根节点的左边
	if newBSTNode.value < root.value {
		if root.left == nil {
			root.left = newBSTNode
		} else {
			insertNode(root.left, newBSTNode)
		}
	} else if newBSTNode.value > root.value {
		// 根节点的右边
		if root.right == nil {
			root.right = newBSTNode
		} else {
			insertNode(root.right, newBSTNode)
		}
	}
}

// 用来递归移除节点的辅助方法.
// 返回替换root的新节点，以及元素是否存在
func remove(root *BSTNode, value int) (*BSTNode, bool) {
	if root == nil {
		return nil, false
	}

	var existed bool
	// 从左边找
	if value < root.value {
		root.left, existed = remove(root.left, value)
		return root, existed
	}

	// 从右边找
	if value > root.value {
		root.right, existed = remove(root.right, value)
		return root, existed
	}

	existed = true

	// 如果此节点没有孩子，直接返回即可
	if root.left == nil && root.right == nil {
		root = nil
		return root, existed
	}

	// 左节点为空，提升右节点
	if root.left == nil {
		root = root.right
		return root, existed
	}

	// 右节点为空，提升左节点
	if root.right == nil {
		root = root.left
		return root, existed
	}

	// 如果左右节点都存在,那么从右边节点找到一个最小的节点提升，这个节点肯定比左子树所有节点都大.
	smallestInRight, _ := min(root.right)
	// 提升
	root.value = smallestInRight
	// 从右边子树中移除此节点
	root.right, _ = remove(root.right, smallestInRight)

	return root, existed
}

// 最小值
func (bst *BST) Min() (int, bool) {
	return min(bst.root)
}

func min(node *BSTNode) (int, bool) {
	if node == nil {
		return 0, false
	}

	n := node

	for {
		if n.left == nil {
			return n.value, true
		}
		n = n.left
	}
}

// 最大值
func (bst *BST) Max() (int, bool) {
	return max(bst.root)
}

func max(node *BSTNode) (int, bool) {
	if node == nil {
		return 0, false
	}

	n := node
	// 从右边找
	for {
		if n.right == nil {
			return n.value, true
		}
		n = n.right
	}
}

// Search 搜索元素(检查元素是否存在)
func (bst *BST) Search(value int) bool {
	return search(bst.root, value)
}
func search(n *BSTNode, value int) bool {
	if n == nil {
		return false
	}
	if value < n.value {
		return search(n.left, value)
	}
	if value > n.value {
		return search(n.right, value)
	}
	return true
}

// PreOrderTraverse 前序遍历
func (bst *BST) PreOrderTraverse(f func(int)) {
	preOlderTraverse(bst.root, f)
}
func preOlderTraverse(n *BSTNode, f func(int)) {
	if n != nil {
		f(n.value)
		preOlderTraverse(n.left, f)
		preOlderTraverse(n.right, f)
	}
}

// PostOrderTraverse 后序遍历
func (bst *BST) PostOrderTraverse(f func(int)) {
	postOrderTraverse(bst.root, f)
}
func postOrderTraverse(n *BSTNode, f func(int)) {
	if n != nil {
		postOrderTraverse(n.left, f)
		postOrderTraverse(n.right, f)
		f(n.value)
	}
}
