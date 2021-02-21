/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return _isValidBST(root, nil, nil)
}

func _isValidBST(root, left, right *TreeNode) bool {
	if root == nil {
		return true
	}

	if left != nil && root.Val <= left.Val {
		return false
	}

	if right != nil && root.Val >= right.Val {
		return false
	}

	return _isValidBST(root.Left, left, root) && _isValidBST(root.Right, root, right)
}
