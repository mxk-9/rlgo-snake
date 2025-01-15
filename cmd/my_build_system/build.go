package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
)

func buildProject(prName, storePath, target *string) (err error) {
	var (
		binExec = *prName
	)

	if *target == "windows" {
		binExec += ".exe"
	}
	if *target != os.Getenv("GOOS") {
		if err = os.Setenv("GOOS", *target); err != nil {
			err = fmt.Errorf("Failed to set env variable 'GOOS' to '%s': %v\n", *target, err)
			return
		}
	}

	argsToBuild := []string{
		"build", "-x", "-o", path.Join(*storePath, binExec),
	}

	if dr, localErr := os.ReadDir(path.Join("cmd", *prName)); localErr == nil {
		for _, item := range dr {
			argsToBuild = append(argsToBuild, path.Join("cmd", *prName, item.Name()))
		}
	} else {
		err = fmt.Errorf("Failed to read '%s': %v\n", path.Join("cmd", *prName), err)
		return
	}

	log.Printf("Project contains:\n%v", argsToBuild)

	cmd = exec.Command("go", argsToBuild...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err = cmd.Run()
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
	cmd = exec.Command(
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
