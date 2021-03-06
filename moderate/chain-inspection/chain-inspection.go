package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func processLine(line string) string {
	linkStr := strings.Split(line, ";")
	links := make(map[string]string)
	for _, str := range linkStr {
		parts := strings.Split(str, "-")
		links[parts[0]] = parts[1]
	}

	current := "BEGIN"
	var ok bool
	for i := 0; i < len(links); i++ {
		current, ok = links[current]
		if ok == false {
			return "BAD"
		}
	}

	if current != "END" {
		return "BAD"
	}

	return "GOOD"
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
