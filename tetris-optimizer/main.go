package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("ERROR,one arguement only")
    }
    
    filename := os.Args[1]
    
    // Parse file and get tetrominoes - directly call functions!
    tetrominoes, err := ParseFile(filename)  
    if err != nil {
        fmt.Println("ERROR")
    }
    
    // Solve the puzzle - directly call functions!
    solution := Solve(tetrominoes)  
    if solution == nil {
        fmt.Println("ERROR")
    }
    
    // Print solution
    for _, row := range solution {
        fmt.Println(string(row))  // Convert []byte to string for printing
    }
}