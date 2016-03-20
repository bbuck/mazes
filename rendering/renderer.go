package rendering

import (
	"image"
	"image/color"

	"github.com/bbuck/mazes/types"
	"github.com/llgcode/draw2d/draw2dimg"
)

// MazeOptions contain the data that is used to alter the output of the Maze
// rendering function
type MazeOptions struct {
	BackgroundColor, BorderColor color.Color
	CellWidth, CellHeight        float64
	FileName                     string
	Maze                         types.Maze
}

// Maze renders a maze object by drawing closed sides for each portion of the
// maze, rendering to an output file
func Maze(opts MazeOptions) {
	dest := image.NewRGBA(image.Rect(0, 0, opts.Maze.Width(opts.CellWidth), opts.Maze.Height(opts.CellHeight)))
	for x := 0; x < opts.Maze.Width(opts.CellWidth); x++ {
		for y := 0; y < opts.Maze.Height(opts.CellHeight); y++ {
			dest.Set(x, y, opts.BackgroundColor)
		}
	}
	gc := draw2dimg.NewGraphicContext(dest)

	gc.SetStrokeColor(opts.BorderColor)
	gc.SetLineWidth(1)

	for r, row := range opts.Maze {
		for c, col := range row {
			x := float64(c) * opts.CellWidth
			y := float64(r) * opts.CellHeight

			if col&types.DirUp != types.DirUp {
				gc.MoveTo(x, y)
				gc.LineTo(x+opts.CellWidth, y)
				gc.Close()
				gc.Stroke()
			}
			if col&types.DirLeft != types.DirLeft {
				gc.MoveTo(x, y)
				gc.LineTo(x, y+opts.CellHeight)
				gc.Close()
				gc.Stroke()
			}
			if col&types.DirRight != types.DirRight {
				gc.MoveTo(x+opts.CellWidth, y)
				gc.LineTo(x+opts.CellWidth, y+opts.CellHeight)
				gc.Close()
				gc.Stroke()
			}
			if col&types.DirDown != types.DirDown {
				gc.MoveTo(x, y+opts.CellHeight)
				gc.LineTo(x+opts.CellWidth, y+opts.CellHeight)
				gc.Close()
				gc.Stroke()
			}
		}
	}

	draw2dimg.SaveToPngFile(opts.FileName, dest)
}
