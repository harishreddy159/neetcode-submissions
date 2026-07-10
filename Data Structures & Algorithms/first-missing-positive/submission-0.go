func firstMissingPositive(nums []int) int {
	n := len(nums)
	seen := make([]bool, n)
	
	for _, num := range nums{
		if num>0 && num<=n{
			seen[num-1] = true
		}
	}

	for i:=0; i<n; i++{
		if !seen[i]{
			return i + 1
		}
	}
	return n+1
}
