func productExceptSelf(nums []int) []int {
	result := []int{}
		for i:= 0; i<len(nums); i++{
		product := 1
		for j := 0; j<len(nums); j++{
			if j == i {
				continue
			}
			product = product * nums[j]
		}
		result = append(result, product)
	}
	return result
}
