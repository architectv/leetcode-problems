/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := new(ListNode)

	cur1 := l1
	cur2 := l2
	curSum := sum
	mod := 0

	for cur1 != nil && cur2 != nil {
		curDigit := cur1.Val + cur2.Val + mod
		fmt.Println(curDigit % 10)
		curSum.Val = curDigit % 10
		mod = curDigit / 10

		cur1 = cur1.Next
		cur2 = cur2.Next
		if cur1 != nil || cur2 != nil || mod != 0 {
			curSum.Next = new(ListNode)
		}
		curSum = curSum.Next
	}

	for cur1 != nil {
		curDigit := cur1.Val + +mod
		fmt.Println(curDigit % 10)
		curSum.Val = curDigit % 10
		mod = curDigit / 10

		cur1 = cur1.Next
		if cur1 != nil || mod != 0 {
			curSum.Next = new(ListNode)
		}
		curSum = curSum.Next
	}

	for cur2 != nil {
		curDigit := cur2.Val + +mod
		fmt.Println(curDigit % 10)
		curSum.Val = curDigit % 10
		mod = curDigit / 10

		cur2 = cur2.Next
		if cur2 != nil || mod != 0 {
			curSum.Next = new(ListNode)
		}
		curSum = curSum.Next
	}

	if mod != 0 {
		curSum.Val = mod
	}

	return sum
}
