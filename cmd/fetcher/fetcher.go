package main

import (
	packageparse "fetcher/pkg/packageParse"
	"io"
	"log"
	"os"
	"path"
	"runtime"
)

func main() {
	var (
		mkFileDest, mkFileSrc *os.File
		err                   error
	)

	if !raylibSrc {
		if targetOS == "" {
			targetOS = runtime.GOOS
		}
	} else {
		targetOS = "raylib"
	}

	if err = packageparse.InstallPackages(targetOS, raylibSrc); err != nil {
		log.Fatalf("Failed to install packages:\n%s\n", err)
	}

	if customMakefile {
		log.Println("Copying custom Makefile")

		if mkFileSrc, err = os.Open(path.Join("third_party", "Makefile")); err != nil {
			log.Fatalf("Failed to open custom Makefile:\n%s\n", err)
		}
		defer mkFileSrc.Close()

		if mkFileDest, err = os.Create(path.Join("third_party", "src", "Makefile")); err != nil {
			log.Fatalf("Failed to create custom Makefile in 'src':\n%s\n", err)
		}
		defer mkFileDest.Close()

		if _, err = io.Copy(mkFileDest, mkFileSrc); err != nil {
			log.Fatalf("Failed to overwrite custom Makefile:\n%s\n", err)
		}

		log.Println("Custom Makefile copyied")
	}

	log.Println("All operations complete")
}
