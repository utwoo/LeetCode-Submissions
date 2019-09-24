/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodes := make([]*ListNode, 0)
	current := head
	for current != nil {
		nodes = append(nodes, current)
		current = current.Next
	}

	l := len(nodes)

	if n == l {
		head = head.Next
	} else {
		nodes[l-1-n].Next = nodes[l-n].Next
	}

	return head
}
