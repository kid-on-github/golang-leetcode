package longestpalindrome

// this file contains 3 implementations

// use a set to keep track of odd letters
// (Go doesn't have sets so I used a hashmap)
func longestPalindrome(s string) int {
	oddLetters := make(map[rune]bool)

	for _, letter := range s {
		_, ok := oddLetters[letter]
		if ok {
			delete(oddLetters, letter)
		} else {
			oddLetters[letter] = true
		}
	}

	oddCount := len(oddLetters)
	if oddCount > 1 {
		return len(s) - (oddCount - 1)
	}

	return len(s)
}

// keep track of which letters are odd via bool

// func longestPalindrome(s string) int {
// 	letterIsOdd := make(map[rune]bool)

// 	for _, letter := range s {
// 		letterIsOdd[letter] = !letterIsOdd[letter]
// 	}

// 	oddCount := 0

// 	for _, odd := range letterIsOdd {
// 		if odd {
// 			oddCount += 1
// 		}
// 	}

// 	if oddCount > 1 {
// 		return len(s) - (oddCount - 1)
// 	}

// 	return len(s)
// }

// counting letters to determine which letters are odd

// func longestPalindrome(s string) (length int) {
//     letterCount := make(map[rune]int)

// 	for _, letter := range s{
// 		letterCount[letter] = 1 + letterCount[letter]
// 	}

// 	oddUsed := false

// 	for _, count := range letterCount{
// 		countIsOdd := count%2 == 1
// 		length += count

// 		if countIsOdd{
// 			if oddUsed{
// 				length -= 1
// 			} else {
// 				oddUsed = true
// 			}
// 		}
// 	}

// 	return
// }
