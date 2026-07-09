func isPalindrome(s string) bool {
	l, h := 0, len(s)-1
	for l<h {
		for l<h && !isAlphaNum(rune(s[l])){
			l++
		}
		for h>l && !isAlphaNum(rune(s[h])){
			h--
		}
		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[h])){
			return false
		}
		l++
		h--
	}
	return true
}

func isAlphaNum(c rune) bool{
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}