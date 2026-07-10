func trap(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}
	area := 0
	prefixArr := make([]int, n)
	suffixArr := make([]int, n)

	prefixArr[0] = heights[0]
	for i := 1; i<n; i++{
		prefixArr[i] = max(prefixArr[i-1], heights[i])
	}

	suffixArr[n-1] = heights[n-1]
	for i := n-2; i>=0; i--{
		suffixArr[i] = max(suffixArr[i+1], heights[i])
	}

	for i:=0; i<n; i++{
		area += min(prefixArr[i], suffixArr[i]) - heights[i]
	}
	return area
}

func max(a, b int) int {
	if a>b{
		return a
	}
	return b
}
func min(a, b int) int {
	if a<b {
		return a
	}
	return b
}
// min(maxLeft, maxRight) - height[i]