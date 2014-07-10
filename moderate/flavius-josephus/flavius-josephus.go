package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

func processLine(line string) string {
	parts := strings.Split(line, ",")
	population, _ := strconv.Atoi(parts[0])
	step, _ := strconv.Atoi(parts[1])
	r := ring.New(population)
	result := ""

	for i := 0; i < r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}
	r = r.Prev()

	for r.Len() > 1 {
		r = r.Move(step - 1)
		result += fmt.Sprintf("%d ", r.Unlink(1).Value)
	}

	result += fmt.Sprintf("%d ", r.Value)

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
		if line != "" {
			fmt.Println(processLine(line))
		}
	}
}
