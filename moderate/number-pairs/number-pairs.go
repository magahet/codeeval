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
    parts := strings.Split(line, ";")
    sum, _ := strconv.Atoi(parts[1])
    numStrs := strings.Split(parts[0], ",")
    numSl := make([]int, len(numStrs))
    numSet := make(map[int]bool)
    usedSet := make(map[int]bool)

    for i, char := range numStrs {
        n, _ := strconv.Atoi(char)
        numSl[i] = n
        numSet[n] = true
    }

    result := ""
    var exists bool
    for _, n := range numSl {
        _, exists = usedSet[n]
        if exists || sum == n*2 {
            break
        }
        _, exists = numSet[sum - n]
        if exists {
            usedSet[n] = true
            usedSet[sum - n] = true
            result += fmt.Sprintf("%d,%d;", n, sum - n)
        }
    }

    if result == "" {
        return "NULL"
    }
    return strings.Trim(result, ";")
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
        //fmt.Println(line)
		fmt.Println(processLine(line))
        //fmt.Println("")
	}
}
