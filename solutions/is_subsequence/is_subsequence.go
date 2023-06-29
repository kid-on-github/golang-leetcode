package subsequence

func isSubsequence(s string, t string) bool {
	sIndex := 0
	tIndex := 0

	for {
		if sIndex == len(s) {
			return true
		}

		if tIndex == len(t) {
			return false
		}

		if s[sIndex] == t[tIndex] {
			sIndex++
		}

		tIndex++
	}
}
