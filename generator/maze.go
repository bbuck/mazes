package generator

import (
	"math/rand"
	"time"

	"github.com/bbuck/mazes/types"
)

// TestMaze is a manually created maze used to test rendering and other things
// with a statically defined maze.
var TestMaze = types.Maze{
	types.MazeRow{types.DirRight, types.DirLeft | types.DirDown, types.DirRight, types.DirLeft | types.DirRight, types.DirLeft | types.DirDown},
	types.MazeRow{types.DirDown, types.DirUp | types.DirRight, types.DirLeft | types.DirDown, types.DirDown | types.DirRight, types.DirLeft | types.DirUp},
	types.MazeRow{types.DirUp | types.DirDown | types.DirRight, types.DirLeft | types.DirRight, types.DirUp | types.DirLeft, types.DirDown | types.DirUp | types.DirRight, types.DirDown | types.DirLeft},
	types.MazeRow{types.DirUp | types.DirRight, types.DirLeft | types.DirRight, types.DirLeft | types.DirDown, types.DirUp, types.DirUp | types.DirDown},
	types.MazeRow{types.DirRight, types.DirLeft | types.DirRight, types.DirUp | types.DirLeft | types.DirRight, types.DirLeft | types.DirRight, types.DirLeft | types.DirUp},
}

// MazeOptions defines options required by the maze generator to perform it's
// generation.
type MazeOptions struct {
	Rows, Cols uint
}

// Maze generates a maze using a breadth-first algorithm for building the
// pathways.
func Maze(opts MazeOptions) types.Maze {
	maze := newMaze(opts)
	st := newStack()
	st = st.push(position{0, 0})
	gen := newRandomGenerator()

	for !st.empty() {
		var cur position
		st, cur = st.pop()
		neighbors := getNeighbors(maze, cur)
		if len(neighbors) == 0 {
			continue
		}
		neighbor := neighbors[gen.Intn(len(neighbors))]
		switch {
		case neighbor.row > cur.row:
			maze[cur.row][cur.col] |= types.DirDown
			maze[neighbor.row][neighbor.col] |= types.DirUp
		case neighbor.row < cur.row:
			maze[cur.row][cur.col] |= types.DirUp
			maze[neighbor.row][neighbor.col] |= types.DirDown
		case neighbor.col > cur.col:
			maze[cur.row][cur.col] |= types.DirRight
			maze[neighbor.row][neighbor.col] |= types.DirLeft
		case neighbor.col < cur.col:
			maze[cur.row][cur.col] |= types.DirLeft
			maze[neighbor.row][neighbor.col] |= types.DirRight
		}

		st = st.push(cur).push(neighbor)
	}

	maze[gen.Intn(len(maze))][0] |= types.DirLeft
	maze[gen.Intn(len(maze))][len(maze[0])-1] |= types.DirRight

	return maze
}

func newRandomGenerator() *rand.Rand {
	src := rand.NewSource(time.Now().UnixNano())

	return rand.New(src)
}

func getNeighbors(maze types.Maze, cur position) []position {
	possible := []position{
		{cur.row - 1, cur.col},
		{cur.row + 1, cur.col},
		{cur.row, cur.col - 1},
		{cur.row, cur.col + 1},
	}
	var neighbors []position
	for _, pos := range possible {
		if pos.row >= 0 && pos.row < len(maze) && pos.col >= 0 && pos.col < len(maze[0]) && maze[pos.row][pos.col] == types.DirNone {
			neighbors = append(neighbors, pos)
		}
	}

	return neighbors
}

func newMaze(opts MazeOptions) types.Maze {
	maze := make(types.Maze, opts.Rows)
	for i, rows := 0, int(opts.Rows); i < rows; i++ {
		maze[i] = make(types.MazeRow, opts.Cols)
	}

	return maze
}
