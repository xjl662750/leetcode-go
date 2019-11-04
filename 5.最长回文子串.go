/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */
func longestPalindrome(s string) string {
	// fmt.Println(len(s))
	var m=make(map[int]string)
	if len(s)>1{
    for k:=0;k<len(s)-1;k++{
			m[k]=string(s[k])
			for s[k]==s[k+1]{
					m[k]=m[k]+string(s[k])
					s=string(s[0:k+1])+string(s[k+2:])
					if k==len(s)-1{
						break
					}
			}
			if k==len(s)-2{
				m[k+1]=string(s[k+1])
				break
			}
		}
	} else if len(s)==1{
		m[0]=string(s[0])
	}
	// fmt.Println(m)
		var r string
		var str string
		// fmt.Println(string(v)+string(s[2:4]))
		// if k==0||k==len(s)-1{
		// 	r=string(v)
		// }else{
			for i:=0;i<len(s);i++{
				r=string(m[i])
				for j:=1;i-j>=0&&i+j<=len(s)-1;j++{
					if s[i-j]==s[i+j]{
						if len(m[i-j])>len(m[i+j]){
							r=string(m[i+j])+r+string(m[i+j])
							break
						}else if len(m[i-j])<len(m[i+j]){
							r=string(m[i-j])+r+string(m[i-j])
							break
						}else{
							r=string(m[i-j])+r+string(m[i-j])
						}
					}else{
						// r=string(m[i])
						break
					}
				}
				fmt.Println(r)
				if len(str)<len(r){
					str=r
				}
			}


	return str

}

