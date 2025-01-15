package main

import "flag"

var (
	customMakefile bool
	raylib         bool
)

func init() {
	flag.BoolVar(&customMakefile, "m", false, "Copy custom Makefile to build raylib.dll")
	flag.BoolVar(&raylib, "raylib", false, "Build raylib")
	flag.Parse()
}
