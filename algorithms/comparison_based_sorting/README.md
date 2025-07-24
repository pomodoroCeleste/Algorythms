# Comparison-Based Sorting Algorithms

This directory contains implementations of various comparison-based sorting algorithms. These algorithms work by comparing elements to determine their relative order.

## Algorithms Included

### Divide and Conquer
- **[Merge Sort](divide_and_conquer/MergeSort/)**: A stable sorting algorithm with guaranteed O(n log n) performance
- **[Quick Sort](divide_and_conquer/QuickSort/)**: An efficient in-place sorting algorithm with average O(n log n) performance

### Simple Comparison-Based
- **[Bubble Sort](simple_comparison_based/BubbleSort/)**: A simple comparison sort with O(n²) time complexity
- **[Heap Sort](simple_comparison_based/HeapSort/)**: An efficient sorting algorithm using the heap data structure
- **[Insertion Sort](simple_comparison_based/InsertionSort/)**: Efficient for small datasets and nearly sorted arrays
- **[Selection Sort](simple_comparison_based/SelectionSort/)**: Simple algorithm that repeatedly finds the minimum element

## Complexity Comparison

| Algorithm | Best Case | Average Case | Worst Case | Space | Stable |
|-----------|-----------|--------------|------------|-------|--------|
| Merge Sort | O(n log n) | O(n log n) | O(n log n) | O(n) | Yes |
| Quick Sort | O(n log n) | O(n log n) | O(n²) | O(log n) | No |
| Heap Sort | O(n log n) | O(n log n) | O(n log n) | O(1) | No |
| Insertion Sort | O(n) | O(n²) | O(n²) | O(1) | Yes |
| Selection Sort | O(n²) | O(n²) | O(n²) | O(1) | No |
| Bubble Sort | O(n) | O(n²) | O(n²) | O(1) | Yes |

## File Structure

Each algorithm directory contains:
- Implementation files in multiple languages (`.py`, `.cpp`, `.go`)
- Jupyter notebook (`readme.ipynb`) with detailed explanation
- Visual diagrams (`.png` files) showing algorithm execution
- Example inputs and outputs

## Usage

Navigate to any algorithm directory to find:
1. **Documentation**: Open the `readme.ipynb` file for detailed explanations
2. **Implementations**: Choose your preferred programming language
3. **Visual Aids**: View PNG files to understand algorithm execution

## When to Use Each Algorithm

- **Merge Sort**: When you need guaranteed O(n log n) performance and stability
- **Quick Sort**: For general-purpose sorting with good average performance
- **Heap Sort**: When you need O(n log n) with minimal extra space
- **Insertion Sort**: For small arrays or nearly sorted data
- **Selection Sort**: When memory writes are expensive
- **Bubble Sort**: For educational purposes only
