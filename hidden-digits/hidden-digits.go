package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
	"unicode"
)

func findDigits(line string) string {
    replacer := strings.NewReplacer("a", "0", "b", "1", "c", "2", "d", "3", "e", "4", "f", "5", "g", "6", "h", "7", "i", "8", "j", "9")
    result := ""
    for _, char := range replacer.Replace(line) {
        if unicode.IsNumber(char) {
            result += string(char)
        }
    }
    if len(result) == 0 {
        return "NONE"
    } else {
        return result
    }
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
	    fmt.Println(findDigits(line))
	}
}
