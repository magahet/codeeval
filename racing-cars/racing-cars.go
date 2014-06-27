package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func processLine(line, lastLine string) string {
    srcI := strings.Index(lastLine, "C")
    if srcI == -1 {
        srcI = strings.Index(lastLine, "_")
    }

    dstI := strings.Index(line, "C")
    if dstI == -1 {
        dstI = strings.Index(line, "_")
    }

    var char string
    if srcI < dstI {
        char = "\\"
    } else if srcI > dstI {
        char = "/"
    } else {
        char = "|"
    }

    return line[:dstI] + char + line[dstI+1:]
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

    lastLine := ""
	for line := range readLine(file) {
        if lastLine == "" {
            lastLine = line
        }
        fmt.Println(processLine(line, lastLine))
        lastLine = line
	}
}
