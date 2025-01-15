package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path"
)

func release(prName, target, version string) (err error) {
	var (
		binFile     = prName
		archiveName = prName
		fZip        *zip.Writer
		fArchive    *os.File
		f           *os.File
		fWrite      io.Writer
	)

	if _, err = os.Stat("releases"); err != nil {
		err = os.MkdirAll("releases", 0754)
	}

	if target == "windows" {
		binFile += ".exe"
		archiveName += "_windows-x86_64_v" + version + ".zip"
	} else if target == "linux" {
		archiveName += "_linux-x86_64_v" + version + ".zip"
	}

	fArchive, err = os.Create(path.Join("releases", archiveName))
	if err != nil {
		return
	}
	defer fArchive.Close()

	fZip = zip.NewWriter(fArchive)
	defer fZip.Close()

	fWrite, err = fZip.Create(binFile)
	if err != nil {
		return
	}

	if f, err = os.Open(path.Join("build", binFile)); err != nil {
		return
	}
	defer f.Close()

	if _, err = io.Copy(fWrite, f); err != nil {
		return
	}

	log.Println("Packed binary")

	if target == "windows" {
		if fWrite, err = fZip.Create("raylib.dll"); err != nil {
			return
		}

		if f, err = os.Open(path.Join("build", "raylib.dll")); err != nil {
			return
		}
		defer f.Close()

		if _, err = io.Copy(fWrite, f); err != nil {
			return
		}

		log.Println("Packed raylib.dll")
	}

	return
}
