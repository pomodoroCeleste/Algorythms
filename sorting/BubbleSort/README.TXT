### Bubble Sort Algorithm

Bubble Sort is a simple comparison-based sorting algorithm that repeatedly steps through the list, compares adjacent elements, and swaps them if they are in the wrong order. The pass through the list is repeated until the list is sorted. Here is a step-by-step explanation of how the Bubble Sort algorithm works:

1. **Compare and Swap**:
   - Starting from the beginning of the list, compare each pair of adjacent elements.
   - If the elements are in the wrong order, swap them.

2. **Repeat**:
   - Repeat the process for the entire list.
   - After each pass, the largest element in the unsorted portion of the list is moved to its correct position.

3. **Stop**:
   - The algorithm stops when no more swaps are needed, indicating that the list is sorted.

### Detailed Steps

1. **Initial Call**:
   - The initial array is passed to the `bubble_sort` function.

2. **Comparison and Swapping**:
   - Compare each pair of adjacent elements and swap them if necessary.
   - Continue this process for the entire list.

3. **Repeat**:
   - Repeat the process for the entire list until no more swaps are needed.

### Example

Consider the array `[5, 1, 4, 2, 8]`.

1. **First Pass**:
   - Compare `5` and `1`, swap them to get `[1, 5, 4, 2, 8]`.
   - Compare `5` and `4`, swap them to get `[1, 4, 5, 2, 8]`.
   - Compare `5` and `2`, swap them to get `[1, 4, 2, 5, 8]`.
   - Compare `5` and `8`, no swap needed.

2. **Second Pass**:
   - Compare `1` and `4`, no swap needed.
   - Compare `4` and `2`, swap them to get `[1, 2, 4, 5, 8]`.
   - Compare `4` and `5`, no swap needed.
   - Compare `5` and `8`, no swap needed.

3. **Third Pass**:
   - Compare `1` and `2`, no swap needed.
   - Compare `2` and `4`, no swap needed.
   - Compare `4` and `5`, no swap needed.
   - Compare `5` and `8`, no swap needed.

4. **Stop**:
   - No more swaps are needed, the array is sorted: `[1, 2, 4, 5, 8]`.

### Pseudocode

Here is a pseudocode representation of the Bubble Sort algorithm:


### Time Complexity

- **Best Case**: O(n) (occurs when the array is already sorted)
- **Average Case**: O(n^2)
- **Worst Case**: O(n^2)

### Space Complexity

- **O(1)**: Requires a constant amount of additional space.

Bubble Sort is a simple and intuitive sorting algorithm, but it is not efficient for large lists. It is mainly used for educational purposes and for small datasets where its simplicity is an advantage.