func reverseString(s []byte) {
	l, h := 0, len(s)-1
	for l<h {
		s[l], s[h] = s[h], s[l]
		l++
		h--
	}
}
