package main

func twoSum(nums []int, target int) []int {
	// Create a map to store the number as the key and its index as the value
	numMap := make(map[int]int)

	// Loop through the array once
	for i, num := range nums {
		complement := target - num

		// Check if the complement is already in the map
		if index, found := numMap[complement]; found {
			// If found, return the indices of the complement and the current number
			return []int{index, i}
		}

		// Store the current number and its index in the map
		numMap[num] = i
	}

	// If no solution is found, return nil
	return nil
}

// honestly dude the hard part of the question was figuring out how to make it less than O(n^2) in a lang im new to
