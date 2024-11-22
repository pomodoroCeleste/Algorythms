def merge_sort(arr):
    # Base case: if the array has 1 or 0 elements, it is already sorted
    if len(arr) <= 1:
        return arr

    # Find the middle point to divide the array into two halves
    mid = len(arr) // 2

    # Call merge_sort recursively on the first half
    left = merge_sort(arr[:mid])

    # Call merge_sort recursively on the second half
    right = merge_sort(arr[mid:])

    # Merge the two halves
    return merge(left, right)

def merge(left, right):
    result = []
    i = j = 0

    # Compare elements from left and right arrays and append the smaller one to the result
    while i < len(left) and j < len(right):
        if left[i] < right[j]:
            result.append(left[i])
            i += 1
        else:
            result.append(right[j])
            j += 1

    # Append any remaining elements from the left array
    result.extend(left[i:])
    # Append any remaining elements from the right array
    result.extend(right[j:])
    return result

# Example usage
if __name__ == "__main__":
    arr = [12, 11, 13, 5, 6, 7]
    print("Given array is:", arr)
    sorted_arr = merge_sort(arr)
    print("Sorted array is:", sorted_arr)
