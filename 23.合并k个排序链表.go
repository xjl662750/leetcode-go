/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个排序链表
 *
 * https://leetcode-cn.com/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (47.76%)
 * Likes:    404
 * Dislikes: 0
 * Total Accepted:    58.8K
 * Total Submissions: 122.4K
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * 合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
 *
 * 示例:
 *
 * 输入:
 * [
 * 1->4->5,
 * 1->3->4,
 * 2->6
 * ]
 * 输出: 1->1->2->3->4->4->5->6
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
var temp *ListNode

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	// if len(lists) == 1 {
	// 	return lists[0]
	// }
	// if len(lists) == 2 {
	// 	return mergeTwoLists(lists[0], lists[1])
	// }
	temp := lists[0]
	for k, v := range lists {
		if k != 0 {
			temp = mergeTwoLists(temp, v)
		}
	}
	// var head *ListNode
	// head = mergeTwoLists(mergeTwoLists(lists[0], lists[1]), lists[2])
	return temp
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	tempList := new(ListNode)
	head := tempList
	for {
		if l1 == nil {
			tempList.Next = l2
			break
		}
		if l2 == nil {
			tempList.Next = l1
			break
		}
		if l1.Val < l2.Val {
			tempList.Next = l1
			tempList = tempList.Next
			l1 = l1.Next
		} else {
			tempList.Next = l2
			tempList = tempList.Next
			l2 = l2.Next
		}
	}
	return head.Next
}

// @lc code=end
