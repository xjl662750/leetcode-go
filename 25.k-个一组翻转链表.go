/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 *
 * https://leetcode-cn.com/problems/reverse-nodes-in-k-group/description/
 *
 * algorithms
 * Hard (54.87%)
 * Likes:    320
 * Dislikes: 0
 * Total Accepted:    30.1K
 * Total Submissions: 54.6K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
 * 
 * k 是一个正整数，它的值小于或等于链表的长度。
 * 
 * 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
 * 
 * 示例 :
 * 
 * 给定这个链表：1->2->3->4->5
 * 
 * 当 k = 2 时，应当返回: 2->1->4->3->5
 * 
 * 当 k = 3 时，应当返回: 3->2->1->4->5
 * 
 * 说明 :
 * 
 * 
 * 你的算法只能使用常数的额外空间。
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 * 
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
func reverseKGroup(head *ListNode, k int) *ListNode {
	var m =make(map[int]*ListNode,k)
	for i:=0;i<k;i++{
		if head!=nil{
			m[i]=head
			head=head.Next
		}else{
			return m[0]
		}
	}
	for i:=k-1;i>0;i--{	
		m[i].Next=m[i-1]
	}
	m[0].Next=reverseKGroup(head,k)
	return m[k-1]
}
// func reverseKGroup(head *ListNode, k int) *ListNode {
// 	if head == nil {
// 		return nil
//  }

//  index := 0
//  p, cur := head, head
//  var pre, temp *ListNode
//  pre = nil
//  temp = nil

//  for p != nil && index < k {

// 		 p = p.Next
// 		 index++
//  }
 
//  if index < k {
// 		 return head
//  } else {
// 		 loop := 0
// 		 for loop < k {
// 				 temp = cur.Next
// 				 cur.Next = pre
// 				 pre = cur
// 				 cur = temp
// 				 loop++
// 		 }       
// 		 head.Next = reverseKGroup(temp, k)
// 		 return pre
//  }  

// }
// @lc code=end

