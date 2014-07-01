package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func processLine(line string, tree map[string]string) string {
	node := strings.Fields(line)
	chain := make(map[string]bool)

	for i := 0; ; i = (i + 1) % 2 {
		_, exists := chain[node[i]]

		if exists {
			return node[i]
		}

		chain[node[i]] = true
		node[i] = tree[node[i]]
	}
	return "None"
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

	tree := map[string]string{
		"30": "30",
		"8":  "30",
		"52": "30",
		"3":  "8",
		"20": "8",
		"10": "20",
		"29": "20",
	}

	for line := range readLine(file) {
		fmt.Println(processLine(line, tree))
	}
}
