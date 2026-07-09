func containsNearbyDuplicate(nums []int, k int) bool {
	hm := make(map[int]int)
	for i, n := range nums{
		if j, ok := hm[n]; ok{
			if abs(i-j) <= k{
				return true
			}
		}
		hm[n] = i
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}