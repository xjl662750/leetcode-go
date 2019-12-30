/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 *
 * https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (36.13%)
 * Likes:    585
 * Dislikes: 0
 * Total Accepted:    92.3K
 * Total Submissions: 253.8K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
 *
 * 示例：
 *
 * 给定一个链表: 1->2->3->4->5, 和 n = 2.
 *
 * 当删除了倒数第二个节点后，链表变为 1->2->3->5.
 *
 *
 * 说明：
 *
 * 给定的 n 保证是有效的。
 *
 * 进阶：
 *
 * 你能尝试使用一趟扫描实现吗？
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var m = make(map[int]*ListNode)
	m[0] = head
	i := 1
	temp := head
	if head.Next == nil {
		return nil
	}
	for temp.Next != nil {
		m[i] = temp.Next
		temp = temp.Next
		i++
	}
	if i == n {
		head = m[1]
	} else {
		m[i-1-n].Next = m[i-n+1]
	}
	return head
}

// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	preHead := new(ListNode)
// 	preHead.Next = head
// 	p := preHead
// 	q := preHead

// 	n ++
// 	for p != nil {
// 		p = p.Next // p先走n+1步，之后和q一起走
// 		if n == 0 {
// 			q = q.Next
// 		} else {
// 			n--
// 		}
// 	}
// 	q.Next = q.Next.Next //要刪除的节点即q.Next
// 	return preHead.Next
// }

// @lc code=end
