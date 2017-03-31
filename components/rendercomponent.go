package components

type RenderComponent struct {
	name   string
	active bool
}

func NewRenderComponent() *RenderComponent {
	return &RenderComponent{
		name:   "Render",
		active: true,
	}
}

func (c *RenderComponent) SetActive(state bool) {
	c.active = state
}

func (c RenderComponent) GetName() string {
	return c.name
}
