package components

// ColorComponent defines a color
type ColorComponent struct {
	R, G, B uint8
}

// NewColorComponent returns a pointer to a new ColorComponent.
func NewColorComponent(r, g, b uint8) *ColorComponent {
	return &ColorComponent{
		R: r,
		G: g,
		B: b,
	}
}
