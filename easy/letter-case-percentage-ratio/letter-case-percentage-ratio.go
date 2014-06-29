package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"unicode"
)

func processLine(line string) string {
    totalUpper := 0.0
    total := 0.0
    for _, char := range line {
        if unicode.IsLetter(char) {
            total++
            if unicode.IsUpper(char) {
                totalUpper++
            }
        }
    }
    upperRatio := totalUpper / total
    return fmt.Sprintf("lowercase: %.2f uppercase: %.2f", 100.0*(1-upperRatio), 100.0*upperRatio)
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
