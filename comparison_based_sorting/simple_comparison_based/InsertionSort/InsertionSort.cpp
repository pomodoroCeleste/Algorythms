#include <iostream>
#include <vector>

// Function to perform Insertion Sort on an array
void insertionSort(std::vector<int>& arr) {
    int n = arr.size();
    // Traverse through 1 to n
    for (int i = 1; i < n; i++) {
        int key = arr[i];
        int j = i - 1;

        // Move elements of arr[0..i-1], that are greater than key, to one position ahead of their current position
        while (j >= 0 && arr[j] > key) {
            arr[j + 1] = arr[j];
            j = j - 1;
        }
        arr[j + 1] = key;
    }
}

int main() {
    // Example array
    std::vector<int> arr = {12, 11, 13, 5, 6};
    // Sort the array using insertionSort
    insertionSort(arr);
    // Print the sorted array
    std::cout << "Sorted array is: ";
    for (int i : arr) {
        std::cout << i << " ";
    }
    std::cout << std::endl;
    return 0;
}