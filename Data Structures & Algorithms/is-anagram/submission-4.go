func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	seen := make(map[rune]int)
	for _, c := range s {
		_, ok := seen[c]
		if !ok {
			seen[c] = 0
		}
		seen[c] += 1
	}

	for _, c := range t {
		_, ok := seen[c]
		if !ok {
			return false
		}
		seen[c] -= 1
		if seen[c] < 0 {
			return false
		}
	}

	for _, v := range seen {
		if v != 0 {
			return false
		}
	}

	return true
}
