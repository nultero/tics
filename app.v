module tics

pub struct App {
pub mut:
	cmds   []string
	paths map[string]string
}

pub fn new_app() App {
	return App{
		cmds: []string{}
		paths: map[string]string{}
	}
}