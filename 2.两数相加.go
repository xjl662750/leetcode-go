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
//  import "fmt"
// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	var r *ListNode
// 	var eachSum int
// 	for l1.Next!=nil{
// 		eachSum=l1.Val+l2.Val
// 		r.Val=r.Val+eachSum%10
// 		r=r.Next
// 		r.Val=r.Val+eachSum/10
// 	}
// 	// fmt.Println(r.Next.Next.Next==nil)
//     return r
// }
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    result := &ListNode{}
	current := result
	each_sums := 0

	for l1 != nil || l2 != nil || each_sums != 0{
		if l1 != nil{
			each_sums += l1.Val
			l1 = l1.Next
		}
		if l2 != nil{
			each_sums += l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: each_sums % 10}
		each_sums /= 10
		current = current.Next
	}

	return result.Next
}

