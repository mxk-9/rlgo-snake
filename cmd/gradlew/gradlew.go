package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var (
		env map[string]string
		err    error
	)

	if env, err = setupenv(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	gradleArgs := make([]string, 0)
	gradleArgsUnclean := getJvmOpts(&env)
	for _, item := range gradleArgsUnclean {
		if item != "" {
			fmt.Println("GRADLE:", item)
			gradleArgs = append(gradleArgs, item)
		}
	}
	
	gradleArgs = append(gradleArgs, "-classpath")
	gradleArgs = append(gradleArgs, env["CLASSPATH"])
	gradleArgs = append(gradleArgs, "org.gradle.wrapper.GradleWrapperMain")
	gradleArgs = append(gradleArgs, "assembleDebug")

	cmd := exec.Command(env["JAVACMD"], gradleArgs...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	fmt.Println("CMD:", env["JAVACMD"], gradleArgs)
	if err = cmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
