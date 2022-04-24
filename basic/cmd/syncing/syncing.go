package syncing

type Counter struct {
	value int
}

type Hoge string

func (c *Counter) Inc() {
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
