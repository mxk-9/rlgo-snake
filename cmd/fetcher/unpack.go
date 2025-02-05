package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

func unpackOneItem(f *zip.File, dest string) (err error) {
	var (
		dstFile       *os.File
		fileInArchive io.ReadCloser
	)

	if f.FileInfo().IsDir() {
		if err = os.MkdirAll(dest, 0754); err != nil {
			err = fmt.Errorf("Failed to create directory '%s':\n%s\n", dest, err)
			return
		}
	} else {
		baseDest := path.Dir(dest)

		if _, err = os.Stat(baseDest); err != nil {
			if err = os.MkdirAll(baseDest, 0754); err != nil {
				err = fmt.Errorf("Failed to create directory '%s':\n%s\n", baseDest, err)
				return
			}
		}

		if dstFile, err = os.Create(dest); err != nil {
			err = fmt.Errorf("Failed to create file '%s':\n%s\n", dest, err)
			return
		}
		defer dstFile.Close()

		if fileInArchive, err = f.Open(); err != nil {
			err = fmt.Errorf("Failed to open archive file '%s':\n%s\n", f.Name, err)
			return
		}
		defer fileInArchive.Close()

		if _, err = io.Copy(dstFile, fileInArchive); err != nil {
			err = fmt.Errorf("Failed to copy data from archive to file:\n%s\n", err)
		}
	}

	return
}

func unpack(archName, dest string) (err error) {
	var (
		zFile   *zip.ReadCloser
		destDir string = dest
	)

	if zFile, err = zip.OpenReader(archName); err != nil {
		err = fmt.Errorf("Failed to open archive '%s':\n%s\n", archName, err)
		return
	}
	defer zFile.Close()

	if destDir == "" || destDir == "." {
		destDir = "downloads"
	}

	if _, err = os.Stat(dest); err != nil {
		err = os.MkdirAll(dest, 0754)
	}

	for _, f := range zFile.File {
		to := path.Join(destDir, f.Name)

		if err = unpackOneItem(f, to); err != nil {
			return
		}
	}

	return
}

func unpackRaylib(archName string) (err error) {
	const sep string = string(os.PathSeparator)

	var (
		zFile *zip.ReadCloser
	)

	if _, err = os.Stat(path.Join("third_party", "src")); err == nil {
		log.Println("Seems archive already unpacked")
		return
	}

	if zFile, err = zip.OpenReader(archName); err != nil {
		err = fmt.Errorf("Failed to open archive '%s':\n%s\n", archName, err)
		return
	}
	defer zFile.Close()

	for _, f := range zFile.File {
		if strings.Contains(f.Name, "src"+sep) {
			arrFileDest := []string{"third_party"}
			arrFileDest = append(arrFileDest, strings.Split(f.Name, sep)[1:]...)

			dest := strings.Join(arrFileDest, sep)

			if err = unpackOneItem(f, dest); err != nil {
				return
			}
		}
	}

	return
}
