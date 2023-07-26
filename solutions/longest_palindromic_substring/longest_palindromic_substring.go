package longestpalindromicsubstring

func longestPalindrome(s string) string {

	if len(s) <= 1 {
		return s
	}

	middles := [][2]int{}

	// find middles
	for i := 0; i < len(s)-1; i++ {

		// even
		if s[i] == s[i+1] {
			middles = append(middles, [2]int{i, i + 1})
		}

		// odd
		if i+2 < len(s) && s[i] == s[i+2] {
			middles = append(middles, [2]int{i, i + 2})
		}
	}

	if len(middles) == 0 {
		return string(s[0])
	}

	longest := ""

	// loop through middles
	for _, middle := range middles {
		front, back := middle[0], middle[1]
		palindrome := getLongestPalindrome(s, front, back)

		// update longest palindrome
		if len(palindrome) > len(longest) {
			longest = palindrome
		}
	}

	return longest
}

func getLongestPalindrome(s string, front int, back int) string {
	for front >= 0 && back < len(s) && s[front] == s[back] {
		front--
		back++
	}
	return s[front+1 : back]
}
