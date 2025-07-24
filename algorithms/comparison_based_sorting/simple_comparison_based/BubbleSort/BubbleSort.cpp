#include <iostream>
#include <vector>

// Function to perform Bubble Sort on an array
void bubbleSort(std::vector<int>& arr) {
    int n = arr.size();
    // Traverse through all array elements
    for (int i = 0; i < n-1; i++) {
        // Last i elements are already in place
        for (int j = 0; j < n-i-1; j++) {
            // Swap if the element found is greater than the next element
            if (arr[j] > arr[j+1]) {
                std::swap(arr[j], arr[j+1]);
            }
        }
    }
}

int main() {
    // Example array
    std::vector<int> arr = {64, 34, 25, 12, 22, 11, 90};
    // Sort the array using bubbleSort
    bubbleSort(arr);
    // Print the sorted array
    std::cout << "Sorted array is: ";
    for (int i : arr) {
        std::cout << i << " ";
    }
    std::cout << std::endl;
    return 0;
}
