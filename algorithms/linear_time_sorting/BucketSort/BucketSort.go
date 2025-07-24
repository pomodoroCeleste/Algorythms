package main

import (
    "fmt"
    "sort"
)

// bucketSort sorts an array using the Bucket Sort algorithm
func bucketSort(arr []float64) []float64 {
    if len(arr) == 0 {
        return arr
    }

    // Create n empty buckets
    n := len(arr)
    buckets := make([][]float64, n)

    // Put array elements in different buckets
    for _, num := range arr {
        index := int(num * float64(n)) // Index in bucket
        buckets[index] = append(buckets[index], num)
    }

    // Sort individual buckets
    for i := range buckets {
        sort.Float64s(buckets[i])
    }

    // Concatenate all buckets into arr
    sortedArr := make([]float64, 0, len(arr))
    for _, bucket := range buckets {
        sortedArr = append(sortedArr, bucket...)
    }

    return sortedArr
}

func main() {
    arr := []float64{0.78, 0.17, 0.39, 0.26, 0.72, 0.94, 0.21, 0.12, 0.23, 0.68}
    fmt.Println("Given array is:", arr)
    sortedArr := bucketSort(arr)
    fmt.Println("Sorted array is:", sortedArr)
}