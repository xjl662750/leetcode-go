/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 *
 * https://leetcode-cn.com/problems/valid-sudoku/description/
 *
 * algorithms
 * Medium (56.66%)
 * Likes:    248
 * Dislikes: 0
 * Total Accepted:    49.1K
 * Total Submissions: 85.1K
 * Testcase Example:  '[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]'
 *
 * 判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。
 * 
 * 
 * 数字 1-9 在每一行只能出现一次。
 * 数字 1-9 在每一列只能出现一次。
 * 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
 * 
 * 
 * 
 * 
 * 上图是一个部分填充的有效的数独。
 * 
 * 数独部分空格内已填入了数字，空白格用 '.' 表示。
 * 
 * 示例 1:
 * 
 * 输入:
 * [
 * ⁠ ["5","3",".",".","7",".",".",".","."],
 * ⁠ ["6",".",".","1","9","5",".",".","."],
 * ⁠ [".","9","8",".",".",".",".","6","."],
 * ⁠ ["8",".",".",".","6",".",".",".","3"],
 * ⁠ ["4",".",".","8",".","3",".",".","1"],
 * ⁠ ["7",".",".",".","2",".",".",".","6"],
 * ⁠ [".","6",".",".",".",".","2","8","."],
 * ⁠ [".",".",".","4","1","9",".",".","5"],
 * ⁠ [".",".",".",".","8",".",".","7","9"]
 * ]
 * 输出: true
 * 
 * 
 * 示例 2:
 * 
 * 输入:
 * [
 * ["8","3",".",".","7",".",".",".","."],
 * ["6",".",".","1","9","5",".",".","."],
 * [".","9","8",".",".",".",".","6","."],
 * ["8",".",".",".","6",".",".",".","3"],
 * ["4",".",".","8",".","3",".",".","1"],
 * ["7",".",".",".","2",".",".",".","6"],
 * [".","6",".",".",".",".","2","8","."],
 * [".",".",".","4","1","9",".",".","5"],
 * [".",".",".",".","8",".",".","7","9"]
 * ]
 * 输出: false
 * 解释: 除了第一行的第一个数字从 5 改为 8 以外，空格内其他数字均与 示例1 相同。
 * ⁠    但由于位于左上角的 3x3 宫内有两个 8 存在, 因此这个数独是无效的。
 * 
 * 说明:
 * 
 * 
 * 一个有效的数独（部分已被填充）不一定是可解的。
 * 只需要根据以上规则，验证已经填入的数字是否有效即可。
 * 给定数独序列只包含数字 1-9 和字符 '.' 。
 * 给定数独永远是 9x9 形式的。
 * 
 * 
 */

// @lc code=start
// func isValidSudoku(board [][]byte) bool {
//     for i:=0;i<9;i++{
// 		var m=make(map[byte] bool)
// 		for j:=0;j<9;j++{
			
// 		if m[board[i][j]]==true{
// 			return false
// 		}else{
// 			if string(board[i][j])!="."{
// 				m[board[i][j]]=true
// 			}
// 		}
// 		}
// 	}
// 	for i:=0;i<9;i++{
// 		var m=make(map[byte] bool)
// 		for j:=0;j<9;j++{
// 		if m[board[j][i]]==true{
// 			return false
// 		}else{
// 			if string(board[j][i])!="."{
// 			m[board[j][i]]=true
// 			}
// 		}
// 		}
// 	}
// 	for k:=0;k<=6;k=k+3{
// 		for l:=0;l<=6;l=l+3{
// 		var m=make(map[byte] bool)
// 		for i:=k;i<k+3;i++{
// 			for j:=l;j<l+3;j++{
// 				if m[board[i][j]]==true{
// 					return false
// 				}else{
// 					if string(board[i][j])!="."{
// 					m[board[i][j]]=true
// 					}
// 				}
// 			}
// 		}
// 	}
// 	}
// 	return true
// }
// func isValidSudoku(board [][]byte) bool {
// 	var row [9][]byte
// 	var columns [9][]byte
// 	var box [9][]byte
//     for i:=0;i<9;i++{
// 		for j:=0;j<9;j++{	
// 			if 	board[i][j]	!=46{
// 				row[i]=append(row[i],board[i][j])
// 				columns[j]=append(columns[j],board[i][j])
// 				box[(i/3)*3+j/3]=append(box[(i/3)*3+j/3],board[i][j])
// 			}
// 		}
// 	}
// 	for i:=0;i<9;i++{
// 		if (isValid(row[i])&&isValid(columns[i])&&isValid(box[i]))==false{
// 			 return false
// 		}
// 	}

// 	return true
// }
// func isValid(b []byte)bool{
//     for i:=0;i<len(b);i++{
// 		for j:=0;j<len(b);j++{
// 			if i!=j&&b[i]==b[j]{
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

func isValidSudoku(board [][]byte) bool {
    var row, col, block [9]uint16
    var cur uint16
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if board[i][j] == '.' {
                continue
            }
            cur = 1 << (board[i][j] - '1')  // 当前数字的 二进制数位 位置
            bi := i/3 + j/3*3  // 3x3的块索引号
            if (row[i] & cur) | (col[j] & cur) | (block[bi] & cur) != 0 { // 使用与运算查重
                return false
            }
            // 在对应的位图上，加上当前数字
            row[i] |= cur
            col[j] |= cur
            block[bi] |= cur
        }
    }
    return true
}

// @lc code=end

