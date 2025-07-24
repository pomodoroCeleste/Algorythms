# Linear Time Sorting Algorithms

This directory contains implementations of sorting algorithms that can achieve linear time complexity under specific conditions. These algorithms do not use comparisons between elements but instead rely on other properties of the data.

## Algorithms Included

- **[Bucket Sort](BucketSort/)**: Distributes elements into buckets and sorts each bucket individually
- **[Counting Sort](CountingSort/)**: Counts occurrences of each distinct element
- **[Radix Sort](RadixSort/)**: Sorts by processing individual digits/characters

## Algorithm Details

### Bucket Sort
- **Time Complexity**: O(n + k) average case, O(n²) worst case
- **Space Complexity**: O(n + k)
- **Best for**: Uniformly distributed floating-point numbers
- **Requirements**: Input should be uniformly distributed over a range

### Counting Sort
- **Time Complexity**: O(n + k) where k is the range of input
- **Space Complexity**: O(k)
- **Best for**: Small range of integer values
- **Requirements**: Input must be integers in a known, small range

### Radix Sort
- **Time Complexity**: O(d × (n + k)) where d is the number of digits
- **Space Complexity**: O(n + k)
- **Best for**: Fixed-width integer or string data
- **Requirements**: Fixed number of digits/characters

## Complexity Comparison

| Algorithm | Time (Avg) | Time (Worst) | Space | Input Requirements |
|-----------|------------|--------------|-------|--------------------|
| Bucket Sort | O(n + k) | O(n²) | O(n + k) | Uniformly distributed |
| Counting Sort | O(n + k) | O(n + k) | O(k) | Integers in range [0, k] |
| Radix Sort | O(d(n + k)) | O(d(n + k)) | O(n + k) | Fixed-width data |

## When to Use

- **Bucket Sort**: When input is uniformly distributed over a continuous range
- **Counting Sort**: When sorting integers with a small, known range
- **Radix Sort**: When sorting fixed-width integers, strings, or other data

## Limitations

These algorithms are faster than comparison-based sorts but have restrictions:
- **Limited input types**: Cannot sort arbitrary data types
- **Memory requirements**: May require significant additional space
- **Input constraints**: Performance depends on specific input characteristics

## File Structure

Each algorithm directory contains:
- Implementation files in multiple languages (`.py`, `.cpp`, `.go`)
- Jupyter notebook (`readme.ipynb`) with detailed explanation
- Visual diagrams (`.png` files) showing algorithm execution
- Example inputs and outputs

## Note

While these algorithms can achieve linear time, they are specialized tools. For general-purpose sorting, comparison-based algorithms like Quick Sort or Merge Sort are often more practical.
