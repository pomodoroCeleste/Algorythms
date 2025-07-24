package main

import (
    "fmt"
)

// countingSort sorts an array using the Counting Sort algorithm
func countingSort(arr []int) []int {
    if len(arr) == 0 {
        return arr
    }

    // Find the maximum and minimum elements in the array
    maxVal := arr[0]
    minVal := arr[0]
    for _, num := range arr {
        if num > maxVal {
            maxVal = num
        }
        if num < minVal {
            minVal = num
        }
    }
    rangeOfElements := maxVal - minVal + 1

    // Create a count array to store the count of each unique element
    count := make([]int, rangeOfElements)
    output := make([]int, len(arr))

    // Store the count of each element
    for _, num := range arr {
        count[num-minVal]++
    }

    // Modify the count array by adding the count of the previous element
    for i := 1; i < rangeOfElements; i++ {
        count[i] += count[i-1]
    }

    // Build the output array
    for i := len(arr) - 1; i >= 0; i-- {
        output[count[arr[i]-minVal]-1] = arr[i]
        count[arr[i]-minVal]--
    }

    // Copy the sorted elements back to the original array
    for i := 0; i < len(arr); i++ {
        arr[i] = output[i]
    }

    return arr
}

func main() {
    arr := []int{4, 2, 2, 8, 3, 3, 1}
    fmt.Println("Given array is:", arr)
    sortedArr := countingSort(arr)
    fmt.Println("Sorted array is:", sortedArr)
}