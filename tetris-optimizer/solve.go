package main

// Grid represents the square where tetrominoes are placed
type Grid struct {
    cells [][]byte 
    size  int      
}

// Solve finds the smallest square arrangement for all tetrominoes
func Solve(pieces []*Tetromino) [][]byte {
    if len(pieces) == 0 {
        return nil
    }
    
    // Calculate minimum possible size
    // Total blocks = number of pieces × 4
    // Square size must be at least ceil(sqrt(total blocks))
    totalBlocks := len(pieces) * 4
    minSize := 1
    for minSize*minSize < totalBlocks {
        minSize++
    }
    
    // Try increasing square sizes until solution found
    // We try up to minSize+3 to allow for some flexibility
    for size := minSize; size <= minSize+3; size++ {
        grid := NewGrid(size)
        if backtrack(grid, pieces, 0) {
            return grid.cells
        }
    }
    
    return nil // No solution found
}

// NewGrid creates an empty grid of given size
func NewGrid(size int) *Grid {
    cells := make([][]byte, size)
    for i := range cells {
        cells[i] = make([]byte, size)
        for j := range cells[i] {
            cells[i][j] = '.' // Empty cell
        }
    }
    return &Grid{
        cells: cells,
        size:  size,
    }
}

// backtrack tries to place pieces recursively using DFS
func backtrack(grid *Grid, pieces []*Tetromino, index int) bool {
    // Base case: all pieces placed successfully
    if index == len(pieces) {
        return true
    }
    
    piece := pieces[index]
    
    // Try all possible rotations of current piece
    for _, rotation := range piece.GetAllRotations() {
        // Try all possible positions on the grid
        for y := 0; y < grid.size; y++ {
            for x := 0; x < grid.size; x++ {
                // Check if piece can be placed at (x, y)
                if canPlace(grid, rotation, x, y) {
                    // Place the piece
                    place(grid, rotation, x, y, piece.Letter)
                    
                    // Recursively try to place remaining pieces
                    if backtrack(grid, pieces, index+1) {
                        return true // Solution found!
                    }
                    
                    // If we get here, this placement didn't work
                    // Backtrack: remove the piece
                    remove(grid, rotation, x, y)
                }
            }
        }
    }
    
    return false // No placement worked for this piece
}

// canPlace checks if a piece fits at position (x, y) on the grid
func canPlace(grid *Grid, piece *Tetromino, x, y int) bool {
    for _, block := range piece.Points {
        newX, newY := x+block.X, y+block.Y
        
        // Check if block is within grid boundaries
        if newX < 0 || newX >= grid.size || newY < 0 || newY >= grid.size {
            return false
        }
        
        // Check if cell is already occupied
        if grid.cells[newY][newX] != '.' {
            return false
        }
    }
    return true
}

// place puts a piece on the grid at position (x, y)
func place(grid *Grid, piece *Tetromino, x, y int, letter byte) {
    for _, block := range piece.Points {
        grid.cells[y+block.Y][x+block.X] = letter
    }
}

// remove takes a piece off the grid from position (x, y)
func remove(grid *Grid, piece *Tetromino, x, y int) {
    for _, block := range piece.Points {
        grid.cells[y+block.Y][x+block.X] = '.'
    }
}