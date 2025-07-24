# Stack Data Structure

A stack is a linear data structure that follows the **LIFO (Last In, First Out)** principle. Elements are added and removed from the same end, called the "top" of the stack.

## Operations

### [Push](push/)
- **Description**: Insert an element onto the top of the stack
- **Time Complexity**: O(1)
- **Space Complexity**: O(1)

### [Pop](pop/)
- **Description**: Remove and return the top element from the stack
- **Time Complexity**: O(1)
- **Space Complexity**: O(1)

### [Empty Check](empty/)
- **Description**: Check if the stack contains no elements
- **Time Complexity**: O(1)
- **Space Complexity**: O(1)

### [General Operations](operations/)
- **Description**: Combined stack operations and examples
- **Includes**: Push, Pop, Empty check implementations

## Stack Properties

- **LIFO Access**: Last element added is first to be removed
- **Single Access Point**: Elements can only be added/removed from the top
- **Dynamic Size**: Can grow and shrink during runtime (implementation dependent)
- **No Random Access**: Cannot access elements in the middle without popping

## Applications

1. **Function Call Management**: Store return addresses and local variables
2. **Expression Evaluation**: Parse and evaluate mathematical expressions
3. **Undo Operations**: Implement undo functionality in applications
4. **Backtracking Algorithms**: Store states for algorithms like maze solving
5. **Browser History**: Navigate back through visited pages

## Implementation Considerations

### Array-Based Stack
- **Pros**: Simple implementation, cache-friendly
- **Cons**: Fixed maximum size (unless dynamic array)
- **Space**: O(n) where n is maximum capacity

### Linked List-Based Stack
- **Pros**: Dynamic size, memory efficient
- **Cons**: Extra memory for pointers
- **Space**: O(n) where n is current number of elements

## Common Stack Operations

```
STACK-EMPTY(S)
  return S.top == 0

PUSH(S, x)
  S.top = S.top + 1
  S[S.top] = x

POP(S)
  if STACK-EMPTY(S)
    error "underflow"
  else
    S.top = S.top - 1
    return S[S.top + 1]
```

## Error Conditions

- **Stack Overflow**: Attempting to push onto a full stack
- **Stack Underflow**: Attempting to pop from an empty stack

## File Structure

```
stacks/
├── README.md                 # This file
├── push/                     # Push operation
├── pop/                      # Pop operation  
├── empty/                    # Empty check operation
└── operations/               # Combined operations and examples
```

Each operation directory contains implementation files, documentation, and visual examples.
