package main

import (
    "fmt"
)

// bubbleSort sorts an array using the Bubble Sort algorithm
func bubbleSort(arr []int) {
    n := len(arr)
    // Traverse through all array elements
    for i := 0; i < n-1; i++ {
        // Last i elements are already in place
        for j := 0; j < n-i-1; j++ {
            // Swap if the element found is greater than the next element
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}

func main() {
    // Example array
    arr := []int{64, 34, 25, 12, 22, 11, 90}
    // Sort the array using bubbleSort
    bubbleSort(arr)
    // Print the sorted array
    fmt.Println("Sorted array is:", arr)
}
