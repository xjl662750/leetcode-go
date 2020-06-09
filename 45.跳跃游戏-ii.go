/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
var arr []int
var k int
for i:=0;i<len(nums);i++{
	arr=append(arr,nums[i]+i)
}
for j:=len(arr)-1;j>0;{
	j=maxIndex(arr[:j],j)
	k++
}
return k
}
func maxIndex(nums []int,m int)int{
	max:=0
	for i:=0;i<len(nums);i++{
		if nums[i]>=m{
			max=i
				break
		}
	}
	return max
}
// @lc code=end

