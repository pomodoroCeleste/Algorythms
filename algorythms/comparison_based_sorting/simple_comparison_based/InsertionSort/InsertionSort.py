def insertion_sort(arr):
    n = len(arr)
    # Traverse through 1 to n
    for i in range(1, n):
        key = arr[i]
        j = i - 1

        # Move elements of arr[0..i-1], that are greater than key, to one position ahead of their current position
        while j >= 0 and arr[j] > key:
            arr[j + 1] = arr[j]
            j -= 1
        arr[j + 1] = key

# Example array
arr = [12, 11, 13, 5, 6]
# Sort the array using insertion_sort
insertion_sort(arr)
# Print the sorted array
print("Sorted array is:", arr)