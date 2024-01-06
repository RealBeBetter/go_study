package tree

// Node
//  @Description: 节点
type Node struct {
	Value int   // 节点值
	Left  *Node // 左子树
	Right *Node // 右子树
}

func traverseTree(root *Node) {
	if root == nil {
		return
	}
	traverseTree(root.Left)
	traverseTree(root.Right)
	print(root.Value, " ")
}

func InsertNode(root *Node, value int) *Node {
	if root == nil {
		return &Node{Value: value}
	}
	if value < root.Value {
		root.Left = InsertNode(root.Left, value)
	} else {
		root.Right = InsertNode(root.Right, value)
	}
	return root
}

func DeleteNode(root *Node, value int) *Node {
	if root == nil {
		return root
	}
	if value < root.Value {
		root.Left = DeleteNode(root.Left, value)
	} else if value > root.Value {
		root.Right = DeleteNode(root.Right, value)
	} else {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		temp := root.Right
		for temp.Left != nil {
			temp = temp.Left
		}
		root.Value = temp.Value
	}

	return root
}
