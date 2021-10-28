package termeneutics

func (a Arg) HasFlag(r rune) bool {
	for _, f := range a.Flags {
		if f == r {
			return true
		}
	}

	return false
}
