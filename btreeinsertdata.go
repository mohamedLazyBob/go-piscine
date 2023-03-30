package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	return BTreeInsertNode(root, &TreeNode{Data: data})
}

func BTreeInsertNode(root, newNode *TreeNode) *TreeNode {
	if root == nil {
		return newNode
	}
	if newNode.Data < root.Data {
		root.Left = BTreeInsertNode(root.Left, newNode)
		root.Left.Parent = root
	} else if newNode.Data > root.Data {
		root.Right = BTreeInsertNode(root.Right, newNode)
		root.Right.Parent = root
	}
	return root
}
