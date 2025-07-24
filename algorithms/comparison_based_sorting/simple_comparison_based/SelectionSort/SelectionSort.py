def selection_sort(arr):
    n = len(arr)
    # Traverse through all array elements
    for i in range(n):
        # Find the minimum element in the unsorted portion of the array
        min_idx = i
        for j in range(i + 1, n):
            if arr[j] < arr[min_idx]:
                min_idx = j
        # Swap the found minimum element with the first element of the unsorted portion
        arr[i], arr[min_idx] = arr[min_idx], arr[i]

# Example array
arr = [64, 25, 12, 22, 11]
# Sort the array using selection_sort
selection_sort(arr)
# Print the sorted array
print("Sorted array is:", arr)