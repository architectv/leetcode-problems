// TODO: refactor
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	res = append(res, []int{root.Val})

	tmp := _levelOrder(root.Left, root.Right)
	return append(res, tmp...)
}

func _levelOrder(left, right *TreeNode) [][]int {
	var res, leftSlice, rightSlice [][]int
	var tmp []int
	if left != nil {
		tmp = append(tmp, left.Val)
		leftSlice = _levelOrder(left.Left, left.Right)
	}
	if right != nil {
		tmp = append(tmp, right.Val)
		rightSlice = _levelOrder(right.Left, right.Right)
	}
	if tmp != nil {
		res = append(res, tmp)
	}
	i, j := 0, 0
	var downLevel [][]int
	for i < len(leftSlice) && j < len(rightSlice) {
		curLevel := append(leftSlice[i], rightSlice[j]...)
		downLevel = append(downLevel, curLevel)
		i++
		j++
	}
	for i < len(leftSlice) {
		curLevel := append([]int{}, leftSlice[i]...)
		downLevel = append(downLevel, curLevel)
		i++
	}
	for j < len(rightSlice) {
		curLevel := append([]int{}, rightSlice[j]...)
		downLevel = append(downLevel, curLevel)
		j++
	}

	res = append(res, downLevel...)
	return res
}
