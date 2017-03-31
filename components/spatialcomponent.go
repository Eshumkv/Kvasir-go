package components

type SpatialComponent struct {
	name   string
	active bool
}

func NewSpatialComponent() *SpatialComponent {
	return &SpatialComponent{
		name:   "Spatial",
		active: true,
	}
}

func (c *SpatialComponent) SetActive(state bool) {
	c.active = state
}

func (c SpatialComponent) GetName() string {
	return c.name
}
