/*
 * @lc app=leetcode.cn id=30 lang=golang
 *
 * [30] 串联所有单词的子串
 *
 * https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/description/
 *
 * algorithms
 * Hard (27.43%)
 * Likes:    180
 * Dislikes: 0
 * Total Accepted:    18.5K
 * Total Submissions: 65.6K
 * Testcase Example:  '"barfoothefoobarman"\n["foo","bar"]'
 *
 * 给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
 * 
 * 注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：
 * ⁠ s = "barfoothefoobarman",
 * ⁠ words = ["foo","bar"]
 * 输出：[0,9]
 * 解释：
 * 从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
 * 输出的顺序不重要, [9,0] 也是有效答案。
 * 
 * 
 * 示例 2：
 * 
 * 输入：
 * ⁠ s = "wordgoodgoodgoodbestword",
 * ⁠ words = ["word","good","best","word"]
 * 输出：[]
 * 
 * 
 */
// @lc code=start
// func findSubstring(s string, words []string) []int {
// 	if len(words)==0{
// 		return nil
// 	}
// 	sLen:=len(s)
// 	wordsLen:=len(words)
// 	var m=make(map[string]int)
// 	for _,v:=range words{
// 		m[v]=m[v]+1
// 	}
// 	var res []int
// 	length:=len(words[0])
//     for k,_:=range s{
// 		t:=cloneTags(m)
// 		if k<=sLen-length*wordsLen&&t[s[k:k+length]]!=0{
// 			n:=0
// 			for i:=k;i<=sLen-length&&t[s[i:i+length]]!=0;i=i+length{
// 				n++
// 				t[s[i:i+length]]=t[s[i:i+length]]-1
// 			}
// 			if n==wordsLen{
// 				res=append(res,k)
// 			}
// 		}
// 	}
// 	return res
// }
// func cloneTags(tags map[string]int) map[string]int {
// 	cloneTags := make(map[string]int)
// 	for k, v := range tags {
// 	   cloneTags[k] = v
// 	}
// 	return cloneTags
//  }

//  func findSubstring(s string, words []string) (result []int) {
// 	n := len(words)
// 	if n == 0 {
// 		return
// 	}
// 	dict := map[string]int{}
// 	wn := 0
// 	for _, w := range words {
// 		dict[w]++
// 		wn += len(w)
// 	}
// 	ns := len(s)
// 	for i := 0; i < ns-wn+1; i++ {
// 		if findIndexes(s[i:i+wn], len(words[0]), dict) {
// 			result = append(result, i)
// 		}
// 	}
// 	return
// }

// func findIndexes(s string, wl int, dict map[string]int) bool {
// 	ns := len(s)
// 	tmp := map[string]int{}
// 	for i := 0; i < ns-wl+1; i++ {
// 		k := s[i : i+wl]
// 		if dict[k] != 0 {
// 			tmp[k]++
//             // 从末尾开始继续匹配，防止单词重叠
// 			i = i + wl - 1
// 		}
// 	}
// 	for k, v := range dict {
// 		if tmp[k] != v {
// 			return false
// 		}
// 	}
// 	return true
// }
func findSubstring(s string, words []string) []int {
    w2num := make(map[string]uint8)
    allNum := len(words)
    if allNum == 0 {
        return []int{}
    }
    length :=  len(words[0])
    if length == 0 {
        return []int{}
    }
    for _, str := range(words) {
        w2num[str] += 1
    }
    tmp := (make([]int, 0, 30))
    re := &tmp
    for i:=0; i<len(s)-allNum*length+1 && i<length; i++ {
        isSub(s[i:], w2num, length, allNum, re, 0, 0, i) 
    }
    return *re
}

func isSub(s string, w2num map[string]uint8, length int, allNum int, re *[]int, start int, end int, offset int) {
    if len(s[end:]) - allNum*length < 0 {
        return
    }
    tmp, ok := w2num[s[end:end+length]]
    if ok && tmp>0 {
        //能进下一状态
        w2num[s[end:end+length]] -= 1
        allNum -= 1
        end += length
        if allNum == 0 {
            *re = append(*re, start+offset)
            w2num[s[start:start+length]] += 1
            allNum += 1
            start += length
            isSub(s, w2num, length, allNum, re, start, end, offset)
            start -= length
            allNum -= 1
            w2num[s[start:start+length]] -= 1
        } else {
            isSub(s, w2num, length, allNum, re, start, end, offset)
        }
        end -= length
        allNum += 1
        w2num[s[end:end+length]] += 1
    } else {
        if end > start {
            w2num[s[start:start+length]] += 1
            allNum += 1
            start += length
            isSub(s, w2num, length, allNum, re, start, end, offset)
            start -= length
            allNum -= 1
            w2num[s[start:start+length]] -= 1
        } else {
            isSub(s, w2num, length, allNum, re, start+length, end+length, offset)
        }
    }
}

// @lc code=end

