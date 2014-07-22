package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

var r = strings.NewReplacer("a", "y", "b", "h", "c", "e", "d", "s", "e", "o", "f", "c", "g", "v", "h", "x", "i", "d", "j", "u", "k", "i", "l", "g", "m", "l", "n", "b", "o", "k", "p", "r", "q", "z", "r", "t", "s", "n", "t", "w", "u", "j", "v", "p", "w", "f", "x", "m", "y", "a", "z", "q")

func processLine(line string) string {
	return r.Replace(line)
}

func readLine(file *os.File) <-chan string {
	out := make(chan string)
	go func() {
		in := bufio.NewReader(file)
		linePartial := ""
		for {
			bytes, isPrefix, err := in.ReadLine()
			if err != nil {
				break
			} else if isPrefix {
				linePartial += string(bytes)
			} else {
				out <- linePartial + string(bytes)
				linePartial = ""
			}
		}
		close(out)
	}()
	return out
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage:", path.Base(os.Args[0]), "file")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	defer file.Close()

	if err != nil {
		fmt.Println("error opening file", os.Args[1], ":", err)
		os.Exit(1)
	}

	for line := range readLine(file) {
		if line != "" {
			fmt.Println(processLine(line))
		}
	}
}
