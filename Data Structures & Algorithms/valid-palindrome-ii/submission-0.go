func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r{
		if s[l] != s[r]{
			skipL := s[l+1 : r+1]
			skipR := s[l:r]
			return isPalindrome(skipL) || isPalindrome(skipR)
		}
		l++
		r--
	}
	return true
}

func isPalindrome(s string) bool{
	l, r := 0, len(s)-1
	for l < r{
		if s[l] != s[r]{
			return false
		}
		l++
		r--
	}
	return true
}
