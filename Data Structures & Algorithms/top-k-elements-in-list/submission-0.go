func topKFrequent(nums []int, k int) []int {
	type Pair struct{
		Num int
		Count int
	}
	m := make(map[int]int)
	for _, num := range nums{
		if val, ok := m[num]; ok {
			m[num] = val+1
		} else {
			m[num] = 1
		}
	}
	pairs := []Pair{}
	for n, c := range m{
		pairs = append(pairs, Pair{
			Num: n,
			Count: c,
		})
	}

	sort.Slice(pairs, func(i, j int) bool{
		return pairs[i].Count > pairs[j].Count
	})
	result := []int{}
	for i := 0; i<k; i++{
		result = append(result, pairs[i].Num)
	}
	return result
}
