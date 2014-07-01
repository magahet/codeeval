package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

func processLine(line string) {
	items := strings.Fields(line)
	//fmt.Printf("%v\n", items)
	rawIndex, _ := strconv.Atoi(items[len(items)-1])
	index := len(items) - rawIndex - 1
	if index >= 0 {
		fmt.Println(items[index])
	}
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
		processLine(line)
	}
}
