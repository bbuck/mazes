package main

import (
	"errors"
	"fmt"
	"image/color"
	"strconv"

	"github.com/bbuck/mazes/generator"
	"github.com/bbuck/mazes/rendering"
	"github.com/spf13/pflag"
)

type colorValue struct {
	color color.Color
}

func (c *colorValue) String() string {
	rr, rg, rb, ra := c.color.RGBA()
	var result string
	r := strconv.FormatUint(uint64(rr>>8), 16)
	g := strconv.FormatUint(uint64(rg>>8), 16)
	b := strconv.FormatUint(uint64(rb>>8), 16)
	a := strconv.FormatUint(uint64(ra>>8), 16)

	if ra>>8 == 255 {
		result = fmt.Sprintf("#%02s%02s%02s", r, g, b)
	} else {
		result = fmt.Sprintf("#%02s%02s%02s%02s", r, g, b, a)
	}

	return result
}

func (c *colorValue) Set(val string) error {
	if val[0] == '#' {
		val = val[1:]
	}
	var rs, gs, bs, as string
	switch len(val) {
	case 3:
		rs = val[0:1] + val[0:1]
		gs = val[1:2] + val[1:2]
		bs = val[2:] + val[2:]
		as = "ff"
	case 6:
		rs = val[:2]
		gs = val[2:4]
		bs = val[4:6]
		as = "ff"
	case 8:
		rs = val[:2]
		gs = val[2:4]
		bs = val[4:6]
		as = val[6:8]
	default:
		return errors.New("invalid color string provided")
	}
	a, err := strconv.ParseUint(as, 16, 8)
	b, err := strconv.ParseUint(bs, 16, 8)
	g, err := strconv.ParseUint(gs, 16, 8)
	r, err := strconv.ParseUint(rs, 16, 8)

	if err != nil {
		return err
	}

	c.color = color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}

	return nil
}

func (c *colorValue) Type() string {
	return "color"
}

var (
	rows        = pflag.UintP("rows", "r", 5, "The width (in cells) you wish the maze to be")
	cols        = pflag.UintP("cols", "c", 5, "The height (in cells) you wish the maze to be")
	cellWidth   = pflag.Float64P("cell-width", "w", 50, "The width (in pixels) each cell of the maze should be")
	cellHeight  = pflag.Float64P("cell-height", "h", 50, "The height (in pixels) each cell of the maze should be")
	fileName    = pflag.StringP("out-file", "o", "rendered_maze.png", "The filename the output maze image should be named, '.png' is appended to this.")
	bgColor     = colorValue{color.White}
	borderColor = colorValue{color.Black}
)

func main() {
	pflag.VarP(&bgColor, "bg-color", "b", "The background color in CSS hex (#rgb, #rrggbb, #rrggbbaa), default is white")
	pflag.VarP(&borderColor, "border-color", "B", "The border color in CSS hex (#rgb, #rrggbb, #rrggbbaa), default is black")

	pflag.Parse()

	maze := generator.Maze(generator.MazeOptions{
		Rows: *rows,
		Cols: *cols,
	})

	rendering.Maze(rendering.MazeOptions{
		Maze:            maze,
		BorderColor:     borderColor.color,
		BackgroundColor: bgColor.color,
		CellWidth:       *cellWidth,
		CellHeight:      *cellHeight,
		FileName:        *fileName,
	})
}
