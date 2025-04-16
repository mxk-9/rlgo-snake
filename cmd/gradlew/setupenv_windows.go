//go:build windows

package main

func appendSystemEnv(env *map[string]string) (err error) {
	// fmt.Println("TODO: appendSystemEnv")
	// err = fmt.Errorf("ImplementationNeeded")
	// _ = env

	(*env)["JAVACMD"] = "java.exe"
	return
}
