package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func remove(s []string, index int) []string {
	n := make([]string, len(s)-1)
	for i := 0; i < len(n); i++ {
		if i < index {
			n[i] = s[i]
		} else {
			n[i] = s[i+1]
		}
	}
	return n
}

func countChains(r rune, words []string, d int) int {
	if len(words) == 0 {
		return d
	}
	max := d
	var c int
	for i, w := range words {
		if r == rune(w[0]) {
			c = countChains(rune(w[1]), remove(words, i), d+1)
			if c > max {
				max = c
				//fmt.Printf("%c, %s, %d\n", r, w, max)
			}
		}
	}
	return max
}

func processLine(line string) string {
	fullWords := strings.Split(line, ",")
	words := make([]string, len(fullWords))
	for i, w := range fullWords {
		words[i] = w[:1] + w[len(w)-1:]
	}
	//fmt.Println(words)

	max := 1
	c := 0
	for i, w := range words {
		c = countChains(rune(w[1]), remove(words, i), 1)
		if c > max {
			max = c
		}
	}

	if max == 1 {
		return "None"
	} else {
		return fmt.Sprintf("%d", max)
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
		if line != "" {
			//fmt.Println(line)
			fmt.Println(processLine(line))
		}
	}
}
