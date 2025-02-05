package main

import (
	"flag"
)

var (
	customMakefile bool
	raylibSrc      bool
	targetOS       string
)

func init() {
	flag.BoolVar(&customMakefile, "m", false, "Copy custom Makefile to build raylib.dll")
	flag.BoolVar(&raylibSrc, "raylib", false, "Only download raylib's source code")
	flag.StringVar(&targetOS, "target", "", "Available targets: windows, linux, if this flag not set, uses runtime.GOOS")
	flag.Parse()
}
