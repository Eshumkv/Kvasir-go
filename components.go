package kvasir

type ColorComponent struct {
	r, g, b uint8
}

func NewColorComponent(r, g, b uint8) *ColorComponent {
	return &ColorComponent{
		r: r,
		g: g,
		b: b,
	}
}
