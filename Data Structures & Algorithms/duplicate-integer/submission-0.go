func hasDuplicate(nums []int) bool {
    countMap := make(map[int]int)
    for _, n := range nums{
        _, valid := countMap[n]
        if valid{
            return true
        } else {
            countMap[n] = 1
        }
    }
    return false
}
