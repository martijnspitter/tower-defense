package context

type Context struct {
	Points int
	Health int
}

func NewContext() *Context {
	return &Context{
		Health: 100,
	}
}

func (c *Context) AddPoints(amount int) {
	c.Points += amount
}

func (c *Context) RemovePoints(amount int) {
	c.Points -= amount
}

func (c *Context) RemoveHealth(hit int) {
	c.Health -= hit
}

func (c *Context) ResetHealth() {
	c.Health = 100
}
