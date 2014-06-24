package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

func swap(line string) string {
	parts := strings.Split(line, ":")
	list := strings.Fields(parts[0])
	cmds := strings.Split(parts[1], ", ")
	for _, cmd := range cmds {
		pos := strings.Split(cmd, "-")
		a, _ := strconv.Atoi(strings.Trim(pos[0], " "))
		b, _ := strconv.Atoi(strings.Trim(pos[1], " "))
		temp := list[a]
		list[a] = list[b]
		list[b] = temp
	}
	result := ""
	for _, num := range list {
		result += string(num) + " "
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
		fmt.Println(swap(line))
	}
}
