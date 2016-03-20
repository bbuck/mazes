package types

type Direction uint8

const DirNone Direction = 0

const (
	DirUp Direction = 1 << iota
	DirRight
	DirDown
	DirLeft
)
