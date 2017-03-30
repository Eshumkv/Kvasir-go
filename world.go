package kvasir

// World represents the world of the game
type World struct {
	columns, rows         int
	tileWidth, tileHeight int
}

// NewWorld creates a new World object.
func NewWorld(cols, rows int) World {
	return World{
		columns:    cols,
		rows:       rows,
		tileWidth:  16,
		tileHeight: 16,
	}
}

func (world *World) GetColumns() int {
	return world.columns
}
func (world *World) GetRows() int {
	return world.rows
}
func (world *World) GetTileWidth() int {
	return world.tileWidth
}
func (world *World) GetTileHeight() int {
	return world.tileHeight
}
