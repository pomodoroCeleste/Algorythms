# Algorithms and Data Structures

A comprehensive repository containing implementations of fundamental algorithms and data structures in multiple programming languages (Python, C++, Go) with detailed documentation and visual examples.

## Repository Structure

```
├── algorithms/                          # Algorithm implementations
│   ├── comparison_based_sorting/        # Comparison-based sorting algorithms
│   │   ├── divide_and_conquer/          # Divide and conquer sorting
│   │   │   ├── MergeSort/               # Merge Sort implementation
│   │   │   └── QuickSort/               # Quick Sort implementation
│   │   └── simple_comparison_based/     # Simple comparison sorts
│   │       ├── BubbleSort/              # Bubble Sort implementation
│   │       ├── HeapSort/                # Heap Sort implementation
│   │       ├── InsertionSort/           # Insertion Sort implementation
│   │       └── SelectionSort/           # Selection Sort implementation
│   ├── linear_time_sorting/             # Linear time sorting algorithms
│   │   ├── BucketSort/                  # Bucket Sort implementation
│   │   ├── CountingSort/                # Counting Sort implementation
│   │   └── RadixSort/                   # Radix Sort implementation
│   └── other_algorithms/                # Other useful algorithms
│       └── factorial/                   # Factorial calculation
├── data_structures/                     # Data structure implementations
│   ├── stacks/                          # Stack operations
│   │   ├── empty/                       # Check if stack is empty
│   │   ├── operations/                  # General stack operations
│   │   ├── pop/                         # Pop operation
│   │   └── push/                        # Push operation
│   ├── queues/                          # Queue operations
│   │   ├── operations/                  # Basic queue operations
│   │   │   ├── accoda/                  # Enqueue operation
│   │   │   ├── extract/                 # Dequeue operation
│   │   │   └── init/                    # Queue initialization
│   │   └── priority/                    # Priority queue operations
│   │       ├── extract_max/             # Extract maximum element
│   │       ├── insert/                  # Insert element
│   │       └── maximum/                 # Find maximum element
│   ├── lists/                           # List implementations (future)
│   ├── trees/                           # Tree implementations (future)
│   └── hash_tables/                     # Hash table implementations (future)
├── graphs/                              # Graph algorithms (future)
└── disjoint_sets/                       # Disjoint set operations (future)
```

## Features

- **Multiple Language Support**: Implementations in Python, C++, and Go
- **Interactive Documentation**: Jupyter notebooks with detailed explanations
- **Visual Examples**: PNG diagrams showing algorithm execution
- **Comprehensive Coverage**: From basic sorting to advanced data structures

## Getting Started

### Prerequisites

- Python 3.7+
- GCC/G++ compiler for C++ code
- Go 1.16+ for Go implementations
- Jupyter Notebook for interactive documentation

### Installation

1. Clone the repository:
```bash
git clone https://github.com/pomodoroCeleste/Algorythms.git
cd Algorythms
```

2. Install Python dependencies (if any):
```bash
pip install jupyter matplotlib numpy
```

## Algorithm Categories

### Comparison-Based Sorting

#### Divide and Conquer
- **Merge Sort**: Stable O(n log n) sorting algorithm with guaranteed performance
- **Quick Sort**: Average O(n log n) in-place sorting algorithm

#### Simple Comparison-Based
- **Bubble Sort**: Simple O(n²) algorithm for educational purposes
- **Heap Sort**: O(n log n) algorithm using heap data structure
- **Insertion Sort**: Efficient for small datasets, O(n²) worst case
- **Selection Sort**: O(n²) algorithm with minimal memory usage

### Linear Time Sorting

- **Bucket Sort**: O(n) average case for uniformly distributed data
- **Counting Sort**: O(n+k) for integers in a specific range
- **Radix Sort**: O(d·n) for integers with d digits

### Other Algorithms

- **Factorial**: Mathematical function implementation with different approaches

## Data Structures

### Stacks
- **Push**: Add element to top of stack
- **Pop**: Remove and return top element
- **Empty**: Check if stack has no elements
- **Operations**: General stack manipulation functions

### Queues
- **Basic Operations**: 
  - **Enqueue (accoda)**: Add element to rear
  - **Dequeue (extract)**: Remove element from front
  - **Initialize**: Set up queue structure
- **Priority Queues**:
  - **Insert**: Add element with priority
  - **Extract Max**: Remove highest priority element
  - **Maximum**: Find highest priority element

## File Organization

Each algorithm/data structure directory contains:
- **Implementation files**: `.py`, `.cpp`, `.go` files with source code
- **Documentation**: `readme.ipynb` with detailed explanations
- **Visual aids**: `.png` files showing algorithm execution
- **Examples**: Sample inputs and outputs

## Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature-name`
3. Add your implementation following the existing structure
4. Include comprehensive documentation and examples
5. Add visual diagrams if helpful
6. Submit a pull request

### Contribution Guidelines

- Follow existing naming conventions
- Include implementations in at least two languages
- Add Jupyter notebook documentation
- Include time and space complexity analysis
- Add test cases and examples

## Planned Features

