package termeneutics

type Cmd struct {
	Name        string
	Flags       []rune
	DescStr     string
	UsageStr    string
	NumPossArgs int
	Args        []string
}

func (c *Cmd) AddDesc(s string) {
	c.DescStr = s
}

func (c *Cmd) AddUsage(s string) {
	c.UsageStr = s
}

func (c *Cmd) HasArgs() bool {
	if c.Args != nil {
		l := len(c.Args)
		if l > 0 {
			if l <= c.NumPossArgs {
				return true
			} else {
				c.throwArgsOverflowErr()
			}
		}
	}
	return false
}

func (c *Cmd) SetNumPossArgs(n int) {
	c.NumPossArgs = n
}
