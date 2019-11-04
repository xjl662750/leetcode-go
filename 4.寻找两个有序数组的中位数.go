/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个有序数组的中位数
 */
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var arr []int
	var i,j,f int
	fmt.Println(len(nums1))
	fmt.Println(len(nums2))
	l:=len(nums1)+len(nums2)
	if len(nums1)==0{
		arr=nums2
	}else if len(nums2)==0{
		arr=nums1
	}else{
	for k:=0;k<l;k++{
		if f!=1&&(nums1[i]>nums2[j]||f==2){
			arr=append(arr,nums2[j])
			if j<len(nums2)-1{
			j++
			}else{
				f=1
			}
		}else if f!=2&&(nums1[i]<=nums2[j]||f==1){
			arr=append(arr,nums1[i])
			if i<len(nums1)-1{
				i++
				}else{
					f=2
				}
		}
		
	}
}
	if l%2==0{
		return (float64(arr[l/2])+float64(arr[l/2-1]))/2
	}else{
		return float64(arr[l/2])
	}
	return 0.0
}

