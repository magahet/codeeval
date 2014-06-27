package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

func processLine(line string) string {
	lists := strings.Split(line, " | ")
	l1 := strings.Fields(lists[0])
	l2 := strings.Fields(lists[1])

	results := ""
	for i := range l1 {
		n1, _ := strconv.Atoi(l1[i])
		n2, _ := strconv.Atoi(l2[i])
		results += fmt.Sprintf("%d ", n1*n2)
	}

	return strings.Trim(results, " ")
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
