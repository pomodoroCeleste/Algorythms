### QuickSort Algorithm

QuickSort is a divide-and-conquer algorithm that sorts an array by partitioning it into two subarrays, sorting the subarrays independently, and then combining the results. Here is a step-by-step explanation of how the QuickSort algorithm works:

1. **Divide**:
   - Choose a pivot element from the array.
   - Partition the array into three parts: elements less than the pivot, elements equal to the pivot, and elements greater than the pivot.

2. **Conquer**:
   - Recursively apply the same process to the subarrays of elements less than and greater than the pivot.

3. **Combine**:
   - Concatenate the sorted subarrays and the pivot elements to produce the final sorted array.

### Detailed Steps

1. **Initial Call**:
   - The initial array is passed to the `quick_sort` function.

2. **Recursive Division**:
   - The array is divided into subarrays based on the pivot element until each subarray contains only one element.

3. **Combining**:
   - The subarrays are combined by concatenating the sorted subarrays and the pivot elements.

### Example

Consider the array `[10, 7, 8, 9, 1, 5]`.

1. **Divide**:
   - Choose `8` as the pivot.
   - Partition into `[7, 1, 5]` (less than `8`), `[8]` (equal to `8`), and `[10, 9]` (greater than `8`).

2. **Conquer**:
   - Recursively apply QuickSort to `[7, 1, 5]` and `[10, 9]`.
   - For `[7, 1, 5]`, choose `5` as the pivot, partition into `[1]`, `[5]`, and `[7]`.
   - For `[10, 9]`, choose `9` as the pivot, partition into `[]`, `[9]`, and `[10]`.

3. **Combine**:
   - Combine the sorted subarrays: `[1, 5, 7]`, `[8]`, and `[9, 10]` to get `[1, 5, 7, 8, 9, 10]`.

### Pseudocode

Here is a pseudocode representation of the QuickSort algorithm:

pivot = arr[length of arr // 2]
left = [x for x in arr if x < pivot]
middle = [x for x in arr if x == pivot]
right = [x for x in arr if x > pivot]

return quick_sort(left) + middle + quick_sort(right)


### Time Complexity

- **Best Case**: O(n log n)
- **Average Case**: O(n log n)
- **Worst Case**: O(n^2) (occurs when the pivot is the smallest or largest element)

### Space Complexity

- **O(log n)**: Requires additional space for the recursive call stack.

QuickSort is an efficient and widely used sorting algorithm. It is not stable, meaning that it does not preserve the relative order of equal elements. However, it is generally faster in practice compared to other O(n log n) algorithms like Merge Sort.