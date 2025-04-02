//go:build darwin

package main

import "fmt"


func appendSystemEnv(env *map[string]string) (err error) {
	(*env)["GRADLE_OPTS"] = fmt.Sprintf(
		"%s '-Xdock:name=%s' '-Xdoc:icon=%s/media/gradle.icns'",
		(*env)["GRADLE_OPTS"], (*env)["APP_NAME"], (*env)["APP_HOME"],
	)
	return
}
