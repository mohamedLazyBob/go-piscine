package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	if root.Left == nil {
		return root
	} else {
		return BTreeMin(root.Left)
	}
}
