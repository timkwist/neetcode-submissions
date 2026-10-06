func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	seen := [26]int{}
	for i := range s {
		seen[s[i]-'a']++
		seen[t[i]-'a']--
	}

	for _, v := range seen {
		if v != 0 {
			return false
		}
	}

	return true
}
