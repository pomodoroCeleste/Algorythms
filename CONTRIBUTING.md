# Contributing Guide

Thank you for your interest in contributing to the Algorithms and Data Structures repository! This guide will help you understand our standards and make meaningful contributions.

## Repository Structure

Our repository follows a consistent structure for organization and maintainability:

```
├── algorithms/                    # Algorithm implementations
├── data_structures/              # Data structure implementations  
├── graphs/                       # Graph algorithms (planned)
├── disjoint_sets/                # Union-Find operations (planned)
└── README.md                     # Main documentation
```

## Contribution Types

### 1. New Algorithm Implementation
- Add algorithms not currently in the repository
- Implement missing language versions of existing algorithms
- Optimize existing implementations

### 2. New Data Structure
- Implement fundamental data structures
- Add operations for existing structures
- Create specialized variants

### 3. Documentation Improvements
- Enhance existing Jupyter notebooks
- Add better visual explanations
- Improve code comments and examples

### 4. Bug Fixes
- Fix incorrect implementations
- Resolve performance issues
- Address documentation errors

## Standards and Requirements

### Code Implementation Standards

#### Multi-Language Support
All algorithms must be implemented in **at least two** of these languages:
- **Python** (primary): Clean, readable code with type hints
- **C++** (performance): Optimized, standard library compliant
- **Go** (modern): Idiomatic Go with proper error handling

#### Code Quality Standards
```python
# Python example - clean and documented
def merge_sort(arr: List[int]) -> List[int]:
    """
    Sorts an array using the merge sort algorithm.
    
    Args:
        arr: List of integers to sort
        
    Returns:
        New sorted list
        
    Time Complexity: O(n log n)
    Space Complexity: O(n)
    """
    if len(arr) <= 1:
        return arr
    
    mid = len(arr) // 2
    left = merge_sort(arr[:mid])
    right = merge_sort(arr[mid:])
    
    return merge(left, right)
```

#### Required Elements for Each Implementation
1. **Clear function/method names**
2. **Comprehensive docstrings/comments**
3. **Type annotations** (Python) or proper typing (C++/Go)
4. **Error handling** for edge cases
5. **Example usage** in comments or separate files

### Documentation Standards

#### Jupyter Notebook Requirements
Each algorithm/data structure must include a `readme.ipynb` with:

1. **Algorithm/Structure Description**
   - Clear explanation of the concept
   - When and why to use it
   - Real-world applications

2. **Step-by-Step Walkthrough**
   - Detailed algorithm steps
   - Example execution with sample data
   - Visual diagrams where helpful

3. **Complexity Analysis**
   - Time complexity (best, average, worst case)
   - Space complexity
   - Comparison with alternatives

4. **Implementation Discussion**
   - Key implementation details
   - Common pitfalls and how to avoid them
   - Optimization opportunities

5. **Code Examples**
   - Working code snippets
   - Test cases with expected outputs
   - Performance comparisons

#### Visual Documentation
- **Include diagrams**: Create or find appropriate `.png` files
- **Show execution steps**: Visualize algorithm progress
- **Highlight key concepts**: Use diagrams to explain complex parts

### Directory Structure for New Contributions

#### Algorithm Contribution Structure
```
algorithms/category/AlgorithmName/
├── AlgorithmName.py              # Python implementation
├── AlgorithmName.cpp             # C++ implementation  
├── AlgorithmName.go              # Go implementation
├── readme.ipynb                  # Comprehensive documentation
├── example_diagram.png           # Visual explanation
├── test_cases.md                 # Test cases and expected outputs
└── README.md                     # Brief overview (optional)
```

#### Data Structure Contribution Structure
```
data_structures/StructureName/
├── operations/                   # Individual operations
│   ├── operation1/
│   │   ├── operation1.py
│   │   ├── operation1.cpp
│   │   ├── operation1.go
│   │   └── readme.ipynb
│   └── operation2/
├── StructureName.py              # Complete implementation
├── StructureName.cpp             # Complete implementation
├── StructureName.go              # Complete implementation
├── readme.ipynb                  # Structure documentation
└── README.md                     # Overview and operation list
```

## Submission Process

### 1. Fork and Clone
```bash
git clone https://github.com/your-username/Algorythms.git
cd Algorythms
```

### 2. Create Feature Branch
```bash
git checkout -b feature/algorithm-name
# or
git checkout -b feature/data-structure-name
```

### 3. Implement Your Contribution
- Follow the structure standards above
- Write comprehensive tests
- Create documentation
- Add visual aids if helpful

### 4. Test Your Implementation
- Verify correctness with multiple test cases
- Check performance characteristics
- Ensure code runs in all target languages
- Validate documentation renders properly

### 5. Submit Pull Request
- Clear title describing the contribution
- Detailed description of what you've added
- List of files changed
- Any special considerations or dependencies

## Pull Request Template

```markdown
## Description
Brief description of the algorithm/data structure added.

## Type of Contribution
- [ ] New algorithm implementation
- [ ] New data structure
- [ ] Documentation improvement
- [ ] Bug fix
- [ ] Performance optimization

## Implementation Details
- **Languages implemented**: Python, C++, Go
- **Time complexity**: O(?)
- **Space complexity**: O(?)
- **Key features**: Notable aspects of the implementation

## Testing
- [ ] Code compiles/runs in all languages
- [ ] Comprehensive test cases included
- [ ] Documentation renders correctly
- [ ] Visual aids are clear and helpful

## Files Changed
- List of new files added
- List of existing files modified

## Additional Notes
Any special considerations, dependencies, or implementation details.
```

## Code Review Process

### Review Criteria
1. **Correctness**: Algorithm implements the intended behavior
2. **Efficiency**: Implementation follows expected complexity bounds
3. **Clarity**: Code is readable and well-documented
4. **Completeness**: All required elements are present
5. **Consistency**: Follows repository conventions

### Review Timeline
- Initial review within 48 hours
- Follow-up reviews within 24 hours of updates
- Merge after approval from at least one maintainer

## Getting Help

### Resources
- **Algorithm references**: CLRS, Algorithm Design Manual
- **Documentation examples**: See existing implementations
- **Style guides**: Language-specific best practices

### Communication
- **Issues**: Open GitHub issues for questions
- **Discussions**: Use GitHub discussions for general topics
- **Email**: Contact maintainers for private concerns

## Recognition

Contributors will be:
- Listed in repository contributors
- Credited in relevant documentation
- Acknowledged in release notes for significant contributions

## Code of Conduct

We maintain a welcoming and inclusive environment:
- Be respectful in all interactions
- Provide constructive feedback
- Help others learn and improve
- Focus on technical merit over personal preferences

Thank you for helping make this repository a valuable resource for the computer science community!
