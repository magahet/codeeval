package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage:", path.Base(os.Args[0]), "file")
		os.Exit(1)
	}

	fileStat, _ := os.Stat(os.Args[1])
	fmt.Println(fileStat.Size())
}
