package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
)

var cmd *exec.Cmd = &exec.Cmd{}

func buildProject(prName, target *string) (err error) {
	binExec := path.Join("build", *prName)

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
		"build", "-x", "-o", binExec,
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
	cmd = exec.Command("make", "CC=x86_64-w64-mingw32-gcc", "PLATFORM=PLATFORM_DESKTOP", "PLATFORM_OS=WINDOWS", "RAYLIB_LIBTYPE=SHARED")

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

func getProjects(pathToProjects string) (projects []string, err error) {
	dir, err := os.ReadDir(pathToProjects)

	if err != nil {
		return
	}

	projects = make([]string, 0)
	for _, item := range dir {
		if item.IsDir() {
			projects = append(projects, item.Name())
		}
	}

	return
}

func main() {
	var (
		projectName  string
		listProjects bool
		targetOS     string
		prebuiltLib  bool
		runExe       bool
		matched      bool = false
	)

	projectList, err := getProjects("cmd")
	if err != nil {
		log.Fatalf("Failed to parse 'cmd' folder: %v\n", err)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	flag.StringVar(&projectName, "name", "remake", "Project name")
	flag.StringVar(&targetOS, "target", "", "Available values 'windows', 'linux'")
	flag.BoolVar(&runExe, "run", false, "Launch executable after building")
	flag.BoolVar(&listProjects, "list", false, "List available projects and exit")
	flag.BoolVar(&prebuiltLib, "prebuilt-lib", false, "Use only with '-target windows', when enabled, just copying pre-built engine 'raylib.dll' from ./third_party")

	flag.Parse()

	if targetOS == "" {
		targetOS = os.Getenv("GOOS")
	}

	if listProjects {
		fmt.Println("Available projects:")
		for _, item := range projectList {
			fmt.Println(item)
		}
		os.Exit(0)
	}

	for _, item := range projectList {
		if item == projectName {
			_, isAlreadyBuilt := os.Stat(path.Join("build", "raylib.dll"))
			if (projectName == "remake" || projectName == "rltest") && targetOS == "windows" && !prebuiltLib && isAlreadyBuilt != nil {
				log.Println("Building engine from source")

				if _, err = os.Stat(path.Join("third_party", "src")); err != nil {
					downloaderExec := path.Join("build", "get_raylib_src")
					if os.Getenv("GOOS") == "windows" {
						downloaderExec += ".exe"
					}

					if _, err = os.Stat(downloaderExec); err != nil {
						log.Println("Cannot find ", downloaderExec, ", please, compile with mbs and then try again:")
						fmt.Println("mbs -name get_raylib_src")
						os.Exit(1)
					}

					cmd = exec.Command(downloaderExec, "-m")
					cmd.Stderr = os.Stderr
					cmd.Stdout = os.Stdout

					if err = cmd.Run(); err != nil {
						log.Fatalln(err)
					}
				}

				if err = makeEngineDLL(); err != nil {
					log.Fatalln(err)
				}
			} else if prebuiltLib && targetOS == "windows" && isAlreadyBuilt != nil {
				log.Println("Using pre built library")
				srcLib, err := os.Open(path.Join("third_party", "raylib.dll"))
				if err != nil {
					log.Fatalf("Failed to open 'raylib.dll': %v\n", err)
				}
				defer srcLib.Close()

				if _, err := os.Stat("build"); err != nil {
					os.MkdirAll("build", os.FileMode(0754))
				}
				dstLib, err := os.Create(path.Join("build", "raylib.dll"))
				if err != nil {
					log.Fatalf("Failed to create new file: %v\n", err)
				}
				defer dstLib.Close()

				_, err = io.Copy(dstLib, srcLib)
				if err != nil {
					log.Fatalf("Failed to copy 'raylib.dll': %v\n", err)
				}
			}

			if err = buildProject(&projectName, &targetOS); err != nil {
				log.Fatal(err)
			}
			matched = true
		}
	}

	if !matched {
		errStd := fmt.Sprintf("Project '%s' not found. Available projects:", projectName)
		for _, item := range projectList {
			errStd = fmt.Sprintf("%s\n%s", errStd, item)
		}

		errStd += "\n"
		log.Fatal(fmt.Errorf("%s\n", errStd))
	}

	if runExe {
		ex := path.Join("build", projectName)
		if targetOS == "windows" {
			ex = ex + ".exe"
		}

		cmd := exec.Command(ex)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}
}
