package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

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

func parseLine(line string) (a int, b int, N int) {
	tokens := strings.Fields(line)
	a, _ = strconv.Atoi(tokens[0])
	b, _ = strconv.Atoi(tokens[1])
	N, _ = strconv.Atoi(tokens[2])
	return
}

func generatePhrase(a int, b int, N int) {
	for i := 1; i <= N; i++ {
		switch {
		default:
			fmt.Printf("%d", i)
		case i%a == 0 && i%b == 0:
			fmt.Print("FB")
		case i%a == 0:
			fmt.Print("F")
		case i%b == 0:
			fmt.Print("B")
		}
		if i != N {
			fmt.Print(" ")
		}
	}
	fmt.Println()
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
		a, b, N := parseLine(line)
		generatePhrase(a, b, N)
	}
}
