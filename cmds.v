module tics

pub struct Cmd {
pub mut:
	cmd_name string
	flags []rune
	desc_str string
	usage_str string
	num_poss_args int
	args []string
}

pub fn (mut c Cmd) add_desc_str(s string) {
	c.desc_str = s
}

pub fn (mut c Cmd) add_usage_str(s string) {
	c.usage_str = s
}

pub fn (mut c Cmd) set_num_poss_args(n int) {
	c.num_poss_args = n
}

pub fn (c Cmd) has_args() bool {
	l := c.args.len
	if l > 0 {
		if l <= c.num_poss_args {
			return true
		}
		// error for where args exceeds set num
	}

	return false
}