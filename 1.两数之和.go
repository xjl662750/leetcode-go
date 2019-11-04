/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */
// func twoSum(nums []int, target int) []int {
// 	var result []int
//     for i:=0;i<len(nums);i++{
// 		for j:=i+1;j<len(nums);j++{
// if nums[i]+nums[j]==target{
// 	result=append(result,i)
// 	result=append(result,j)
// 	break
// }
// 		}
// 	}
// 	return result
// }
// func twoSum(nums []int, target int) []int {
// 	s := map[int]int{}
//     for i:=0;i<len(nums);i++{
// 		if v,ok:=s[target-nums[i]];ok{
// 			return []int{v, i}
// 		}
// 		s[nums[i]]=i
// 	}
// 	return []int{}
// }
func twoSum(nums []int, target int) []int {

	numsMap := make(map[int]int)
	for i, v := range nums {

		if d, ok := numsMap[target - v]; ok {
			return []int{d, i}
		}
		numsMap[v] = i

	}

	return []int{}
}

