# Algorythms and data structures


- **ALGORITHMS EXAMINED**:

algorithms/
├── comparison_based_sorting/
│   ├── divide_and_conquer/
│   │   ├── mergesort/
│   │   └── quicksort/
│   ├── simple_comparison_based/
│   │   ├── insertionsort/
│   │   ├── selectionsort/
│   │   ├── bubblesort/
│   │   └── heapsort/
├── linear_time_sorting/
│   ├── countingsort/
│   ├── radixsort/
│   └── bucketsort/
└── other_algorithms/
    └── factorial/

- **ELEMENTARY DATA STRUCTURES**:

data_structures/
├── stack_of_activation_records/
│   ├── stack_empty/
│   ├── push/
│   └── pop/
├── queues/
│   ├── priority_queues/
│   │   ├── maximum/
│   │   ├── extractMax/
│   │   └── insert/
│   ├── queue_operations/
│   │   ├── accoda/
│   │   ├── estrai_da_coda/
│   │   └── init_coda/
├── lists/
│   ├── singly_linked_lists/
│   │   ├── lista_inserisci_testa/
│   │   ├── lista_estrai_testa/
│   │   ├── lista_vuota/
│   │   ├── lista_inserisci_coda/
│   │   └── lista_estrai_coda/
│   ├── linked_lists/
│   │   ├── lista_cerca/
│   │   ├── lista_cerca_ric/
│   │   └── lista_inserisci_dopo/
│   ├── doubly_linked_lists/
│   │   └── listaD_cancella/
│   ├── lists_with_sentinel/
│   │   ├── lista_cancellaS/
│   │   ├── lista_inserisciS_testa/
│   │   └── lista_cercaS/
│   └── lists_with_indices/
│       ├── alloca_elemento/
│       └── dealloca_elemento/

- **ORDINAL SELECTION ALGORITHMS**:

algorithms/
├── ordinal_selection/
│   ├── min_max/
│   └── rand_select/

- **DICTIONARIES AND HASH FUNCTIONS**:

dictionaries_and_hash_functions/
├── basic_operations/
│   ├── search/
│   ├── insert/
│   └── delete/
├── collision_resolution_with_chaining/
│   ├── search/
│   ├── insert/
│   └── delete/
├── collision_resolution_with_open_addressing/
│   ├── search/
│   ├── insert/
│   └── delete/
└── quadratic_probing/
    └── verification_algorithm/

- **TREES**:

trees/
├── depth_first_traversal/
│   ├── pre_order/
│   ├── in_order/
│   └── post_order/
├── breadth_first_traversal/
│   └── breadth/
├── binary_search_trees/
│   ├── search/
│   │   ├── BST_search/
│   │   ├── BST_minimum/
│   │   ├── BST_maximum/
│   │   └── BST_successor/
│   ├── node_insertion/
│   │   └── BST_insert/
│   └── node_deletion/
│       └── BST_delete/
├── rb_trees/
│   ├── rb_tree_rotation/
│   │   ├── rotate_right/
│   │   └── rotate_left/
│   ├── rb_tree_insertion/
│   │   └── RB_insert/
│   └── rb_tree_deletion/
│       ├── RB_delete/
│       └── RB_Fix_delete/
└── b_trees/
    ├── elementary_operations/
    │   ├── BTree/
    │   ├── Search/
    │   ├── Insert/
    │   └── Delete/
    └── auxiliary_procedures/
        ├── SearchSubtree/
        ├── SplitChild/
        └── InsertNonfull/

- **DISJOINT SET DATA STRUCTURES**:

disjoint_set_data_structures/
├── MakeSet/
├── FindSet/
└── Union/

- **GRAPHS**:

graphs/
├── graph_algorithms/
│   ├── breadth_first_search/
│   ├── depth_first_search/
│   └── topological_sorting/
├── greedy_algorithms/
│   ├── kruskal/
│   └── prim/
└── shortest_path_algorithms/
    └── dijkstra/
_________________________________________________________________________________________________________________________
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