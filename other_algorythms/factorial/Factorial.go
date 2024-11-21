package main

import "fmt"

// factorial calculates the factorial of a non-negative integer n using recursion
func factorial(n int) int {
    // Base case: if n is 0, return 1
    if n == 0 {
        return 1
    }
    // Recursive case: n * factorial of (n-1)
    return n * factorial(n-1)
}

func main() {
    // Example value
    n := 5
    // Calculate the factorial of n
    result := factorial(n)
    // Print the result
    fmt.Printf("Factorial of %d is %d\n", n, result)
}