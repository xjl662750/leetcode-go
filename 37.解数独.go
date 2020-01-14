/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 *
 * https://leetcode-cn.com/problems/sudoku-solver/description/
 *
 * algorithms
 * Hard (57.51%)
 * Likes:    286
 * Dislikes: 0
 * Total Accepted:    16.2K
 * Total Submissions: 27.7K
 * Testcase Example:  '[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]'
 *
 * 编写一个程序，通过已填充的空格来解决数独问题。
 * 
 * 一个数独的解法需遵循如下规则：
 * 
 * 
 * 数字 1-9 在每一行只能出现一次。
 * 数字 1-9 在每一列只能出现一次。
 * 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
 * 
 * 
 * 空白格用 '.' 表示。
 * 
 * 
 * 
 * 一个数独。
 * 
 * 
 * 
 * 答案被标成红色。
 * 
 * Note:
 * 
 * 
 * 给定的数独序列只包含数字 1-9 和字符 '.' 。
 * 你可以假设给定的数独只有唯一解。
 * 给定数独永远是 9x9 形式的。
 * 
 * 
 */

// @lc code=start
// func solveSudoku(board [][]byte)  {
// 	// fmt.Println(board)
// 	// fmt.Println('0')
// 	// ascii :=[9]byte{1,2,3,4,5,6,7,8,9}
// 	for k:=0;k<9;k++{
// 		switch k{
// 		case 0:checkFill(board,49)
// 		case 1:checkFill(board,50)
// 		case 2:checkFill(board,51)
// 		case 3:checkFill(board,52)
// 		case 4:checkFill(board,53)
// 		case 5:checkFill(board,54)
// 		case 6:checkFill(board,55)
// 		case 7:checkFill(board,56)
// 		case 8:checkFill(board,57)
// 		}
		
// 	}
// 	total:=0
// 	for i:=0;i<9;i++{
// 		for j:=0;j<9;j++{
// 			if string(board[i][j])=="."{
// 				t:=board[i][j]
// 				n:=0
// 				for k:=0;k<9;k++{
// 					switch k{
// 					case 0:board[i][j]=49
// 					case 1:board[i][j]=50
// 					case 2:board[i][j]=51
// 					case 3:board[i][j]=52
// 					case 4:board[i][j]=53
// 					case 5:board[i][j]=54
// 					case 6:board[i][j]=55
// 					case 7:board[i][j]=56
// 					case 8:board[i][j]=57
// 					}
// 					// board[i][j]=k
// 					if isValidSudoku(board){
						
// 						t=board[i][j]
// 						n++
// 					}
// 				}
// 				if n==1{
// 					board[i][j]=t
// 					total++
// 					fmt.Println(i,j,t)
// 				}else{
// 					board[i][j]='.'
// 				}
// 			}else{
// 				total++
// 			}
// 		}
// 	}

// 	fmt.Println(total)
// 	if total!=81{
// 		solveSudoku(board)
// 	}
    
// }

// func isValidSudoku(board [][]byte) bool {
//     var row, col, block [9]uint16
//     var cur uint16
//     for i := 0; i < 9; i++ {
//         for j := 0; j < 9; j++ {
//             if board[i][j] == '.' {
//                 continue
//             }
//             cur = 1 << (board[i][j] - '1')  // 当前数字的 二进制数位 位置
//             bi := i/3 + j/3*3  // 3x3的块索引号
//             if (row[i] & cur) | (col[j] & cur) | (block[bi] & cur) != 0 { // 使用与运算查重
//                 return false
//             }
//             // 在对应的位图上，加上当前数字
//             row[i] |= cur
//             col[j] |= cur
//             block[bi] |= cur
//         }
//     }
//     return true
// }
// func checkFill(board [][]byte,n byte){
// 	var removeRow =make(map[int]bool)
// 	var removeColumn =make(map[int]bool)
// 	var removeBox =make(map[int]bool)
// 	for i := 0; i < 9; i++ {
//         for j := 0; j < 9; j++ {
//             if board[i][j]==n{
// 				boxIndex:=i/3 + j/3*3
// 				removeRow[i]=true
// 				removeColumn[j]=true
// 				removeBox[boxIndex]=true
// 			}
// 		}
// 	}
// 	// var m1 =make(map[int]int)
// 	// var m2 =make(map[int]int)
// 	m1 := map[int]int{
// 		1: -1,
// 		2: -1,
// 		3: -1,
// 		4: -1,
// 		5: -1,
// 		6: -1,
// 		7: -1,
// 		8: -1,
// 	}
// 	m2 := map[int]int{
// 		1: -1,
// 		2: -1,
// 		3: -1,
// 		4: -1,
// 		5: -1,
// 		6: -1,
// 		7: -1,
// 		8: -1,
// 	}
// 	for i := 0; i < 9; i++ {
//         for j := 0; j < 9; j++ {
// 			boxIndex:=i/3 + j/3*3
//             if (!removeRow[i])&&(!removeColumn[j])&&(!removeBox[boxIndex])&&(string(board[i][j])=="."){
// 				if m1[i]==-1{
// 					m1[i]=j
// 				}else{
// 					m1[i]=-2
// 				}
// 				if m2[j]==-1{
// 					m2[j]=i
// 				}else{
// 					m2[j]=-2
// 				}
// 			}
// 		}
// 	}
// 	for k,v:=range m1{
// 		if v!=-2&&v!=-1&&m2[v]!=-2&&m2[v]!=-1{
// 			board[k][v]=n
// 			fmt.Println(k,v,n)

