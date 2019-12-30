import (
	"errors"
)

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 *
 * https://leetcode-cn.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (39.92%)
 * Likes:    1187
 * Dislikes: 0
 * Total Accepted:    154.6K
 * Total Submissions: 386.1K
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
 *
 * 有效字符串需满足：
 *
 *
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 *
 *
 * 注意空字符串可被认为是有效字符串。
 *
 * 示例 1:
 *
 * 输入: "()"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: "()[]{}"
 * 输出: true
 *
 *
 * 示例 3:
 *
 * 输入: "(]"
 * 输出: false
 *
 *
 * 示例 4:
 *
 * 输入: "([)]"
 * 输出: false
 *
 *
 * 示例 5:
 *
 * 输入: "{[]}"
 * 输出: true
 *
 */

// @lc code=start
func isValid(s string) bool {
	var arr Stack
	for _, v := range s {
		if arr.IsEmpty() {
			arr.Push(string(v))
			continue
		}
		if temp, _ := arr.Top(); match(temp.(string), string(v)) {
			arr.Pop()
		} else {
			arr.Push(string(v))
		}
	}
	return arr.IsEmpty()
}
func match(a, b string) bool {
	if a == "(" && b == ")" || a == ")" && b == "(" {
		return true
	}
	if a == "[" && b == "]" || a == "]" && b == "[" {
		return true
	}
	if a == "{" && b == "}" || a == "}" && b == "{" {
		return true
	}
	return false
}

type Stack []interface{}

//Push()方法的接收器为一个Stack类型的指针（Go指针的写法与C/C++类似，类型前面加上*号）。Go语言的所有方法参数都是值传递，接收器实际也是作为方法的一个参数传递进入方法的。如果传递一切片或者数组进方法，实际是将切片或数组的值拷贝了一份传入了方法之中，此时在方法之中对该切片或数组做的操作都不会影响方法之外的原始值。如果想要方法之中的操作影响到方法外的原始值，则应该使用指针作为参数，对指针的操作会直接反应到内存之中的原始值上去。在这里我们希望更改原始值（往原始的stack之中添加数据）， 所以接收器是一个指针。方法的参数是一个interface{}类型的值，也就是说该方法可以接受任意类型作为参数。方法的实现使用了内建函数append()，往切片对尾部中添加新值。
func (stack *Stack) Push(value interface{}) {
	*stack = append(*stack, value)
}

//Top()方法返回一个任意类型的值以及一个error（是的没错，Go语言的方法可以返回多个值）。当stack为空时，返回一个空值和一个error类型的值（这里使用errors包的New()函数创建）。当stack不为空时，返回底层切片的最后一个值和一个空的error。
func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("Out of index, len is 0")
	}
	return stack[len(stack)-1], nil
}

//Pop()方法的接收器同样是一个Stack的指针。其中theStack[:len(theStack) - 1]这种写法，是Go中取子切片的方法，:两边是起始index和结束index。起始index为0时可以省略。结束的index也可以省略，省略时结束index为切片的len()值。
func (stack *Stack) Pop() (interface{}, error) {
	theStack := *stack
	if len(theStack) == 0 {
		return nil, errors.New("Out of index, len is 0")
	}
	value := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return value, nil
}

func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}

// @lc code=end
