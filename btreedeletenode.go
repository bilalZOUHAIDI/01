package piscine

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
