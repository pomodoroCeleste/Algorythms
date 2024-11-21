#include <iostream>
#include <vector>
#include <algorithm>

// Function to perform counting sort
void countingSort(std::vector<int>& arr) {
    if (arr.empty()) return;

    // Find the maximum and minimum elements in the array
    int max_val = *std::max_element(arr.begin(), arr.end());
    int min_val = *std::min_element(arr.begin(), arr.end());
    int range = max_val - min_val + 1;

    // Create a count array to store the count of each unique element
    std::vector<int> count(range, 0);
    std::vector<int> output(arr.size());

    // Store the count of each element
    for (int num : arr) {
        count[num - min_val]++;
    }

    // Modify the count array by adding the count of the previous element
    for (int i = 1; i < range; i++) {
        count[i] += count[i - 1];
    }

    // Build the output array
    for (int i = arr.size() - 1; i >= 0; i--) {
        output[count[arr[i] - min_val] - 1] = arr[i];
        count[arr[i] - min_val]--;
    }

    // Copy the sorted elements back to the original array
    for (int i = 0; i < arr.size(); i++) {
        arr[i] = output[i];
    }
}

int main() {
    std::vector<int> arr = {4, 2, 2, 8, 3, 3, 1};
    std::cout << "Given array is: ";
    for (int num : arr) {
        std::cout << num << " ";
    }
    std::cout << std::endl;

    countingSort(arr);

    std::cout << "Sorted array is: ";
    for (int num : arr) {
        std::cout << num << " ";
    }
    std::cout << std::endl;

    return 0;
}