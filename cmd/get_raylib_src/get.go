package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const raylibSrcPath = "https://github.com/raysan5/raylib/archive/refs/tags/5.5.zip"

type Downloader struct {
	Client *http.Client
}

func (d *Downloader) GetFile(link, destPath string, overwrite bool) (err error) {
	if _, err = os.Stat(destPath); err == nil && !overwrite {
		fmt.Printf("File '%s' exists, skip downloading\n", destPath)
		err = nil
		return
	}

	var resp *http.Response
	if resp, err = d.Client.Get(link); err != nil {
		return
	}
	defer resp.Body.Close()

	src := &PassThru{Reader: resp.Body}

	currDir := filepath.Dir(destPath)

	if currDir != "" && currDir != "." && currDir != "."+string(os.PathSeparator) {
		err = os.MkdirAll(currDir, 0754)
		if err != nil {
			err = fmt.Errorf("Failed to create a folder:\n%s\n", err)
			return
		}
	}

	out, err := os.Create(destPath)
	if err != nil {
		return
	}
	defer out.Close()

	fmt.Println()
	if _, err = io.Copy(out, src); err != nil {
		err = fmt.Errorf("Failed to download a file:\n%s\n", err)
	}
	fmt.Println()

	return
}
