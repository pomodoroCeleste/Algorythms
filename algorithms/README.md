# Algorithms

This directory contains implementations of fundamental algorithms organized by category. Each algorithm includes multiple language implementations, detailed documentation, and visual examples.

## Categories

### [Comparison-Based Sorting](comparison_based_sorting/)
Sorting algorithms that work by comparing elements to determine order:
- **Divide and Conquer**: Merge Sort, Quick Sort
- **Simple Methods**: Bubble Sort, Selection Sort, Insertion Sort, Heap Sort

### [Linear Time Sorting](linear_time_sorting/)
Specialized sorting algorithms that can achieve O(n) performance:
- **Bucket Sort**: For uniformly distributed data
- **Counting Sort**: For integers in a known range
- **Radix Sort**: For fixed-width data

### [Other Algorithms](other_algorithms/)
Additional fundamental algorithms:
- **Factorial**: Mathematical function computation

## Algorithm Selection Guide

### For General Sorting
- **Small datasets (< 50 elements)**: Insertion Sort
- **General purpose**: Quick Sort
- **Stability required**: Merge Sort
- **Memory constrained**: Heap Sort
- **Educational purposes**: Bubble Sort

### For Specialized Sorting
- **Integers in small range**: Counting Sort
- **Uniformly distributed floats**: Bucket Sort
- **Fixed-width integers/strings**: Radix Sort

## Complexity Overview

| Category | Algorithm | Best | Average | Worst | Space | Stable |
|----------|-----------|------|---------|-------|-------|--------|
| **Divide & Conquer** | Merge Sort | O(n log n) | O(n log n) | O(n log n) | O(n) | Yes |
| | Quick Sort | O(n log n) | O(n log n) | O(n²) | O(log n) | No |
| **Simple Comparison** | Insertion Sort | O(n) | O(n²) | O(n²) | O(1) | Yes |
| | Selection Sort | O(n²) | O(n²) | O(n²) | O(1) | No |
| | Bubble Sort | O(n) | O(n²) | O(n²) | O(1) | Yes |
| | Heap Sort | O(n log n) | O(n log n) | O(n log n) | O(1) | No |
| **Linear Time** | Counting Sort | O(n+k) | O(n+k) | O(n+k) | O(k) | Yes |
| | Radix Sort | O(d(n+k)) | O(d(n+k)) | O(d(n+k)) | O(n+k) | Yes |
| | Bucket Sort | O(n+k) | O(n+k) | O(n²) | O(n) | Yes |

## File Organization

Each algorithm directory contains:
- **Implementation files**: `.py`, `.cpp`, `.go` source code
- **Documentation**: `readme.ipynb` Jupyter notebooks
- **Visual aids**: `.png` diagrams showing execution
- **Examples**: Sample inputs and expected outputs

## Language Support

All algorithms are implemented in multiple languages:
- **Python**: Clean, readable implementations
- **C++**: Performance-optimized versions
- **Go**: Modern, concurrent-ready implementations

## Learning Path

### Beginner
1. Start with **Bubble Sort** for basic understanding
2. Learn **Insertion Sort** for practical small-data sorting
3. Study **Selection Sort** for simple algorithm design

### Intermediate
1. Master **Merge Sort** for divide-and-conquer concepts
2. Understand **Quick Sort** for efficient general-purpose sorting
3. Learn **Heap Sort** for heap data structure

### Advanced
1. Explore **Counting Sort** for specialized integer sorting
2. Study **Radix Sort** for multi-digit/character data
3. Implement **Bucket Sort** for distributed data

## Performance Testing

Each algorithm includes:
- **Correctness tests**: Verify proper sorting
- **Performance benchmarks**: Compare execution times
- **Memory usage analysis**: Understand space requirements
- **Stability tests**: Check relative order preservation

## Contributing

When adding new algorithms:
1. Follow the existing directory structure
2. Implement in all three languages (Python, C++, Go)
3. Create comprehensive Jupyter notebook documentation
4. Include visual diagrams where helpful
5. Add performance analysis and complexity discussion
6. Provide test cases and examples

## Future Algorithms

Planned additions:
- **Selection algorithms**: Quick Select, Median of Medians
- **String algorithms**: KMP, Rabin-Karp, Boyer-Moore
- **Dynamic programming**: Longest Common Subsequence, Knapsack
- **Greedy algorithms**: Activity Selection, Huffman Coding
- **Graph algorithms**: BFS, DFS, Dijkstra's, Kruskal's
