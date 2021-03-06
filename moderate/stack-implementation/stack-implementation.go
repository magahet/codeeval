package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"path"
	"strings"
)

func processLine(line string) string {
	l := list.New()
	for _, char := range strings.Fields(line) {
		l.PushFront(char)
	}

	result := ""
	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		if i%2 == 0 {
			result += fmt.Sprintf("%s ", e.Value)
		}
		i++
	}
	return strings.Trim(result, " ")
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
		fmt.Println(processLine(line))
	}
}
