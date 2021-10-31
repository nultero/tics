package termeneutics

import "fmt"

func (c *Cmd) throwArgsOverflowErr() {
	s := fmt.Sprintf("%v cmd `%v` passed too many args", redError(), c.Name)
	quit(s)
}
