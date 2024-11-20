def quick_sort(arr):
    # Base case: if the array has 1 or 0 elements, it is already sorted
    if len(arr) <= 1:
        return arr

    # Choose the pivot element
    pivot = arr[len(arr) // 2]

    # Partition the array into three parts: elements less than the pivot, elements equal to the pivot, and elements greater than the pivot
    left = [x for x in arr if x < pivot]
    middle = [x for x in arr if x == pivot]
    right = [x for x in arr if x > pivot]

    # Recursively apply quick_sort to the left and right parts, and concatenate the results with the middle part
    return quick_sort(left) + middle + quick_sort(right)

# Example usage
arr = [10, 7, 8, 9, 1, 5]
sorted_arr = quick_sort(arr)
print("Sorted array is:", sorted_arr)
