#include <iostream>
#include <vector>
#include <algorithm>

// Function to get the maximum value in the array
int getMax(const std::vector<int>& arr) {
    return *std::max_element(arr.begin(), arr.end());
}

// Function to perform counting sort on the array based on the digit represented by exp
void countingSort(std::vector<int>& arr, int exp) {
    int n = arr.size();
    std::vector<int> output(n); // output array
    std::vector<int> count(10, 0); // count array to store the count of occurrences of digits

    // Store count of occurrences in count[]
    for (int i = 0; i < n; i++) {
        int index = (arr[i] / exp) % 10;
        count[index]++;
    }

    // Change count[i] so that count[i] contains the actual position of this digit in output[]
    for (int i = 1; i < 10; i++) {
        count[i] += count[i - 1];
    }

    // Build the output array
    for (int i = n - 1; i >= 0; i--) {
        int index = (arr[i] / exp) % 10;
        output[count[index] - 1] = arr[i];
        count[index]--;
    }

    // Copy the output array to arr[], so that arr[] contains sorted numbers according to the current digit
    for (int i = 0; i < n; i++) {
        arr[i] = output[i];
    }
}

// Function to perform Radix Sort on the array
void radixSort(std::vector<int>& arr) {
    // Find the maximum number to know the number of digits
    int max = getMax(arr);

    // Do counting sort for every digit. Note that instead of passing the digit number, exp is passed.
    // exp is 10^i where i is the current digit number
    for (int exp = 1; max / exp > 0; exp *= 10) {
        countingSort(arr, exp);
    }
}

int main() {
    // Example array
    std::vector<int> arr = {170, 45, 75, 90, 802, 24, 2, 66};
    // Sort the array using radixSort
    radixSort(arr);
    // Print the sorted array
    std::cout << "Sorted array is: ";
    for (int num : arr) {
        std::cout << num << " ";
    }
    std::cout << std::endl;
    return 0;
}