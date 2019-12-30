/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 *
 * https://leetcode-cn.com/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (36.08%)
 * Likes:    450
 * Dislikes: 0
 * Total Accepted:    60.5K
 * Total Submissions: 167K
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
 * 
 * ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
 * 
 * 搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
 * 
 * 你可以假设数组中不存在重复的元素。
 * 
 * 你的算法时间复杂度必须是 O(log n) 级别。
 * 
 * 示例 1:
 * 
 * 输入: nums = [4,5,6,7,0,1,2], target = 0
 * 输出: 4
 * 
 * 
 * 示例 2:
 * 
 * 输入: nums = [4,5,6,7,0,1,2], target = 3
 * 输出: -1
 * 
 */

// @lc code=start
func search(nums []int, target int) int {
	n:=len(nums)
	if n == 1&&nums[0] == target{
		return  0 
	}
	if n == 0||n == 1{
		return -1;
	}

  rotateIndex:=findRotateIndex(nums,0,len(nums)-1)
	if nums[rotateIndex]==target{
		return rotateIndex
	}
	if rotateIndex==0{
		return binarySearch(nums,target,0,len(nums)-1)
	}
	if target<nums[0]{
		return binarySearch(nums,target,rotateIndex,len(nums)-1)
	}
	return binarySearch(nums,target,0,rotateIndex)
	
}

func findRotateIndex(nums []int,left,right int) int{
	if nums[left]<nums[right]{
		return 0
	}
	for left<=right{
		pivot :=(left+right)/2
		if nums[pivot]>nums[pivot+1]{
			return pivot+1
		}else{
			if nums[pivot]<nums[left]{
				right=pivot-1
			}else{
				left=pivot+1
			}
		}
	}
	return 0
}
func binarySearch(nums []int,target,left,right int) int{
	for left<=right{
		pivot := (left + right) / 2;
		if nums[pivot]==target{
			return pivot
		}else{
			if target<nums[pivot]{
				right=pivot-1	
			}else{
				left=pivot+1
			}
		}
	}
	return -1
}

// @lc code=end

