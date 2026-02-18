package main

import (
    "bufio"
	"fmt"
    "os"
    "strings"
	
)

// ParseFile reads and parses the tetromino file
func ParseFile(filename string) ([]*Tetromino, error) {
    // Read all lines
    lines, err := readLines(filename)
    if err != nil {
        return nil, err
    }
    
    // Clean lines (remove special characters from the file)
    // Makes lines consistent for processing
    cleanLines := cleanLines(lines)

	if err := ValidateFile(cleanLines); err != nil {
		return nil, err
    }
    
    // Split tetromino into pieces
    pieces, err := splitPieces(cleanLines)
if err != nil {
    return nil, err
}
    
    var tetrominos []*Tetromino // Hold converted tetrominos
    for i, piece := range pieces { // Loop through the pieces
        shape := make([][]rune, 4) // Converts them into a 2D character array
        for r, line := range piece {
            shape[r] = []rune(line)
        }
        
        // Creates the tetromino object assigning letters A, B, C...
        t, err := New(shape, byte('A'+i))
        if err != nil { // If wrong shape, blocks not connected
            return nil, err
        }
        
        tetrominos = append(tetrominos, t) // Add valid tetromino to slice
    }
    
    return tetrominos, nil // Return all tetrominos
}

// readLines reads all lines from file
func readLines(filename string) ([]string, error) {
    file, err := os.Open(filename) // Opens file, if it can't open, throw error
    if err != nil {
        return nil, err
    }
    defer file.Close() // File closes when function exits
    
    var lines []string
    scanner := bufio.NewScanner(file) // Creates scanner to read file line by line
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    
    return lines, scanner.Err()
}

// cleanLines removes any special characters
func cleanLines(lines []string) []string {
    cleaned := make([]string, 0, len(lines))
    for _, line := range lines {
        // Remove \r if present (Windows line endings)
        line = strings.TrimRight(line, "\r")
        
        // Remove $ if present (from cat -e output)
        line = strings.TrimSuffix(line, "$")
        
        cleaned = append(cleaned, line)
    }
    return cleaned
}

// splitPieces groups lines into tetromino pieces
func splitPieces(lines []string) ([][]string, error) {
    var pieces [][]string
    var currentPiece []string
    expectEmptyLine := false
    
    for i, line := range lines {
        // After completing a piece, we MUST have an empty line
        if expectEmptyLine {
            if line != "" {
                return nil, fmt.Errorf("missing empty line between pieces at line %d", i+1)
            }
            expectEmptyLine = false
            continue
        }
        
        // Empty line when not expected = error
        if line == "" {
            if len(currentPiece) == 0 {
                return nil, fmt.Errorf("extra empty line at beginning of file")
            }
            if len(currentPiece) != 4 {
                return nil, fmt.Errorf("incomplete piece (%d lines) followed by empty line", len(currentPiece))
            }
            expectEmptyLine = true
            continue
        }
        
        // Normal line - add to current piece
        if len(line) != 4 {
            return nil, fmt.Errorf("invalid line length %d at line %d", len(line), i+1)
        }
        
        currentPiece = append(currentPiece, line)
        
        // If we have 4 lines, that's a complete piece
        if len(currentPiece) == 4 {
            pieces = append(pieces, currentPiece)
            currentPiece = []string{}
            expectEmptyLine = true
        }
        
        // If we somehow get more than 4 lines, error
        if len(currentPiece) > 4 {
            return nil, fmt.Errorf("piece has more than 4 lines at line %d", i+1)
        }
    }
    
    // Check for incomplete piece at end
    if len(currentPiece) > 0 {
        if len(currentPiece) == 4 {
            pieces = append(pieces, currentPiece)
        } else {
            return nil, fmt.Errorf("incomplete piece at end of file (%d lines)", len(currentPiece))
        }
    }
    
    return pieces, nil
}