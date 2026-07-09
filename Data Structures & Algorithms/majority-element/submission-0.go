func majorityElement(nums []int) int {
    cm := make(map[int]int)
	for _, x := range nums{
		val, found := cm[x]; if found{
			cm[x] = val+1
		}else{
			cm[x] = 1
		}
	}
	for key, val := range cm{
		fmt.Println("key ", key, "value ", val)
		if val >= len(nums)/2{
			return key
		}
	}
	return 0
}
