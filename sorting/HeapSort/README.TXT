### HeapSort Algorithm

HeapSort is a comparison-based sorting algorithm that uses a binary heap data structure to sort elements. It works by building a max heap from the input data, then repeatedly extracting the maximum element from the heap and rebuilding the heap until all elements are sorted. Here is a step-by-step explanation of how the HeapSort algorithm works:

1. **Build a Max Heap**:
   - Convert the input array into a max heap, where the largest element is at the root.

2. **Extract Elements**:
   - Repeatedly extract the maximum element from the heap and move it to the end of the array.
   - Reduce the size of the heap and rebuild the heap to maintain the max heap property.

### Detailed Steps

1. **Initial Call**:
   - The initial array is passed to the `heap_sort` function.

2. **Build a Max Heap**:
   - Convert the input array into a max heap by calling the `heapify` function on each non-leaf node, starting from the bottom.

3. **Extract Elements**:
   - Swap the root (maximum element) with the last element of the heap.
   - Reduce the size of the heap by one and call `heapify` on the root to rebuild the heap.
   - Repeat this process until the heap is empty.

### Example

Consider the array `[4, 10, 3, 5, 1]`.

1. **Build a Max Heap**:
   - Convert the array into a max heap: `[10, 5, 3, 4, 1]`.

2. **Extract Elements**:
   - Swap `10` with `1` to get `[1, 5, 3, 4, 10]`.
   - Rebuild the heap to get `[5, 4, 3, 1, 10]`.
   - Swap `5` with `1` to get `[1, 4, 3, 5, 10]`.
   - Rebuild the heap to get `[4, 1, 3, 5, 10]`.
   - Swap `4` with `1` to get `[1, 3, 4, 5, 10]`.
   - Rebuild the heap to get `[3, 1, 4, 5, 10]`.
   - Swap `3` with `1` to get `[1, 3, 4, 5, 10]`.
   - Rebuild the heap to get `[1, 3, 4, 5, 10]`.

### Pseudocode

Here is a pseudocode representation of the HeapSort algorithm:

# Build a max heap
for i from n/2 - 1 to 0:
    heapify(arr, n, i)

# Extract elements from heap one by one
for i from n-1 to 0:
    swap arr[0] with arr[i]
    heapify(arr, i, 0)

    if left < n and arr[left] > arr[largest]:
    largest = left

if right < n and arr[right] > arr[largest]:
    largest = right

if largest != i:
    swap arr[i] with arr[largest]
    heapify(arr, n, largest)


### Time Complexity

- **Best Case**: O(n log n)
- **Average Case**: O(n log n)
- **Worst Case**: O(n log n)

### Space Complexity

- **O(1)**: Requires a constant amount of additional space.

HeapSort is an efficient sorting algorithm with a time complexity of O(n log n) in all cases. It is not a stable sorting algorithm, meaning that it does not preserve the relative order of equal elements. However, it is widely used due to its efficiency and simplicity.