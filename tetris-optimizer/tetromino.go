package main 

import "fmt"

type Block struct {
    X, Y int  // Using X,Y for grid coordinates
}

type Tetromino struct {
    Points []Block  // The 4 block positions (normalized to 0,0)
    Letter byte     // 'A', 'B', 'C', etc.
}

// New creates a tetromino from 4x4 grid
func New(shape [][]rune, letter byte) (*Tetromino, error) {
    // 1. Extract all '#' positions
    var points []Block
    for y := 0; y < 4; y++ {
        for x := 0; x < 4; x++ {
            if shape[y][x] == '#' {
                points = append(points, Block{x, y})
            }
        }
    }
    
    // 2. Must have exactly 4 blocks
    if len(points) != 4 {
        return nil, fmt.Errorf("invalid tetromino: must have 4 blocks")
    }
    
    // 3. Check connectivity (BFS)
    if !isConnected(points) {
        return nil, fmt.Errorf("invalid tetromino: blocks not connected")
    }
    
    // 4. Normalize to origin
    points = normalize(points)
    
    return &Tetromino{
        Points: points,
        Letter: letter,
    }, nil
}

// normalize shifts all points so min X=0 and min Y=0
func normalize(points []Block) []Block {
    minX, minY := points[0].X, points[0].Y
    for _, p := range points {
        if p.X < minX {
            minX = p.X
        }
        if p.Y < minY {
            minY = p.Y
        }
    }
    
    normalized := make([]Block, len(points))
    for i, p := range points {
        normalized[i] = Block{p.X - minX, p.Y - minY}
    }
    return normalized
}

// isConnected checks if all 4 blocks are connected orthogonally
func isConnected(points []Block) bool {
    // Build adjacency map
    pointMap := make(map[Block]bool)
    for _, p := range points {
        pointMap[p] = true
    }
    
    // BFS from first point
    visited := make(map[Block]bool)
    queue := []Block{points[0]}
    visited[points[0]] = true
    count := 1
    
    directions := []Block{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        
        for _, d := range directions {
            next := Block{current.X + d.X, current.Y + d.Y}
            if pointMap[next] && !visited[next] {
                visited[next] = true
                queue = append(queue, next)
                count++
            }
        }
    }
    
    return count == 4
}

// Rotate returns new tetromino rotated 90° clockwise
func (t *Tetromino) Rotate() *Tetromino {
    rotated := make([]Block, 4)
    for i, p := range t.Points {
        // For 90° clockwise: (x,y) -> (y, 3-x)
        rotated[i] = Block{p.Y, 3 - p.X}
    }
    rotated = normalize(rotated)
    
    return &Tetromino{
        Points: rotated,
        Letter: t.Letter,
    }
}

// GetAllRotations returns all unique rotations
func (t *Tetromino) GetAllRotations() []*Tetromino {
    rotations := []*Tetromino{t}
    current := t
    
    for i := 0; i < 3; i++ {
        current = current.Rotate()
        
        // Check if unique
        unique := true
        for _, existing := range rotations {
            if existing.Equals(current) {
                unique = false
                break
            }
        }
        
        if unique {
            rotations = append(rotations, current)
        }
    }
    
    return rotations
}

// Equals checks if two tetrominos have same points
func (t *Tetromino) Equals(other *Tetromino) bool {
    if len(t.Points) != len(other.Points) {
        return false
    }
    
    // Create maps for comparison
    tMap := make(map[Block]bool)
    for _, p := range t.Points {
        tMap[p] = true
    }
    
    for _, p := range other.Points {
        if !tMap[p] {
            return false
        }
    }
    
    return true
}