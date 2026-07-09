func isAnagram(s string, t string) bool {
	sorted1 := sortString(s)
	sorted2 := sortString(t)
	return sorted1 == sorted2
}

func sortString(s string) string{
	runes := []rune(s)
	
	sort.Slice(runes, func(i, j int)bool{
		return runes[i] < runes[j]
	})

	return string(runes)
}