package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

var r = strings.NewReplacer("(", "", ")", "")

func processLine(line string) int {
	parts := strings.Fields(line)
	strC := strings.Split(r.Replace(parts[0]), ",")
	aveC := strings.Split(r.Replace(parts[1]), ",")
	str := make([]int, len(strC))
	ave := make([]int, len(aveC))
	aveMap := make(map[int]bool)

	for i, char := range strC {
		str[i], _ = strconv.Atoi(char)
	}

	for i, char := range aveC {
		ave[i], _ = strconv.Atoi(char)
		aveMap[ave[i]] = true
	}

	sLast := str[len(str)-1]
	aLast := ave[len(ave)-1]
	count := len(str) + len(ave)
	for _, s := range str {
		if (aLast*s)%sLast != 0 {
			continue
		}
		_, exists := aveMap[(aLast*s)/sLast]
		if exists {
			count--
		}
	}

	return count - 1
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
			//fmt.Println(line)
			fmt.Println(processLine(line))
		}
	}
}
