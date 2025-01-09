package main

import "flag"

var (
	destPath       string
	downloadPath   string
	overwrite      bool
	customMakefile bool
)

func init() {
	flag.StringVar(&destPath, "src", "third_party", "Folder to copy src file")
	flag.StringVar(&downloadPath, "d", "downloads", "Folder where store downloaded archive")
	flag.BoolVar(&overwrite, "o", false, "Force overwrite")
	flag.BoolVar(&customMakefile, "m", false, "Copy custom Makefile to build raylib.dll")
	flag.Parse()
}
