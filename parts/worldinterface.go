package parts

// WorldInterface represents the world of the game
type WorldInterface interface {
	GetColumns() int
	GetRows() int
	GetTileWidth() int
	GetTileHeight() int
}
