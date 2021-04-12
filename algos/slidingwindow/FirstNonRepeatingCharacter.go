package slidingwindow

// FirstNonRepeatingCharacter finds the first character in a string which does not repeat itself
func FirstNonRepeatingCharacter(s string) string {

	if len(s) == 0 {
		return ""
	}

	window := make(map[byte]int)

	for i := 0; i < len(s); i++ {

		if val, ok := window[s[i]]; ok {
			window[s[i]] = val + 1
		} else {
			window[s[i]] = 1
		}

	}

	for i := 0; i < len(s); i++ {

		if window[s[i]] == 1 {
			return string(s[i])
		}

	}

	return ""
}
