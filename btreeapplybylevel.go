package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	h := BTreeLevelCount(root)
	for i := 0; i < h; i++ {
		ApplyLevel(root, i, f)
	}
}

func ApplyLevel(root *TreeNode, level int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if level == 0 {
		f(root.Data)
	} else if level > 0 {
		ApplyLevel(root.Left, level-1, f)
		ApplyLevel(root.Right, level-1, f)
	}
}
