package bst

// Node represents a node in the binary search tree.
type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

// BST represents the binary search tree.
type BST struct {
	Root *Node
}

// Public Methods

func (bst *BST) Balance() {
	bst.Root = balance(bst.Root)
}

func (bst *BST) Delete(data int) {
	bst.Root = delete(bst.Root, data)
}

func (bst *BST) Height() int {
	return height(bst.Root)
}

func (bst *BST) InOrder() []int {
	var result []int
	inOrder(bst.Root, &result)
	return result
}

func (bst *BST) Insert(data int) {
	if bst.Root == nil {
		bst.Root = &Node{Data: data}
	} else {
		insert(bst.Root, data)
	}
}

func (bst *BST) IsBalanced() bool {
	return isBalanced(bst.Root)
}

func (bst *BST) IsBST() bool {
	return isBST(bst.Root, nil, nil)
}

func (bst *BST) LevelOrder() []int {
	var result []int
	queue := []*Node{bst.Root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			result = append(result, node.Data)
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}
	return result
}

func (bst *BST) Max() int {
	return findMax(bst.Root).Data
}

func (bst *BST) Min() int {
	return findMin(bst.Root).Data
}

func (bst *BST) PostOrder() []int {
	var result []int
	postOrder(bst.Root, &result)
	return result
}

func (bst *BST) PreOrder() []int {
	var result []int
	preOrder(bst.Root, &result)
	return result
}

func (bst *BST) Search(data int) bool {
	return search(bst.Root, data)
}

func (bst *BST) Size() int {
	return size(bst.Root)
}

// Private Methods

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func balance(node *Node) *Node {
	if node == nil {
		return nil
	}
	node.Left = balance(node.Left)
	node.Right = balance(node.Right)
	balanceFactor := height(node.Left) - height(node.Right)
	if balanceFactor > 1 {
		if height(node.Left.Left) >= height(node.Left.Right) {
			node = rotateRight(node)
		} else {
			node.Left = rotateLeft(node.Left)
			node = rotateRight(node)
		}
	} else if balanceFactor < -1 {
		if height(node.Right.Right) >= height(node.Right.Left) {
			node = rotateLeft(node)
		} else {
			node.Right = rotateRight(node.Right)
			node = rotateLeft(node)
		}
	}
	return node
}

func delete(node *Node, data int) *Node {
	if node == nil {
		return nil
	}
	if data < node.Data {
		node.Left = delete(node.Left, data)
	} else if data > node.Data {
		node.Right = delete(node.Right, data)
	} else {
		if node.Left == nil {
			return node.Right
		}
		if node.Right == nil {
			return node.Left
		}
		minNode := findMin(node.Right)
		node.Data = minNode.Data
		node.Right = delete(node.Right, minNode.Data)
	}
	return node
}

func findMax(node *Node) *Node {
	for node.Right != nil {
		node = node.Right
	}
	return node
}

func findMin(node *Node) *Node {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func height(node *Node) int {
	if node == nil {
		return 0
	}
	leftHeight := height(node.Left)
	rightHeight := height(node.Right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func inOrder(node *Node, result *[]int) {
	if node != nil {
		inOrder(node.Left, result)
		*result = append(*result, node.Data)
		inOrder(node.Right, result)
	}
}

func insert(node *Node, data int) {
	if data < node.Data {
		if node.Left == nil {
			node.Left = &Node{Data: data}
		} else {
			insert(node.Left, data)
		}
	} else {
		if node.Right == nil {
			node.Right = &Node{Data: data}
		} else {
			insert(node.Right, data)
		}
	}
}

func isBalanced(node *Node) bool {
	if node == nil {
		return true
	}
	leftHeight := height(node.Left)
	rightHeight := height(node.Right)
	return abs(leftHeight-rightHeight) <= 1 && isBalanced(node.Left) && isBalanced(node.Right)
}

func isBST(node *Node, minNode *Node, maxNode *Node) bool {
	if node == nil {
		return true
	}
	if minNode != nil && node.Data <= minNode.Data {
		return false
	}
	if maxNode != nil && node.Data >= maxNode.Data {
		return false
	}
	return isBST(node.Left, minNode, node) && isBST(node.Right, node, maxNode)
}

func postOrder(node *Node, result *[]int) {
	if node != nil {
		postOrder(node.Left, result)
		postOrder(node.Right, result)
		*result = append(*result, node.Data)
	}
}

func preOrder(node *Node, result *[]int) {
	if node != nil {
		*result = append(*result, node.Data)
		preOrder(node.Left, result)
		preOrder(node.Right, result)
	}
}

func rotateLeft(node *Node) *Node {
	newNode := node.Right
	node.Right = newNode.Left
	newNode.Left = node
	return newNode
}

func rotateRight(node *Node) *Node {
	newNode := node.Left
	node.Left = newNode.Right
	newNode.Right = node
	return newNode
}

func search(node *Node, data int) bool {
	if node == nil {
		return false
	}
	if data == node.Data {
		return true
	}
	if data < node.Data {
		return search(node.Left, data)
	}
	return search(node.Right, data)
}

func size(node *Node) int {
	if node == nil {
		return 0
	}
	return size(node.Left) + size(node.Right) + 1
}
