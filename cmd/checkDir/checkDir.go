package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func dirOutput(dr *[]os.DirEntry) {
	for i, item := range *dr {
		log.Printf("[%1d]: %v, %v\n", i, item.Name(), item.IsDir())
	}
}

func main() {
	dr, err := os.ReadDir("cmd")
	if err != nil {
		log.Fatalln(err)
	}

	dirOutput(&dr)

	fmt.Printf("\n\n")

	dr, err = os.ReadDir(path.Join("pkg", "app"))
	if err != nil {
		log.Fatalln(err)
	}

	dirOutput(&dr)
}
