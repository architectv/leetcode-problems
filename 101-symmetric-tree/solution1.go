/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return _isSymmetric(root.Left, root.Right)
}

func _isSymmetric(left, right *TreeNode) bool {
	if left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}

		return _isSymmetric(left.Right, right.Left) && _isSymmetric(left.Left, right.Right)
	}

	return left == nil && right == nil
}
