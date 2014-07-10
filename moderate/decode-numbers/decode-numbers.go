package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func countPermutaions(line string) int {
	l := len(line)
	count := 0
	for i, char := range line {
		if char == '1' && i+1 < l {
			count += countPermutaions(line[i+1:])
			if i+2 < l {
				count += countPermutaions(line[i+2:])
			} else {
				count++
			}
			return count
		} else if char == '2' && i+1 < l {
			if strings.ContainsAny(string(line[i+1]), "0123456") == false {
				continue
			}
			count += countPermutaions(line[i+1:])
			if i+2 < l {
				count += countPermutaions(line[i+2:])
			} else {
				count++
			}
			return count
		}
	}
	return 1
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
			fmt.Println(countPermutaions(line))
		}
	}
}
