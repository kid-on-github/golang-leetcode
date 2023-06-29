package validpalindrome

func isPalindrome(s string) bool {
	converted := make([]byte, len(s))
	convertedLen := 0

	for i := 0; i < len(s); i++ {
		char := convertToUpperAlphaNumericByte(s[i])
		if char != 0 {
			converted[convertedLen] = char
			convertedLen++
		}
	}
	for i, char := range converted[:convertedLen/2] {
		if char != converted[(convertedLen-1)-i] {
			return false
		}
	}

	return true
}

func convertToUpperAlphaNumericByte(char byte) byte {

	if 65 <= char && char <= 90 {
		return char
	}

	if 97 <= char && char <= 122 {
		return char - 32
	}

	if 48 <= char && char <= 57 {
		return char
	}

	return 0
}
