package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"
)

func main() {
	var (
		target        string
		projectName   string
		projectOutput string
		projectDir    string
		createRelease string

		verbose    bool
		runProgram bool
		clean      bool

		args          []string = []string{"build"}
		cmd           *exec.Cmd
		currentTarget TargetSystem
		err           error
	)

	flag.StringVar(&target, "t", "", "Target to build")
	flag.StringVar(&projectName, "p", "", "Build another project in cmd")
	flag.StringVar(
		&createRelease,
		"release", "",
		"Creates a folder containing all executables for all supported platforms",
	)
	flag.BoolVar(&verbose, "v", false, "Printing commands")
	flag.BoolVar(&runProgram, "r", false, "Launch program after building")
	flag.BoolVar(&clean, "c", false, "Delete project's output files")
	flag.Parse()

	// TODO:
	if clean {
		fmt.Println("This feature is not implemented yet")
		return
		// targetsToDestroy := []string{
		// 	path.Join("build"),
		// 	path.Join("scripts"),
		// }
	}

	// TODO:
	if createRelease != "" {
		fmt.Println("This feature is not implemented yet")
		return
	}

	if verbose {
		args = append(args, "-x", "-o")
	} else {
		args = append(args, "-o")
	}

	if target == "" {
		target = runtime.GOOS
	}

	currentTarget, err = getTarget(target)
	fmt.Println("Got target:", currentTarget)
	if err != nil && errors.Is(err, UnknownSystem) {
		fmt.Println(err)
		fmt.Println("Available systems:")
		for k := range supportedSystems {
			fmt.Printf("\t+ %s\n", k)
		}
		os.Exit(1)
	}

	if projectName != "" {
		projectDir = path.Join("cmd", projectName)
		projectOutput = path.Join("build", projectName)
		projectDir = "." + string(os.PathSeparator) + projectDir
	} else {
		projectDir = "."
		projectOutput = path.Join("build", "rltest")
	}

	if currentTarget == Windows {
		projectOutput += ".exe"
	}
	args = append(args, projectOutput)

	fmt.Printf(
		"Build info:\nTARGET: %s\nPROJECT: '%s'\nDIR: '%s'\n",
		target, projectName, projectDir,
	)

	switch currentTarget {
	case Linux:
		args = append(args, projectDir)
		cmd = exec.Command("go", args...)
	case Windows:
		os.Setenv("CGO_ENABLED", "1")
		os.Setenv("CC", "x86_64-w64-mingw32-gcc")
		os.Setenv("GOOS", "windows")
		os.Setenv("GOARCH", "amd64")
		ldflags := "-s -w -H=windowsgui"
		args = append(args, "-ldflags", ldflags, projectDir)

		cmd = exec.Command("go", args...)
	case Android:
		if runtime.GOOS != "linux" {
			fmt.Printf(
				"This build system cannot build apk from '%s', only from"+
					" 'linux'\n",
				runtime.GOOS,
			)
			os.Exit(1)
		}

		env := make(map[string]string)
		env["ANDROID_NDK_HOME"] = "/opt/android-ndk"
		env["ANDROID_HOME"] = "/opt/android-sdk"
		env["PATH"] = os.Getenv("PATH")
		env["PATH"] = env["ANDROID_NDK_HOME"] +
			"/toolchains/llvm/prebuilt/linux-x86_64/bin" + ":" + env["PATH"]

		env["ANDROID_SYSROOT"] = env["ANDROID_NDK_HOME"] +
			"/toolchains/llvm/prebuilt/linux-x86_64/sysroot"

		env["ANDROID_TOOLCHAIN"] = env["ANDROID_NDK_HOME"] +
			"/toolchains/arm-linux-androideabi-4.9/prebuilt/linux-x86_64"

		env["ANDROID_API"] = "26"

		env["CC"] = "armv7a-linux-androideabi" + env["ANDROID_API"] + "-clang"

		env["CGO_CFLAGS"] = fmt.Sprintf(
			"-I%s/usr/include"+" "+
				"-I%s/usr/include/arm-linux-androideabi"+" "+
				"--sysroot=%s"+" ", // +
			// "-D__ANDROID_API__=%s",
			env["ANDROID_SYSROOT"], env["ANDROID_SYSROOT"],
			env["ANDROID_SYSROOT"], /*env["ANDROID_API"],*/
		)

		env["CGO_LDFLAGS"] = fmt.Sprintf(
			"-L%s/usr/lib/arm-linux-androideabi/%s"+" "+
				"-L%s/arm-linux-androideabi/lib"+" "+
				"--sysroot=%s",
			env["ANDROID_SYSROOT"], env["ANDROID_API"],
			env["ANDROID_TOOLCHAIN"], env["ANDROID_SYSROOT"],
		)

		env["CGO_ENABLED"] = "1"
		env["GOOS"] = "android"
		env["GOARCH"] = "arm"

		ldflags := "-s -w -extldflags=-Wl,-soname,-libexample.so"
		out := "android/libs/armeabi-v7a/libexample.so"

		for k, v := range env {
			os.Setenv(k, v)
		}

		cmd = exec.Command("go", []string{
			"build", "-x", "-buildmode=c-shared", "-ldflags", ldflags, "-o=" + out,
		}...)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err = cmd.Run(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		os.Setenv("CC", "")
		os.Setenv("GOOS", runtime.GOOS)
		os.Setenv("GOARCH", runtime.GOARCH)

		execName := "./scripts/gradlew"

		args := []string{
			"build", "-o", execName,
			"./cmd/gradlew",
		}

		cmd = exec.Command("go", args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err = cmd.Run(); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("\nLaunching GRADLEW\n\n")
		cmd = exec.Command(execName)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err = cmd.Run(); err != nil {
			fmt.Println(err)
		}
		return
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if runProgram {
		cmd = exec.Command(projectOutput)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
}
