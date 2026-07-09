func majorityElement(nums []int) []int {
	m := make(map[int]int)
	for _, n := range nums{
		val, found := m[n]; if found{
			m[n] = val+1
		} else{
			m[n] = 1
		}
	}

	res := []int{}
	for key, val := range m{
		if val > len(nums)/3{
			res = append(res, key)
		}
	}
	return res
}
