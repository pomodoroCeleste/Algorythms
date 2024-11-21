#include <iostream>
#include <vector>

// Function to perform Selection Sort on an array
void selectionSort(std::vector<int>& arr) {
    int n = arr.size();
    // Traverse through all array elements
    for (int i = 0; i < n-1; i++) {
        // Find the minimum element in the unsorted portion of the array
        int minIdx = i;
        for (int j = i + 1; j < n; j++) {
            if (arr[j] < arr[minIdx]) {
                minIdx = j;
            }
        }
        // Swap the found minimum element with the first element of the unsorted portion
        std::swap(arr[i], arr[minIdx]);
    }
}

int main() {
    // Example array
    std::vector<int> arr = {64, 25, 12, 22, 11};
    // Sort the array using selectionSort
    selectionSort(arr);
    // Print the sorted array
    std::cout << "Sorted array is: ";
    for (int i : arr) {
        std::cout << i << " ";
    }
    std::cout << std::endl;
    return 0;
}