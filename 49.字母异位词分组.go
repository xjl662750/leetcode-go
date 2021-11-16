/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, str := range strs {
	  k := strArray(str)
	  s, ok := m[k]
	  if !ok {
		s = make([]string, 0)
	  }
	  s = append(s, str)
	  m[k] = s
	}
	res := make([][]string, 0, len(m))
	for _, v := range m {
	  res = append(res, v)
	}
	return res
  }
  // 用作计算每个单词的Key
  func strArray(s string) [26]int {
	res := [26]int{}
	for _, v := range s {
	  res[v - 'a']++
	}
	return res
  }
  
// 素数法
// func groupAnagrams(strs []string) [][]string {
// 	prime := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103}

// 	var key_val int
// 	lookup := map[int][]string{}

// 	for i:=0; i<len(strs); i++ {
// 		key_val = 1
// 		for _, s:= range []rune(strs[i]) {
// 			key_val *= prime[s-97]
// 		}
// 		lookup[key_val] = append(lookup[key_val], strs[i])
// 	}
// 	res := [][]string{}
// 	for _, v := range lookup {
// 		res = append(res, v)
// 	}
// 	return res
// }
// @lc code=end

