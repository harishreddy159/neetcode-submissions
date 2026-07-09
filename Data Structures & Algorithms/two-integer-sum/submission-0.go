func twoSum(nums []int, target int) []int {
    hm := make(map[int]int)
	for i, val := range nums{
		diff := target - val
		v, found := hm[diff]; if found{
			return []int{v, i}
		}
		hm[val] = i
	}
	return []int{0, 0}
}
