package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	if root.Right == nil {
		return root
	} else {
		return BTreeMax(root.Right)
	}
}
