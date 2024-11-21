package main

import (
    "fmt"
)

// merge merges two sorted slices into a single sorted slice
func merge(left, right []int) []int {
    // Create a result slice with a capacity equal to the sum of the lengths of left and right slices
    result := make([]int, 0, len(left)+len(right))
    i, j := 0, 0

    // Compare elements from left and right slices and append the smaller one to the result
    for i < len(left) && j < len(right) {
        if left[i] < right[j]) {
            result = append(result, left[i])
            i++
        } else {
            result = append(result, right[j])
            j++
        }
    }

    // Append any remaining elements from the left slice
    result = append(result, left[i:]...)
    // Append any remaining elements from the right slice
    result = append(result, right[j:]...)

    return result
}

// mergeSort sorts a slice of integers using the merge sort algorithm
func mergeSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }

    // Find the middle point to divide the array into two halves
    mid := len(arr) / 2

    // Call mergeSort recursively on the first half
    left := mergeSort(arr[:mid])
    // Call mergeSort recursively on the second half
    right := mergeSort(arr[mid:])

    // Merge the two halves
    return merge(left, right)
}

func main() {
    arr := []int{12, 11, 13, 5, 6, 7}
    fmt.Println("Given array is:", arr)
    sortedArr := mergeSort(arr)
    fmt.Println("Sorted array is:", sortedArr)
}