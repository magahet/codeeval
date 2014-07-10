package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func processLine(line string) string {
	parts := strings.Split(line, ",")
	a := parts[0]
	b := parts[1]

	if len(a) != len(b) {
		return "False"
	}

	for i := 0; i < len(a); i++ {
		if a == b[i:]+b[:i] {
			return "True"
		}
	}

	return "False"
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
