package main

import (
	"flag"
	"path"
)

func setupenv() (env map[string]string, err error) {
	env = make(map[string]string)
	// Add default JVM options here. You can also use JAVA_OPTS and GRADLE_OPTS
	// to pass JVM options to this script.
	env["DEFAULT_JVM_OPTS"] = ""
	env["GRADLE_OPTS"] = ""
	env["JAVA_OPTS"] = ""

	env["APP_NAME"] = "Gradle"

	env["CLASSPATH"] = path.Join("gradle", "wrapper", "gradle-wrapper.jar")

	// env["JAVACMD"] = "/usr/lib/jvm/java-17-openjdk/bin/java"
	// env["JAVA_HOME"] = "/usr/lib/jvm/java-17-openjdk/"
	env["JAVACMD"] = "java"

	err = appendSystemEnv(&env)
	return
}

func getJvmOpts(env *map[string]string) (jvmArgs []string) {
	jvmArgs = flag.Args()
	jvmArgs = append(jvmArgs, (*env)["DEFAULT_JVM_OPTS"])
	jvmArgs = append(jvmArgs, (*env)["JAVA_OPTS"])
	jvmArgs = append(jvmArgs, (*env)["GRADLE_OPTS"])
	jvmArgs = append(jvmArgs, "-Dorg.gradle.appname=gradle")
	return
}
