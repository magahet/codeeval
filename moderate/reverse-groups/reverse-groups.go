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
	swapCount, _ := strconv.Atoi(parts[1])
	buffer := make([]string, swapCount)
	nums := strings.Split(parts[0], ",")
	remainingStart := (len(nums) / swapCount) * swapCount
	result := ""

	for i, num := range nums {
		if i > swapCount-1 && i%swapCount == 0 {
			for j := 1; j <= swapCount; j++ {
				result += fmt.Sprintf("%s,", buffer[swapCount-j])
			}
		} else if i == remainingStart {
			break
		}
		buffer[i%swapCount] = num
	}

	for i := remainingStart; i < len(nums); i++ {
		result += fmt.Sprintf("%s,", nums[i])
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
		if line != "" {
			fmt.Println(line)
			fmt.Println(processLine(line))
		}
	}
}
