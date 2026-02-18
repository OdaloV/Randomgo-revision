package main

import (
    "os"
    "testing"
)

// TestInvalidFiles tests only the most critical error cases
func TestInvalidFiles(t *testing.T) {
    tests := []struct {
        name    string
        content string
    }{
        {
            name: "empty file",
            content: "",
        },
        {
            name: "invalid character",
            content: "...#\n...#\n..x#\n...#\n",
        },
        {
            name: "missing separator",
            content: "...#\n...#\n...#\n...#\n....\n....\n....\n####\n",
        },
        {
            name: "extra empty line",
            content: "...#\n...#\n...#\n...#\n\n\n....\n....\n....\n####\n",
        },
        {
            name: "wrong line length",
            content: "...#\n...#\n...#\n..##..\n",
        },
        {
            name: "wrong hash count (3)",
            content: "...#\n...#\n...#\n....\n",
        },
        {
            name: "wrong hash count (5)",
            content: ".##.\n####\n.##.\n....\n",
        },
        {
            name: "disconnected blocks",
            content: "#...\n.#..\n..#.\n...#\n",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Create temp file
            tmpfile, err := os.CreateTemp("", "test*.txt")
            if err != nil {
                t.Fatal(err)
            }
            defer os.Remove(tmpfile.Name())

            if _, err := tmpfile.WriteString(tt.content); err != nil {
                t.Fatal(err)
            }
            tmpfile.Close()

            // Parse should return error
            _, err = ParseFile(tmpfile.Name())
            if err == nil {
                t.Errorf("ParseFile() with %s: expected error, got nil", tt.name)
            }
        })
    }
}

// TestValidShapes tests that all 7 tetromino shapes are accepted
func TestValidShapes(t *testing.T) {
    shapes := []struct {
        name    string
        content string
    }{
        {
            name: "I shape",
            content: "...#\n...#\n...#\n...#\n",
        },
        {
            name: "O shape",
            content: "....\n.##.\n.##.\n....\n",
        },
        {
            name: "T shape",
            content: "....\n###.\n.#..\n....\n",
        },
        {
            name: "L shape",
            content: "#...\n#...\n##..\n....\n",
        },
        {
            name: "J shape",
            content: "..#.\n..#.\n.##.\n....\n",
        },
        {
            name: "S shape",
            content: ".##.\n##..\n....\n....\n",
        },
        {
            name: "Z shape",
            content: "##..\n.##.\n....\n....\n",
        },
    }

    for _, tt := range shapes {
        t.Run(tt.name, func(t *testing.T) {
            tmpfile, err := os.CreateTemp("", "test*.txt")
            if err != nil {
                t.Fatal(err)
            }
            defer os.Remove(tmpfile.Name())

            if _, err := tmpfile.WriteString(tt.content); err != nil {
                t.Fatal(err)
            }
            tmpfile.Close()

            tetros, err := ParseFile(tmpfile.Name())
            if err != nil {
                t.Errorf("ParseFile() with %s: unexpected error %v", tt.name, err)
            }
            if len(tetros) != 1 {
                t.Errorf("ParseFile() with %s: got %d tetrominos, want 1", tt.name, len(tetros))
            }
        })
    }
}

// TestSolveEdgeCases tests solver with minimal cases
func TestSolveEdgeCases(t *testing.T) {
    // Test with no pieces
    solution := Solve([]*Tetromino{})
    if solution != nil {
        t.Error("Solve() with no pieces: expected nil, got grid")
    }

    // Test with single O piece (smallest square)
    oShape := [][]rune{
        {'.', '.', '.', '.'},
        {'.', '#', '#', '.'},
        {'.', '#', '#', '.'},
        {'.', '.', '.', '.'},
    }
    oPiece, _ := New(oShape, 'A')
    
    solution = Solve([]*Tetromino{oPiece})
    if solution == nil {
        t.Error("Solve() with O piece: expected solution, got nil")
    }
    if len(solution) != 2 {
        t.Errorf("Solve() with O piece: grid size %d, want 2", len(solution))
    }
}

// TestFileNotFound tests missing file
func TestFileNotFound(t *testing.T) {
    _, err := ParseFile("does_not_exist.txt")
    if err == nil {
        t.Error("ParseFile() with missing file: expected error, got nil")
    }
}