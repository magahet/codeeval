package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
	"strconv"
)

func processLine(line string) int64 {
    char := ""
    binStr := ""
    for i, str := range strings.Fields(line) {
        // 1 or 0
        if i %2 == 0 {
            if len(str) == 1 {
                char = "0"
            } else {
                char = "1"
            }
        // how many
        } else {
            for j := 0; j < len(str); j++ {
                binStr += char
            }
        }
    }
    //fmt.Println(binStr)
    i, _ := strconv.ParseInt(binStr, 2, 64)
    return i
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
