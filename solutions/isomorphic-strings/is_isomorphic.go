package isomorphicstrings

func isIsomorphic(s string, t string) bool {

	// this will keep track of ascii characters seen in string t
	tSeen := [128]bool{}
	stMap := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		sChar := s[i]
		tChar := t[i]

		tAlreadySeen := tSeen[tChar]
		tExpected, ok := stMap[sChar]

		if tAlreadySeen {
			if !ok || tChar != tExpected {
				return false
			}
		} else {
			if ok {
				return false
			}
			tSeen[tChar] = true
			stMap[sChar] = tChar
		}
	}

	return true
}
