package main

import (
    "fmt"
    "math"
)

// getMax returns the maximum value in the array
func getMax(arr []int) int {
    max := arr[0]
    for _, num := range arr {
        if num > max {
            max = num
        }
    }
    return max
}

// countingSort sorts the array based on the digit represented by exp
func countingSort(arr []int, exp int) {
    n := len(arr)
    output := make([]int, n) // output array
    count := make([]int, 10) // count array to store the count of occurrences of digits

    // Store count of occurrences in count[]
    for i := 0; i < n; i++ {
        index := (arr[i] / exp) % 10
        count[index]++
    }

    // Change count[i] so that count[i] contains the actual position of this digit in output[]
    for i := 1; i < 10; i++ {
        count[i] += count[i-1]
    }

    // Build the output array
    for i := n - 1; i >= 0; i-- {
        index := (arr[i] / exp) % 10
        output[count[index]-1] = arr[i]
        count[index]--
    }

    // Copy the output array to arr[], so that arr[] contains sorted numbers according to the current digit
    for i := 0; i < n; i++ {
        arr[i] = output[i]
    }
}

// radixSort sorts the array using Radix Sort algorithm
func radixSort(arr []int) {
    // Find the maximum number to know the number of digits
    max := getMax(arr)

    // Do counting sort for every digit. Note that instead of passing the digit number, exp is passed.
    // exp is 10^i where i is the current digit number
    for exp := 1; max/exp > 0; exp *= 10 {
        countingSort(arr, exp)
    }
}

func main() {
    // Example array
    arr := []int{170, 45, 75, 90, 802, 24, 2, 66}
    // Sort the array using radixSort
    radixSort(arr)
    // Print the sorted array
    fmt.Println("Sorted array is:", arr)
}