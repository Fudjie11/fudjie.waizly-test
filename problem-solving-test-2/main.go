package main

import (
	"fmt"
)

// calculateOccurranceValue calculates the proportions of positive, negative, and zero values
// in an array. It returns three float32 values representing the proportions of positive,
// negative, and zero elements respectively.
func calculateOccurranceValue(arr []int) (float32, float32, float32) {
	// Initialize counters for positive, negative, and zero values
	var (
		occurrancePositiveValue float32
		occurranceNegativeValue float32
		occurranceZeroValue     float32
	)

	// Iterate over the array to count occurrences of positive, negative, and zero values
	for _, number := range arr {
		if number > 0 {
			occurrancePositiveValue++
		} else if number < 0 {
			occurranceNegativeValue++
		} else {
			occurranceZeroValue++
		}
	}

	// Calculate the proportion of positive values in the array
	positiveValue := occurrancePositiveValue / float32(len(arr))

	// Calculate the proportion of negative values in the array
	negativeValue := occurranceNegativeValue / float32(len(arr))

	// Calculate the proportion of zero values in the array
	zeroValue := occurranceZeroValue / float32(len(arr))

	// Return the calculated proportions
	return positiveValue, negativeValue, zeroValue
}

func main() {
	var n int

	// Prompt the user to enter the number of elements in the array
	fmt.Print("Enter the number of elements: ")
	fmt.Scan(&n)

	// Ensure the user enters at least 1 element
	if n < 1 {
		fmt.Println("Please enter at least 1 number.")
		return
	}

	// Create a slice to store the input values
	arr := make([]int, n)

	// Prompt the user to enter the elements of the array
	fmt.Println("Enter the elements:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		// Validate that the input numbers are within the specified range (-100 to 100)
		if arr[i] < -100 || arr[i] > 100 {
			fmt.Println("Wrong Number!!!")
			return
		}
	}

	// Display the entered array
	fmt.Println("Array:", arr)

	// Calculate the proportions of positive, negative, and zero values
	positiveValue, negativeValue, zeroValue := calculateOccurranceValue(arr)

	// Print the calculated proportions
	fmt.Printf("Proportion of Positive Values: %f, Proportion of Negative Values: %f, Proportion of Zero Values: %f\n", positiveValue, negativeValue, zeroValue)
}
