func mergeAlternately(word1 string, word2 string) string {
    var res strings.Builder
    i, j := 0, 0

    for i < len(word1) && j < len(word2) {
        res.WriteByte(word1[i])
        res.WriteByte(word2[j])
        i++
        j++
    }

    res.WriteString(word1[i:])
    res.WriteString(word2[j:])

    return res.String()
}