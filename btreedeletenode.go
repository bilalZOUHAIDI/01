package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root
	} else {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	}
	return root
}

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	} else if elem > root.Data {
		return BTreeSearchItem(root.Right, elem)
	}
	return root
}

func BTreeIsBinary(root *TreeNode) bool {
	return isBST(root, nil, nil)
}

func isBST(node *TreeNode, min, max *string) bool {
	if node == nil {
		return true
	}
	if min != nil && node.Data <= *min {
		return false
	}
	if max != nil && node.Data >= *max {
		return false
	}
	return isBST(node.Left, min, &node.Data) && isBST(node.Right, &node.Data, max)
}

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil || node == nil {
		return root
	}
	if node.Left == nil {
		root = btreeTransplant(root, node, node.Right)
	} else if node.Right == nil {
		root = btreeTransplant(root, node, node.Left)
	} else {
		y := btreeMin(node.Right)
		if y.Parent != node {
			root = btreeTransplant(root, y, y.Right)
			y.Right = node.Right
			y.Right.Parent = y
		}
		root = btreeTransplant(root, node, y)
		y.Left = node.Left
		y.Left.Parent = y
	}
	return root
}

func btreeMin(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil {
		return root
	}
	return btreeMin(root.Left)
}

func btreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node == nil {
		return root
	}
	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Left {
		node.Parent.Left = rplc
	} else {
		node.Parent.Right = rplc
	}
	if rplc != nil {
		rplc.Parent = node.Parent
	}
	return root
}
