package main

import "fmt"

// selectionSort sorts an array using the Selection Sort algorithm
func selectionSort(arr []int) {
    n := len(arr)
    // Traverse through all array elements
    for i := 0; i < n-1; i++ {
        // Find the minimum element in the unsorted portion of the array
        minIdx := i
        for j := i + 1; j < n; j++ {
            if arr[j] < arr[minIdx] {
                minIdx = j
            }
        }
        // Swap the found minimum element with the first element of the unsorted portion
        arr[i], arr[minIdx] = arr[minIdx], arr[i]
    }
}

func main() {
    // Example array
    arr := []int{64, 25, 12, 22, 11}
    // Sort the array using selectionSort
    selectionSort(arr)
    // Print the sorted array
    fmt.Println("Sorted array is:", arr)
}