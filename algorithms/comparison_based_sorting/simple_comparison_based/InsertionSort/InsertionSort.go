package main

import "fmt"

// insertionSort sorts an array using the Insertion Sort algorithm
func insertionSort(arr []int) {
    n := len(arr)
    // Traverse through 1 to n
    for i := 1; i < n; i++ {
        key := arr[i]
        j := i - 1

        // Move elements of arr[0..i-1], that are greater than key, to one position ahead of their current position
        for j >= 0 && arr[j] > key {
            arr[j+1] = arr[j]
            j = j - 1
        }
        arr[j+1] = key
    }
}

func main() {
    // Example array
    arr := []int{12, 11, 13, 5, 6}
    // Sort the array using insertionSort
    insertionSort(arr)
    // Print the sorted array
    fmt.Println("Sorted array is:", arr)
}