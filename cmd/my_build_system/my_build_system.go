package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
)

var cmd *exec.Cmd = &exec.Cmd{}

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

const defaultProject string = "rltest"

func main() {
	var (
		projectName  string
		execName     string
		listProjects bool
		targetOS     string
		prebuiltLib  bool
		runExe       bool
		isScript     bool
		makeRelease  string
		storeBinary  string = "build"
		matched      bool   = false
	)

	projectList, err := getProjects("cmd")
	if err != nil {
		log.Fatalf("Failed to parse 'cmd' folder: %v\n", err)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	flag.StringVar(&projectName, "name", defaultProject, "Project name")
	flag.StringVar(&targetOS, "target", "", "Available values 'windows', 'linux'")
	flag.BoolVar(&runExe, "run", false, "Launch executable after building")
	flag.BoolVar(&listProjects, "list", false, "List available projects and exit")
	flag.BoolVar(&prebuiltLib, "prebuilt-lib", false, "Use only with '-target windows', when enabled, just copying pre-built engine 'raylib.dll' from ./third_party")
	flag.BoolVar(&isScript, "script", false, "Put binary to 'scripts', not 'build'")
	flag.StringVar(&makeRelease, "release", "", "Enter a version for program")

	flag.Parse()

	if targetOS == "" {
		targetOS = runtime.GOOS
	}

	if listProjects {
		fmt.Println("Available projects:")
		for _, item := range projectList {
			fmt.Println(item)
		}
		os.Exit(0)
	}

	if makeRelease != "" {
		if err = release(projectName, targetOS, makeRelease); err != nil {
			log.Fatalf("Failed to create release:%s\n", err)
		}

		os.Exit(0)
	}

	if isScript {
		storeBinary = "scripts"
	}

	for _, item := range projectList {
		if item == projectName {
			_, isAlreadyBuilt := os.Stat(path.Join("build", "raylib.dll"))
			if projectName == defaultProject && targetOS == "windows" && !prebuiltLib && isAlreadyBuilt != nil {
				log.Println("Building engine from source")

				if _, err = os.Stat(path.Join("third_party", "src")); err != nil {
					downloaderExec := path.Join("scripts", "fetcher")
					if os.Getenv("GOOS") == "windows" {
						downloaderExec += ".exe"
					}

					if _, err = os.Stat(downloaderExec); err != nil {
						log.Println("Cannot find ", downloaderExec, ", please, compile with mbs and then try again:")
						fmt.Printf("%s -name fetcher -script\n", path.Join("scripts", "mbs"))
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

			if execName, err = buildProject(&projectName, &storeBinary, &targetOS); err != nil {
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
		ex := path.Join("build", execName)
		cmd := exec.Command(ex)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}
}
