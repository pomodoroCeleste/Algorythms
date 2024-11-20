#include <iostream>
#include <vector>

// Function to partition the array into two parts
int partition(std::vector<int>& arr, int low, int high) {
    int pivot = arr[high]; // Pivot element
    int i = (low - 1); // Index of smaller element

    // Traverse through all elements and rearrange them based on the pivot
    for (int j = low; j <= high - 1; j++) {
        if (arr[j] < pivot) {
            i++; // Increment index of smaller element
            std::swap(arr[i], arr[j]); // Swap the elements
        }
    }
    std::swap(arr[i + 1], arr[high]); // Place the pivot element in the correct position
    return (i + 1); // Return the partitioning index
}

// Function to implement QuickSort
void quickSort(std::vector<int>& arr, int low, int high) {
    if (low < high) {
        // Partition the array and get the partitioning index
        int pi = partition(arr, low, high);

        // Recursively sort the elements before and after partition
        quickSort(arr, low, pi - 1);
        quickSort(arr, pi + 1, high);
    }
}

int main() {
    std::vector<int> arr = {10, 7, 8, 9, 1, 5};
    int n = arr.size();
    quickSort(arr, 0, n - 1);
    std::cout << "Sorted array is: ";
    for (int i : arr) {
        std::cout << i << " ";
    }
    std::cout << std::endl;
    return 0;
}
