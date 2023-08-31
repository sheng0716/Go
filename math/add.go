package main

import "fmt"

// Function to calculate the sum of two integers
func add(a, b int) int {
    return a + b
}

func main() {
    // Call the add function and print the result
    result := add(5, 3)
    fmt.Println("The sum is:", result)
}
