# Data Structures

This directory contains implementations of fundamental data structures that are essential building blocks for many algorithms and applications.

## Implemented Structures

### [Stacks](stacks/)
- **LIFO (Last In, First Out)** data structure
- **Operations**: Push, Pop, Empty check
- **Applications**: Function calls, expression evaluation, backtracking

### [Queues](queues/)
- **FIFO (First In, First Out)** data structure
- **Basic Operations**: Enqueue, Dequeue, Initialize
- **Priority Queues**: Insert, Extract Max, Find Maximum
- **Applications**: Task scheduling, breadth-first search, simulation

## Planned Structures

### Lists (Coming Soon)
- Singly linked lists
- Doubly linked lists  
- Lists with sentinels
- Array-based lists

### Trees (Coming Soon)
- Binary trees
- Binary search trees
- Balanced trees (AVL, Red-Black)
- B-trees

### Hash Tables (Coming Soon)
- Direct addressing
- Chaining for collision resolution
- Open addressing methods
- Perfect hashing

## Data Structure Comparison

| Structure | Access | Insert | Delete | Search | Space |
|-----------|--------|--------|--------|--------|-------|
| Stack | O(1) top | O(1) | O(1) | O(n) | O(n) |
| Queue | O(1) front/rear | O(1) | O(1) | O(n) | O(n) |
| Priority Queue | O(1) max | O(log n) | O(log n) | O(n) | O(n) |

## File Organization

Each data structure directory contains:
- **Implementation files**: Source code in Python, C++, and Go
- **Operation subdirectories**: Individual operations with examples
- **Documentation**: Jupyter notebooks with explanations
- **Visual aids**: Diagrams showing structure and operations

## Usage Guidelines

1. **Choose the right structure**: Consider access patterns and performance requirements
2. **Understand operations**: Each structure has specific strengths and limitations
3. **Study implementations**: Compare different language approaches
4. **Practice with examples**: Use provided test cases and documentation

## Applications

### Stacks
- Function call management
- Expression parsing and evaluation
- Undo operations in applications
- Backtracking algorithms
- Browser history

### Queues
- Process scheduling in operating systems
- Breadth-first search in graphs
- Print job management
- Handling requests in web servers
- Simulation of real-world queuing systems

### Priority Queues
- Task scheduling with priorities
- Dijkstra's shortest path algorithm
- Huffman coding trees
- A* search algorithm
- Event simulation
