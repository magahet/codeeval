package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
	"strconv"
	"sort"
)

func processLine(line string) string {
    line = strings.Trim(line, ";")
	tokens := strings.Split(line, "; ")
	distances := make([]int, len(tokens))

	for i, token := range tokens {
	    parts := strings.Split(token, ",")
	    dist, _ := strconv.Atoi(parts[1])
	    distances[i] = dist
	}

	sort.Ints(distances)
	//fmt.Printf("%v\n", distances)

    pos := 0
    result := ""
	for _, dist := range distances {
	    result += fmt.Sprintf("%d,", dist - pos)
	    pos = dist
	}
	return strings.Trim(result, ",")
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
