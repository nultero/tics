package termeneutics

type Cmd struct {
	Name     string
	Flags    []rune
	DescStr  string
	UsageStr string
}

func (c *Cmd) AddDesc(s string) {
	c.DescStr = s
}

func (c *Cmd) AddUsage(s string) {
	c.UsageStr = s
}
