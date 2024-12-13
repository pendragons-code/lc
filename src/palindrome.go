package main

func isPalindrome(x int) bool {
	// Guard against negatives
	if x < 0 {
		return false
	}

	n := x  // Assign x to n for processing
	reversedXInt := 0
	for n > 0 {
		remainder := n % 10
		reversedXInt = reversedXInt*10 + remainder
		n /= 10
	}

	return reversedXInt == x
}

//