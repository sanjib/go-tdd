package sync

type Counter struct {
	value int
}

func (c *Counter) Add() {
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
