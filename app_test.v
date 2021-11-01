module tics

fn test_new_app() {
	mut s := 'test'
	e := ' app'
	a := new_app('$s$e')
	s = s + e
	assert a.name == s
}