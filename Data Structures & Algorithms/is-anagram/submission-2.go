func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	map1 := make(map[rune]int)
	map2 := make(map[rune]int)
	for _, c := range []rune(s){
		val, x := map1[c]
		if x{
			map1[c] = val+1
		} else {
			map1[c] = 1
		}
	}

		for _, c := range []rune(t){
		val, x := map2[c]
		if x{
			map2[c] = val+1
		} else {
			map2[c] = 1
		}
	}

	for key, val := range map1{
		if val != map2[key] {
			return false
		}
	}
	return true
}
