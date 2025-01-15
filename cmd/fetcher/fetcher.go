package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

const archiveName string = raylibVersion + ".zip"

func main() {
	var (
		mkFileDest, mkFileSrc *os.File
		err                   error
	)

	if raylib {
		dstArchPath := path.Join("downloads", archiveName)
		d := &Downloader{Client: &http.Client{}}

		log.Println("Downloading archive")
		if err := d.GetFile(raylibSrcPath, dstArchPath); err != nil {
			log.Fatalf("Failed to download raylib's source code:\n%s\n", err)
		}

		fmt.Println()
		log.Println("Unpacking 'src' folder into 'third_party' directory")
		unpack(dstArchPath)

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

	}

	log.Println("All operations complete")
}
