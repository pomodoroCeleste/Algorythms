### Merge Sort Algorithm

Merge Sort is a divide-and-conquer algorithm that divides the input array into two halves, recursively sorts each half, and then merges the two sorted halves to produce the final sorted array. Here is a step-by-step explanation of how the Merge Sort algorithm works:

1. **Divide**:
   - If the array has more than one element, find the middle point to divide the array into two halves.
   - Recursively apply the same process to each half.

2. **Conquer**:
   - Once the subarrays have been divided down to single elements (which are inherently sorted), start merging them back together in sorted order.

3. **Combine**:
   - Merge the two sorted halves to produce a single sorted array.

### Detailed Steps

1. **Initial Call**:
   - The initial array is passed to the `merge_sort` function.

2. **Recursive Division**:
   - The array is divided into two halves until each subarray contains only one element.

3. **Merging**:
   - The `merge` function is called to merge two sorted subarrays into a single sorted subarray.
   - This process continues recursively until the entire array is merged and sorted.

### Example

Consider the array `[12, 11, 13, 5, 6, 7]`.

1. **Divide**:
   - Split into `[12, 11, 13]` and `[5, 6, 7]`.
   - Further split `[12, 11, 13]` into `[12]` and `[11, 13]`.
   - Further split `[11, 13]` into `[11]` and `[13]`.
   - Similarly, split `[5, 6, 7]` into `[5]` and `[6, 7]`, and then `[6, 7]` into `[6]` and `[7]`.

2. **Conquer**:
   - Merge `[11]` and `[13]` to get `[11, 13]`.
   - Merge `[12]` and `[11, 13]` to get `[11, 12, 13]`.
   - Merge `[6]` and `[7]` to get `[6, 7]`.
   - Merge `[5]` and `[6, 7]` to get `[5, 6, 7]`.

3. **Combine**:
   - Merge `[11, 12, 13]` and `[5, 6, 7]` to get the final sorted array `[5, 6, 7, 11, 12, 13]`.

### Pseudocode

Here is a pseudocode representation of the Merge Sort algorithm:

mid = length of arr / 2
left = merge_sort(arr[0:mid])
right = merge_sort(arr[mid:end])

return merge(left, right)

while i < length of left and j < length of right:
    if left[i] < right[j]:
        append left[i] to result
        i += 1
    else:
        append right[j] to result
        j += 1

append remaining elements of left to result
append remaining elements of right to result

return result

### Time Complexity

- **Best Case**: O(n log n)
- **Average Case**: O(n log n)
- **Worst Case**: O(n log n)

### Space Complexity

- **O(n)**: Requires additional space proportional to the size of the input array for the temporary arrays used during merging.

Merge Sort is a stable sorting algorithm, meaning that it preserves the relative order of equal elements in the sorted array. It is also well-suited for sorting linked lists and external sorting (sorting large amounts of data that do not fit into memory).