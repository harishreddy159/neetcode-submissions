func rotate(nums []int, k int) {
	n := len(nums)
	if n == 0{
		return
	}
	k = k%n
	reverseArr(nums)
	reverseArr(nums[:k])
	reverseArr(nums[k:])
}

func reverseArr(nums []int){
	for i, j:= 0, len(nums)-1; i<j; i,j = i+1,j-1{
		nums[i], nums[j] = nums[j], nums[i]
	}
}
