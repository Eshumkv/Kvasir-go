package parts

// CameraInterface defines the interface a camera has to follow.
type CameraInterface interface {
	GetX() int32
	GetY() int32
	SetX(x int32)
	SetY(y int32)
	GetScreenLocation(x, y float64) (int, int)
}
