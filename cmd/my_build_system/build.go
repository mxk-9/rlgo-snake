package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

func buildProject(prName, storePath, target *string) (execName string, err error) {
	if *target != os.Getenv("GOOS") {
		if err = os.Setenv("GOOS", *target); err != nil {
			err = fmt.Errorf("Failed to set env variable 'GOOS' to '%s': %v\n", *target, err)
			return
		}
	}

	if err = os.Chdir(path.Join("cmd", *prName)); err != nil {
		return
	}

	cmd := exec.Command(
		"go", "build", "-x", "-o",
		path.Join("..", "..", *storePath+string(os.PathSeparator)),
	)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err = cmd.Run()

	fModule, err := os.Open("go.mod")
	if err != nil {
		return
	}
	defer fModule.Close()

	buf := bufio.NewScanner(fModule)

	for buf.Scan() {
		if strings.HasPrefix(buf.Text(), "module") {
			execName = strings.Split(buf.Text(), " ")[1]
			break
		}
	}

	os.Chdir(path.Join("..", ".."))
	return
}

func makeEngineDLL() (err error) {
	const (
		engine string = "raylib.dll"
	)

	var (
		engineDLL  *os.File
		destDLL    *os.File
		destEngine string = path.Join("..", "..", "build", engine)
	)

	log.Printf("\n\nBuilding engine:\n\n")
	os.Chdir(path.Join("third_party", "src"))
	cmd := exec.Command(
		"make",
		"CC=x86_64-w64-mingw32-gcc",
		"PLATFORM=PLATFORM_DESKTOP",
		"PLATFORM_OS=WINDOWS",
		"RAYLIB_LIBTYPE=SHARED",
	)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err = cmd.Run(); err != nil {
		err = fmt.Errorf("Failed to build '%s': %v\n", engine, err)
		return
	}
	log.Printf("\n\nSucces!\nNow copying to build folder:\n\n")

	engineDLL, err = os.Open(engine)
	if err != nil {
		err = fmt.Errorf("Failed to open '%s': %v\n", engine, err)
		return
	}
	defer engineDLL.Close()

	if _, err = os.Stat(path.Dir(destEngine)); err != nil {
		if err = os.MkdirAll(path.Dir(destEngine), os.ModeDir); err != nil {
			err = fmt.Errorf("Failed to create folder '%s':\n%v\n", path.Dir(destEngine), err)
			return
		}
	}

	destDLL, err = os.Create(destEngine)
	if err != nil {
		err = fmt.Errorf("Failed to create '%s': %v\n", engine, err)
	}
	defer destDLL.Close()

	if _, err = io.Copy(destDLL, engineDLL); err != nil {
		err = fmt.Errorf("Failed to copy '%s': %v\n", engine, err)
	}

	os.Chdir(path.Join("..", ".."))

	return
}
