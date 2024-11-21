package main

import "fmt"

// heapify ensures the subtree rooted at index i satisfies the max heap property
func heapify(arr []int, n, i int) {
    largest := i       // Initialize largest as root
    left := 2*i + 1    // Left child index
    right := 2*i + 2   // Right child index

    // If left child is larger than root
    if left < n && arr[left] > arr[largest] {
        largest = left
    }

    // If right child is larger than largest so far
    if right < n && arr[right] > arr[largest] {
        largest = right
    }

    // If largest is not root
    if largest != i {
        arr[i], arr[largest] = arr[largest], arr[i] // Swap
        heapify(arr, n, largest)                    // Recursively heapify the affected subtree
    }
}

// heapSort sorts an array using the Heap Sort algorithm
func heapSort(arr []int) {
    n := len(arr)

    // Build a max heap
    for i := n/2 - 1; i >= 0; i-- {
        heapify(arr, n, i)
    }

    // One by one extract elements from heap
    for i := n - 1; i >= 0; i-- {
        arr[0], arr[i] = arr[i], arr[0] // Move current root to end
        heapify(arr, i, 0)              // Call heapify on the reduced heap
    }
}

func main() {
    // Example array
    arr := []int{12, 11, 13, 5, 6, 7}
    // Sort the array using heapSort
    heapSort(arr)
    // Print the sorted array
    fmt.Println("Sorted array is:", arr)
}
