package termeneutics

func (c *Cmd) AddFlags(flags []rune) {
	c.Flags = append(c.Flags, flags...)
}

func (c *Cmd) HasFlag(r rune) bool {
	for _, f := range c.Flags {
		if f == r {
			return true
		}
	}

	return false
}
