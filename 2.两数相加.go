/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)
	result.Val = (l1.Val + l2.Val) % 10
	carry := (l1.Val + l2.Val) / 10

	if l1.Next == nil && l2.Next == nil {
		if carry == 0 {
			result.Next = nil
		} else {
			result.Next = &ListNode{carry, nil}
		}
	} else if l1.Next == nil {
		result.Next = addTwoNumbers(&ListNode{carry, nil}, l2.Next)
	} else if l2.Next == nil {
		result.Next = addTwoNumbers(l1.Next, &ListNode{carry, nil})
	} else {
		l1.Next.Val += carry
		result.Next = addTwoNumbers(l1.Next, l2.Next)
	}

	return result
}
