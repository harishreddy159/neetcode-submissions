func longestCommonPrefix(strs []string) string {
    prefix := strs[0]
	for _, str := range strs{
		fmt.Println("before ", prefix)
		prefix = getLongestPrefix(prefix, str)
		fmt.Println("after ", prefix)
	}
	return prefix
}

func getLongestPrefix(prefix string, str string) string {
	p := []rune{}
	sRunes := []rune(str)
	j := 0

	for _, i := range prefix{
		if j<len(str) && sRunes[j] == i {
			p = append(p, i)
		} else{
			break
		}
		j++
	}
	return string(p)
}
