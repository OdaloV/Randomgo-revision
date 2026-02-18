package main

import (
    "errors"
)

var (
    ErrInvalidChar      = errors.New("invalid character in piece")
    ErrInvalidLength    = errors.New("line length not 4")
    ErrWrongHashCount   = errors.New("piece does not have exactly 4 hashes")
    ErrNotConnected     = errors.New("blocks are not connected")
    ErrEmptyFile        = errors.New("file is empty")
    ErrInvalidSeparator = errors.New("invalid piece separator")
)

// ValidateFile checks the entire raw file content
func ValidateFile(lines []string) error {
    // 1. Check file not empty
    if len(lines) == 0 {
        return ErrEmptyFile
    }
    
    // 2. Split into pieces (every 4 lines, separated by empty lines)
    pieces := splitIntoPieces(lines)
    
    // 3. Validate each piece
    for _, piece := range pieces {
        if err := ValidatePiece(piece); err != nil {
            return err
        }
    }
    
    return nil
}

// ValidatePiece validates a single 4-line tetromino
func ValidatePiece(piece []string) error {
    // 1. Check we have exactly 4 lines
    if len(piece) != 4 {
        return ErrInvalidLength
    }
    
    // 2. Validate each line
    hashCount := 0
    for _, line := range piece {
        // Check line length
        if len(line) != 4 {
            return ErrInvalidLength
        }
        
        // Check characters and count hashes
        for _, ch := range line {
            if ch != '#' && ch != '.' {
                return ErrInvalidChar
            }
            if ch == '#' {
                hashCount++
            }
        }
    }
    
    // 3. Must have exactly 4 hashes
    if hashCount != 4 {
        return ErrWrongHashCount
    }
    
    // 4. Check connectivity
    if !areBlocksConnected(piece) {
        return ErrNotConnected
    }
    
    return nil
}

// Helper: split raw lines into pieces
func splitIntoPieces(lines []string) [][]string {
    var pieces [][]string
    var currentPiece []string
    
    for _, line := range lines {
        if line == "" {
            // Empty line separates pieces
            if len(currentPiece) > 0 {
                if len(currentPiece) == 4 {
                    pieces = append(pieces, currentPiece)
                }
                currentPiece = []string{}
            }
            continue
        }
        
        currentPiece = append(currentPiece, line)
        
        // If we have 4 lines, that's a complete piece
        if len(currentPiece) == 4 {
            pieces = append(pieces, currentPiece)
            currentPiece = []string{}
        }
    }
    
    // Handle last piece (if file doesn't end with empty line)
    if len(currentPiece) == 4 {
        pieces = append(pieces, currentPiece)
    }
    
    return pieces
}

// Helper: check if 4 blocks are connected using BFS
func areBlocksConnected(piece []string) bool {
    // 1. Find first '#' as starting point
    startX, startY := -1, -1
    found := false
    
    for y := 0; y < 4; y++ {
        for x := 0; x < 4; x++ {
            if piece[y][x] == '#' {
                startX, startY = x, y
                found = true
                break
            }
        }
        if found {
            break
        }
    }
    
    if !found {
        return false // No blocks found (shouldn't happen due to earlier validation)
    }
    
    // 2. BFS to find all connected blocks
    visited := make(map[Block]bool)
    queue := []Block{{startX, startY}}
    visited[Block{startX, startY}] = true
    connectedCount := 1
    
    // Directions: right, left, down, up
    directions := []Block{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        
        for _, dir := range directions {
            next := Block{current.X + dir.X, current.Y + dir.Y}
            
            // Check bounds
            if next.X >= 0 && next.X < 4 && next.Y >= 0 && next.Y < 4 {
                // Check if there's a block and not visited
                if piece[next.Y][next.X] == '#' && !visited[next] {
                    visited[next] = true
                    queue = append(queue, next)
                    connectedCount++
                }
            }
        }
    }
    
    return connectedCount == 4
}