// 		}
// 	}
// 	for k,v:=range m2{
// 		if v!=-2&&v!=-1&&m1[v]!=-2&&m1[v]!=-1{
// 			board[v][k]=n
// 			fmt.Println(v,k,n)
// 		}
// 	}
// }

func solveSudoku(board [][]byte) {
	lines := makeMaps()
	cols := makeMaps()
	blocks := makeMaps()

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				v := int(board[i][j]) - int('0')
				lines[i][v] = true
				cols[j][v] = true
				blocks[i/3*3+j/3][v] = true
			}
		}
	}

search:
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				// search
				mp := make(map[int]bool)
				for k, v := range lines[i] {
					if v {
						mp[k] = true
					}
				}
				for k, v := range cols[j] {
					if v {
						mp[k] = true
					}
				}
				for k, v := range blocks[i/3*3+j/3] {
					if v {
						mp[k] = true
					}
				}
				buf := []int{}
				for k := 1; k < 10; k++ {
					if !mp[k] {
						buf = append(buf, k)
					}
				}
				if len(buf) == 1 {
					board[i][j] = byte(buf[0] + '0')

					lines[i][buf[0]] = true
					cols[j][buf[0]] = true
					blocks[i/3*3+j/3][buf[0]] = true
					goto search
				}
				
			}
		}
	}

	next(board, lines, cols, blocks)
	showBoard(board)

}

func next(board [][]byte, lines, cols, blocks []map[int]bool) bool {
	// 开始回溯
	f := true
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				bi := i/3*3 + j/3
				for c := 1; c < 10; c++ {
					if !lines[i][c] && !cols[j][c] && !blocks[bi][c] {
						f = false
						board[i][j] = byte(c + '0')
						lines[i][c] = true
						cols[j][c] = true
						blocks[bi][c] = true
						if next(board, lines, cols, blocks) {
							return true
						} else {
							board[i][j] = '.'
							lines[i][c] = false
							cols[j][c] = false
							blocks[bi][c] = false
						}
					}
				}
				return false
			}

		}
	}
	return f
}

func showBoard(board [][]byte) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				fmt.Print("  ", "X")
			} else {
				fmt.Print("  ", board[i][j]-'0')
			}
		}
		fmt.Println()
	}
}

func makeMaps() []map[int]bool {
	mps := make([]map[int]bool, 9)
	for i := 0; i < 9; i++ {
		mps[i] = make(map[int]bool)
	}
	return mps
}

// [[".",".","9","7","4","8",".",".","2"],["7",".",".","6",".","2",".",".","9"],[".","2",".","1",".","9",".",".","."],[".",".","7","9","8","6","2","4","1"],["2","6","4","3","1","7","5","9","8"],["1","9","8","5","2","4","3","6","7"],["9",".",".","8","6","3",".","2","."],[".",".","2","4","9","1",".",".","6"],[".",".",".","2","7","5","9",".","."]]

// ["?",".","9","7","4","8","?",".","2"],
// ["7",".",".","6",".","2",".",".","9"],
// ["?","2","?","1",".","9","?",".","."],
// [".",".","7","9","8","6","2","4","1"],
// ["2","6","4","3","1","7","5","9","8"],
// ["1","9","8","5","2","4","3","6","7"],
// ["9",".",".","8","6","3",".","2","."],
// [".",".","2","4","9","1",".",".","6"],
// [".",".",".","2","7","5","9",".","."]]

// [[".",".","9","7","4","8",".",".","."],["7",".",".",".",".",".",".",".","."],[".","2",".","1",".","9",".",".","."],[".",".","7",".",".",".","2","4","."],[".","6","4",".","1",".","5","9","."],[".","9","8",".",".",".","3",".","."],[".",".",".","8",".","3",".","2","."],[".",".",".",".",".",".",".",".","6"],[".",".",".","2","7","5","9",".","."]]
// @lc code=end

