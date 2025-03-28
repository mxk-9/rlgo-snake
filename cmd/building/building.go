package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
)

func main() {
	var (
		target        string
		projectName   string
		projectOutput string
		projectDir    string

		verbose    bool
		runProgram bool

		args []string = []string{"build"}

		cmd *exec.Cmd
	)

	flag.StringVar(&target, "t", "linux", "Target to build")
	flag.StringVar(&projectName, "p", "", "Build another project in cmd")
	flag.BoolVar(&verbose, "v", false, "Printing commands")
	flag.BoolVar(&runProgram, "r", false, "Launch program after building")
	flag.Parse()

	if verbose {
		args = append(args, "-x", "-o")
	} else {
		args = append(args, "-o")
	}

	if projectName != "" {
		projectDir = path.Join("cmd", projectName)
		projectOutput = path.Join("build", projectName)
		projectDir = "." + string(os.PathSeparator) + projectDir
	} else {
		projectDir = "."
		projectOutput = path.Join("build", "rltest")
	}

	if target == "windows" {
		projectOutput += ".exe"
	}
	args = append(args, projectOutput)

	fmt.Printf(
		"Build info:\nTARGET: %s\nPROJECT: '%s'\nDIR: '%s'\n",
		target, projectName, projectDir,
	)

	if target == "linux" {
		args = append(args, projectDir)
		cmd = exec.Command("go", args...)
	} else if target == "windows" {
		os.Setenv("CGO_ENABLED", "1")
		os.Setenv("CC", "x86_64-w64-mingw32-gcc")
		os.Setenv("GOOS", "windows")
		os.Setenv("GOARCH", "amd64")
		ldflags := "-s -w"
		args = append(args, "-ldflags", ldflags, projectDir)

		cmd = exec.Command("go", args...)
	} else {
		fmt.Printf("Target '%s' is not supported\n", target)
		os.Exit(1)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
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
