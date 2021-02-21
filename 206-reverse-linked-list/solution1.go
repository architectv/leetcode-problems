/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode // nil
	node := head
	for node != nil {
		tmp := node.Next
		node.Next = prev
		prev = node
		node = tmp
	}
	head = prev
	return head
}
