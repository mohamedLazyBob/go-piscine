package piscine

func BTreeDeleteNode(root, nodeToRemove *TreeNode) *TreeNode {
	if nodeToRemove == nil {
		return root
	}
	if nodeToRemove.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, nodeToRemove)
	} else if nodeToRemove.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, nodeToRemove)
	} else {
		if root.Left == nil {
			temp := root.Right
			root = nil
			return temp
		} else if root.Right == nil {
			temp := root.Left
			root = nil
			return temp
		} else {
			temp := BTreeMin(root.Right)
			root.Data = temp.Data
			root.Right = BTreeDeleteNode(root.Right, temp)
		}
	}
	return root
}
