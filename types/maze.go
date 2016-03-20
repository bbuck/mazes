package types

// MazeRow represents a single row of the maze, which is a list of directions.
type MazeRow []Direction

// Maze represents a maze structured as a list of MazeRows which are themselves
// a list of valid directions.
type Maze []MazeRow

// Height returns the pixel height of the maze give the configurations height
// value for the maze
func (m Maze) Height(ch float64) int {
	return len(m) * int(ch)
}

// Width returns the pixel width of the maze give the configurations width
// value for the maze
func (m Maze) Width(cw float64) int {
	if len(m) > 0 {
		return len(m[0]) * int(cw)
	}

	return 0
}
