#include <iostream>
#include <vector>

// Function to ensure the subtree rooted at index i satisfies the max heap property
void heapify(std::vector<int>& arr, int n, int i) {
    int largest = i; // Initialize largest as root
    int left = 2 * i + 1; // Left child index
    int right = 2 * i + 2; // Right child index

    // If left child is larger than root
    if (left < n && arr[left] > arr[largest])
        largest = left;

    // If right child is larger than largest so far
    if (right < n && arr[right] > arr[largest])
        largest = right;

    // If largest is not root
    if (largest != i) {
        std::swap(arr[i], arr[largest]); // Swap
        heapify(arr, n, largest); // Recursively heapify the affected subtree
    }
}

// Function to perform Heap Sort on an array
void heapSort(std::vector<int>& arr) {
    int n = arr.size();

    // Build a max heap
    for (int i = n / 2 - 1; i >= 0; i--)
        heapify(arr, n, i);

    // One by one extract elements from heap
    for (int i = n - 1; i >= 0; i--) {
        std::swap(arr[0], arr[i]); // Move current root to end
        heapify(arr, i, 0); // Call heapify on the reduced heap
    }
}

int main() {
    // Example array
    std::vector<int> arr = {12, 11, 13, 5, 6, 7};
    // Sort the array using heapSort
    heapSort(arr);
    // Print the sorted array
    std::cout << "Sorted array is: ";
    for (int i : arr) {
        std::cout << i << " ";
    }
    std::cout << std::endl;
    return 0;
}
