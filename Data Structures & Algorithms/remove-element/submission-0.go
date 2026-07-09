func removeElement(nums []int, val int) int {
    j := 0
	temp := 0
	for i, v := range nums{
		if v != val{
			temp = nums[j]
			nums[j] = nums[i]
			nums[i] = temp
			j++
		}
	}
	fmt.Println(nums)
	return j
}
