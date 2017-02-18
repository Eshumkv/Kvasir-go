package components

type ColorComponent struct {
	R, G, B uint8
}

func NewColorComponent(r, g, b uint8) *ColorComponent {
	return &ColorComponent{
		R: r,
		G: g,
		B: b,
	}
}
