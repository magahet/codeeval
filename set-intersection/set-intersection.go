package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func processLine(line string) string {
	setStrs := strings.Split(line, ";")
	set1 := strings.Split(setStrs[0], ",")
	set2 := strings.Split(setStrs[1], ",")

	result := ""
    n := 0

	for i, num := range set1 {
		if num == set2[n] {
    		result += set1[i] + ","
    		n++
		}
		if n > len(set2) - 1 {
		    break
		}
	}

	if result == "" {
    	for i, num := range set2 {
    		if num == set1[n] {
        		result += set2[i] + ","
        		n++
    		}
    		if n > len(set1) - 1 {
    		    break
    		}
    	}
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
