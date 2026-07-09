func longestConsecutive(nums []int) int {
	setMap := make(map[int]struct{})
	for _, num := range nums{
		setMap[num] = struct{}{}
	}
	longest := 0
	for num := range setMap{
		if _, ok := setMap[num-1]; !ok{
			length := 1
			for {
				if _, ok := setMap[num + length]; ok{
					length++
				} else {
					break
				}
			}
			if length > longest{
				longest = length
			}
		}
	}
	return longest
}
