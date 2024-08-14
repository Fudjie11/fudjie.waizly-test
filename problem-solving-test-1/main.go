package main

import (
	"fmt"
)

// minimaxSum calculates and prints the minimum and maximum sum of 4 out of 5 elements in an array.
// The minimum sum is calculated by excluding the maximum element, and the maximum sum is calculated
// by excluding the minimum element.
func minimaxSum(arr []int) {
	// Initialize min and max values to the first element of the array
	minVal, maxVal := arr[0], arr[0]
	totalSum := 0

	// Iterate through the array to find the total sum, and the minimum and maximum values
	for _, value := range arr {
		if value < minVal {
			minVal = value
		}
		if value > maxVal {
			maxVal = value
		}
		totalSum += value
	}

	// Calculate the minimum sum by excluding the maximum value from the total sum
	// Calculate the maximum sum by excluding the minimum value from the total sum
	minSum := totalSum - maxVal
	maxSum := totalSum - minVal

	// Print the minimum and maximum sums
	fmt.Printf("Min Sum: %d, Max Sum: %d\n", minSum, maxSum)
}

func main() {
	var n int

	// Prompt the user to enter the number of elements in the array
	fmt.Print("Enter the number of elements (Example: 5/6/7/8/9/10): ")
	fmt.Scan(&n)

	// Ensure the user enters at least 5 elements, as the problem requires summing 4 out of 5 elements
	if n < 5 {
		fmt.Println("Please enter at least 5 numbers.")
		return
	}

	// Create a slice to store the input values
	arr := make([]int, n)

	// Prompt the user to enter the elements of the array
	fmt.Println("Enter the elements:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	// Call the minimaxSum function to calculate and display the result
	minimaxSum(arr)
}
