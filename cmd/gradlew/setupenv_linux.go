//go:build linux

package main

func appendSystemEnv(env *map[string]string) (err error) {
	// Maybe I should incerase the maximum file descriptors...
	// Use the maximum available, or set MAX_FD != -1 to use value.
	(*env)["MAX_FD"] = "maximum"
	return
}
