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

func unpack(archName string) (err error) {
	const sep string = string(os.PathSeparator)

	var (
		zFile         *zip.ReadCloser
		dstFile       *os.File
		fileInArchive io.ReadCloser
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

			if f.FileInfo().IsDir() {
				if err = os.MkdirAll(dest, 0754); err != nil {
					err = fmt.Errorf("Failed to create directory '%s':\n%v\n", dest, err)
					return
				}
			} else {
				baseDest := path.Dir(dest)

				if _, err = os.Stat(baseDest); err != nil {
					if err = os.MkdirAll(baseDest, 0754); err != nil {
						err = fmt.Errorf("Failed to create directory '%s':\n%v\n", baseDest, err)
						return
					}
				}

				if dstFile, err = os.Create(dest); err != nil {
					err = fmt.Errorf("Failed to create file '%s':\n%s\n", dest, err)
					return
				}
				defer dstFile.Close()

				if fileInArchive, err = f.Open(); err != nil {
					err = fmt.Errorf("Failed to open archive file '%s':\n%v\n", f.Name, err)
					return
				}
				defer fileInArchive.Close()

				if _, err = io.Copy(dstFile, fileInArchive); err != nil {
					err = fmt.Errorf("Failed to copy data from archive to file:\n%s\n", err)
					return
				}
			}
		}
	}

	return
}
