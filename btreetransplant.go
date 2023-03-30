package piscine

func BTreeTransplant(root, oldNode, newNode *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if oldNode.Parent == nil {
		root = newNode
	} else if oldNode == oldNode.Parent.Left {
		oldNode.Parent.Left = newNode
	} else {
		oldNode.Parent.Right = newNode
	}
	newNode.Parent = oldNode.Parent
	return root
}

func BTreeReplace(root, oldNode, newNode *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root = BTreeDeleteNode(root, oldNode)
	root = BTreeInsertNode(root, newNode)
	return root
}
