package main

func lengthOfLongestSubstring(s string) int {
	// bro i was tryna figure out how to remove a null byte and then i learnt that they are not null terminated in golang

	// today i also learnt that golang does not have sets

	// Create a map to store the characters that we have seen in the current window
	seen := make(map[rune]int)
	left := 0      // The left pointer of the sliding window
	maxLength := 0 // To store the length of the longest substring

	// Iterate through the string
	for right, char := range s {
		// If the character was already seen and its index is greater than or equal to the left pointer,
		// move the left pointer to the right of the last occurrence of this character
		if lastIndex, found := seen[char]; found && lastIndex >= left {
			left = lastIndex + 1
		}

		// Update the last index of the current character
		seen[char] = right

		// Calculate the current window length and update the maximum length if needed
		windowLength := right - left + 1
		if windowLength > maxLength {
			maxLength = windowLength
		}
	}

	// Return the maximum length found
	return maxLength
}
