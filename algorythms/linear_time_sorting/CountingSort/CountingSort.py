def counting_sort(arr):
    if len(arr) == 0:
        return arr

    # Find the maximum and minimum elements in the array
    max_val = max(arr)
    min_val = min(arr)
    range_of_elements = max_val - min_val + 1

    # Create a count array to store the count of each unique element
    count = [0] * range_of_elements
    output = [0] * len(arr)

    # Store the count of each element
    for num in arr:
        count[num - min_val] += 1

    # Modify the count array by adding the count of the previous element
    for i in range(1, range_of_elements):
        count[i] += count[i - 1]

    # Build the output array
    for num in reversed(arr):
        output[count[num - min_val] - 1] = num
        count[num - min_val] -= 1

    # Copy the sorted elements back to the original array
    for i in range(len(arr)):
        arr[i] = output[i]

    return arr

# Example usage
if __name__ == "__main__":
    arr = [4, 2, 2, 8, 3, 3, 1]
    print("Given array is:", arr)
    sorted_arr = counting_sort(arr)
    print("Sorted array is:", sorted_arr)