- [ ] Binary Search Trees (BST)
- [ ] Red-Black Trees
- [ ] B-Trees
- [ ] Hash Tables with collision resolution
- [ ] Graph algorithms (BFS, DFS, Dijkstra, etc.)
- [ ] Disjoint Set (Union-Find) operations
- [ ] Dynamic programming algorithms
- [ ] String algorithms
- [ ] Geometric algorithms

## License

This project is open source and available under the [MIT License](LICENSE).

## Acknowledgments

- Algorithms are based on classic computer science textbooks
- Visual diagrams help illustrate complex concepts
- Multiple language implementations provide broader accessibility

---

For questions or suggestions, please open an issue or submit a pull request.
# Algorithms and Data Structures

## Folder Descriptions

### Comparison Based Sorting
Contains comparison-based sorting algorithms.
- **divide_and_conquer**: Algorithms that use the divide and conquer technique.
  - **mergesort**: Implementation and documentation of Merge Sort.
  - **quicksort**: Implementation and documentation of Quick Sort.
- **simple_comparison_based**: Simple comparison-based sorting algorithms.
  - **insertionsort**: Implementation and documentation of Insertion Sort.
  - **selectionsort**: Implementation and documentation of Selection Sort.
  - **bubblesort**: Implementation and documentation of Bubble Sort.
  - **heapsort**: Implementation and documentation of Heap Sort.

### Linear Time Sorting
Contains sorting algorithms that operate in linear time.
- **countingsort**: Implementation and documentation of Counting Sort.
- **radixsort**: Implementation and documentation of Radix Sort.
- **bucketsort**: Implementation and documentation of Bucket Sort.

### Other Algorithms
Contains other algorithms not classified in the previous categories.
- **factorial**: Implementation and documentation of the factorial calculation.

### Elementary Data Structures
Contains implementations of basic data structures.
- **Stack of Activation Records**: Implementations of stack operations.
  - **stack_empty**: Checks if the stack is empty.
  - **push**: Inserts an element into the stack.
  - **pop**: Removes an element from the stack.
- **Queues**: Implementations of queue operations.
  - **maximum**: Finds the maximum element in the priority queue.
  - **extractMax**: Extracts the maximum element from the priority queue.
  - **insert**: Inserts an element into the priority queue.
  - **enqueue**: Inserts an element into the queue.
  - **dequeue**: Removes an element from the queue.
  - **init_queue**: Initializes the queue with checks.
- **Lists**: Implementations of list operations.
  - **singly linked lists**: Operations on singly linked lists.
    - **lista_inserisci_testa**: Inserts an element at the head of the list.
    - **lista_estrai_testa**: Extracts an element from the head of the list.
    - **lista_vuota**: Checks if the list is empty.
    - **lista_inserisci_coda**: Inserts an element at the tail of the list.
    - **lista_estrai_coda**: Extracts an element from the tail of the list.
  - **linked lists**: Operations on linked lists.
    - **lista_cerca**: Searches for an element in the list.
    - **lista_cerca_ric**: Recursively searches for an element in the list.
    - **lista_inserisci_dopo**: Inserts an element after a specific node in the list.
  - **doubly linked lists**: Operations on doubly linked lists.
    - **listaD_cancella**: Deletes an element from the doubly linked list.

### Tree Traversals
Contains tree traversal algorithms.
- **Depth-First Traversal**:
  - **Pre Order**
  - **In Order**
  - **Post Order**
- **Breadth-First Traversal**:
  - **Level Order**

### Binary Search Trees
Contains operations on binary search trees.
- **search**:
  - **BST_search**: Searches for an element in the BST.
  - **BST_minimum**: Finds the minimum element in the BST.
  - **BST_maximum**: Finds the maximum element in the BST.
  - **BST_successor**: Finds the successor of an element in the BST.
- **node insertion**:
  - **BST_insert**: Inserts a node into the BST.
- **node deletion**:
  - **BST_delete**: Deletes a node from the BST.

### B-Trees
Contains implementations of B-Tree operations.
- **elementary operations**: Basic operations on B-Trees.
  - **BTree**: General B-Tree operations.
  - **Search**: Searches for an element in the B-Tree.
  - **Insert**: Inserts an element into the B-Tree.
  - **Delete**: Deletes an element from the B-Tree.
- **auxiliary procedures**: Helper procedures for B-Tree operations.
  - **SearchSubtree**: Searches within a subtree.
  - **SplitChild**: Splits a child node.
  - **InsertNonfull**: Inserts an element into a non-full node.

### Disjoint Set Data Structures
Contains implementations of disjoint set operations.
- **MakeSet**: Creates a new set.
- **FindSet**: Finds the representative of a set.
- **Union (Link)**: Unites two sets.

### Graphs
Contains graph algorithms and operations.
- **Graph Algorithms**: Basic graph traversal and sorting algorithms.
  - **Breadth-First-Search**: Implementation of breadth-first search.
  - **Depth-First-Search**: Implementation of depth-first search.
  - **Topological Sorting**: Implementation of topological sorting.
- **Greedy Algorithms (Minimum Spanning Tree)**: Algorithms for finding minimum spanning trees.
  - **Kruskal's Algorithm**: Implementation of Kruskal's algorithm.
  - **Prim's Algorithm**: Implementation of Prim's algorithm.
- **Shortest Path Algorithms**: Algorithms for finding the shortest path in a graph.
  - **Dijkstra's Algorithm**: Implementation of Dijkstra's algorithm.