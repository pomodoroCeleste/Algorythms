package main

import "fmt"

// quickSort sorts an array using the QuickSort algorithm
func quickSort(arr []int, low, high int) {
	if low < high {
		// Partition the array and get the partitioning index
		pi := partition(arr, low, high)

		// Recursively sort the elements before and after partition
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

// partition partitions the array into two parts
func partition(arr []int, low, high int) int {
	pivot := arr[high] // Pivot element
	i := low - 1       // Index of smaller element

	// Traverse through all elements and rearrange them based on the pivot
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++                             // Increment index of smaller element
			arr[i], arr[j] = arr[j], arr[i] // Swap the elements
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1] // Place the pivot element in the correct position
	return i + 1                              // Return the partitioning index
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	n := len(arr)
	quickSort(arr, 0, n-1)               // Sort the array
	fmt.Println("Sorted array is:", arr) // Print the sorted array
}
