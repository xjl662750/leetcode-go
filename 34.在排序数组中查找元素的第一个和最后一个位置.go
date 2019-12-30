/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 *
 * https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 *
 * algorithms
 * Medium (37.93%)
 * Likes:    269
 * Dislikes: 0
 * Total Accepted:    49.8K
 * Total Submissions: 130.3K
 * Testcase Example:  '[5,7,7,8,8,10]\n8'
 *
 * 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
 * 
 * 你的算法时间复杂度必须是 O(log n) 级别。
 * 
 * 如果数组中不存在目标值，返回 [-1, -1]。
 * 
 * 示例 1:
 * 
 * 输入: nums = [5,7,7,8,8,10], target = 8
 * 输出: [3,4]
 * 
 * 示例 2:
 * 
 * 输入: nums = [5,7,7,8,8,10], target = 6
 * 输出: [-1,-1]
 * 
 */

// @lc code=start
// func searchRange(nums []int, target int) []int {
// 	if len(nums)==0||len(nums)==1&&nums[0]!=target{
// 		return []int{-1,-1}
// 	}
// 	if len(nums)==1{
// 		return []int{0,0}
// 	}
// 	left,right:=0,len(nums)-1
// 	for left<=right{
// 		pivot := (left + right) / 2;
// 		if nums[pivot]==target{
// 			i,j:=pivot,pivot
// 			for ;j<=len(nums)-1&&nums[j]==target;j++{
// 			}
// 			for ;i>=0&&nums[i]==target;i--{
// 			}
// 			return []int{i+1,j-1}
// 		}else{
// 			if target<nums[pivot]{
// 				right=pivot-1	
// 			}else{
// 				left=pivot+1
// 			}
// 		}
// 	}
// 	return []int{-1,-1}
// }

func searchRange(nums []int, target int) []int {
	if len(nums)==0||len(nums)==1&&nums[0]!=target{
		return []int{-1,-1}
	}
	if len(nums)==1{
		return []int{0,0}
	}
	flag:=false
	resleft,resright:=-1,-1
	left,right:=0,len(nums)-1
	for left<=right{
		pivot := (left + right) / 2;
		if nums[pivot]==target{
			flag=true
			left=pivot+1
			if pivot==len(nums)-1||nums[pivot+1]>target{
				resright=pivot
				break
			}	
		}else{
			if target<nums[pivot]{
				if flag&&nums[pivot-1]==target{
					resright=pivot-1
					break
				}
				right=pivot-1	
			}else{
				left=pivot+1
			}
		}
	}
	flag=false
	left,right=0,len(nums)-1
	for left<=right{
		pivot := (left + right) / 2;
		if nums[pivot]==target{
			flag=true
			right=pivot-1
			if pivot==0||nums[pivot-1]<target{
				resleft=pivot
				break
			}	
		}else{
			if target<nums[pivot]{
				right=pivot-1	
			}else{
				if flag&&nums[pivot+1]==target{
					resleft=pivot+1
					break
				}
				left=pivot+1
			}
		}
	}
	return []int{resleft,resright}
}
// @lc code=end

