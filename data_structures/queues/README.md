# Queue Data Structure

A queue is a linear data structure that follows the **FIFO (First In, First Out)** principle. Elements are added at one end (rear) and removed from the other end (front).

## Basic Queue Operations

### [Operations](operations/)
Contains basic queue operations:

#### [Enqueue (Accoda)](operations/accoda/)
- **Description**: Insert an element at the rear of the queue
- **Time Complexity**: O(1)
- **Space Complexity**: O(1)

#### [Dequeue (Extract)](operations/extract/)
- **Description**: Remove and return the front element from the queue
- **Time Complexity**: O(1)
- **Space Complexity**: O(1)

#### [Initialize](operations/init/)
- **Description**: Initialize an empty queue
- **Time Complexity**: O(1)
- **Space Complexity**: O(1)

## Priority Queue Operations

### [Priority Queues](priority/)
A priority queue maintains elements with associated priorities:

#### [Insert](priority/insert/)
- **Description**: Insert an element with a given priority
- **Time Complexity**: O(log n) for heap-based implementation
- **Space Complexity**: O(1)

#### [Extract Maximum](priority/extract_max/)
- **Description**: Remove and return the element with highest priority
- **Time Complexity**: O(log n) for heap-based implementation
- **Space Complexity**: O(1)

#### [Find Maximum](priority/maximum/)
- **Description**: Return the element with highest priority without removing it
- **Time Complexity**: O(1) for heap-based implementation
- **Space Complexity**: O(1)

## Queue Properties

- **FIFO Access**: First element added is first to be removed
- **Two Access Points**: Add at rear, remove from front
- **Dynamic Size**: Can grow and shrink during runtime
- **Sequential Processing**: Natural for handling ordered tasks

## Applications

### Basic Queues
1. **Process Scheduling**: Operating system task management
2. **Print Job Management**: Handle print requests in order
3. **Breadth-First Search**: Graph traversal algorithm
4. **Buffer for Data Streams**: Handle incoming data
5. **Web Server Request Handling**: Process requests in order

### Priority Queues
1. **Task Scheduling**: Handle tasks based on priority levels
2. **Dijkstra's Algorithm**: Find shortest paths in graphs
3. **Huffman Coding**: Build optimal compression trees
4. **A* Search Algorithm**: Pathfinding with heuristics
5. **Event Simulation**: Process events by timestamp

## Implementation Types

### Array-Based Queue
- **Pros**: Simple implementation, cache-friendly
- **Cons**: Fixed size, potential for false fullness
- **Solution**: Circular array to maximize space usage

### Linked List-Based Queue
- **Pros**: Dynamic size, no wasted space
- **Cons**: Extra memory for pointers, potential cache misses

### Priority Queue (Heap)
- **Pros**: Efficient priority operations
- **Cons**: More complex than basic queue
- **Implementation**: Binary heap for optimal performance

## Common Queue Operations

```
ENQUEUE(Q, x)
  Q[Q.tail] = x
  if Q.tail == Q.length
    Q.tail = 1
  else
    Q.tail = Q.tail + 1

DEQUEUE(Q)
  x = Q[Q.head]
  if Q.head == Q.length
    Q.head = 1
  else
    Q.head = Q.head + 1
  return x
```

## Error Conditions

- **Queue Overflow**: Attempting to enqueue into a full queue
- **Queue Underflow**: Attempting to dequeue from an empty queue

## File Structure

```
queues/
├── README.md                 # This file
├── operations/               # Basic queue operations
│   ├── accoda/              # Enqueue operation
│   ├── extract/             # Dequeue operation
│   └── init/                # Queue initialization
└── priority/                # Priority queue operations
    ├── insert/              # Insert with priority
    ├── extract_max/         # Extract highest priority
    └── maximum/             # Find highest priority
```

Each operation directory contains implementation files, documentation, and visual examples.
