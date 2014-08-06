package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
)

var reSpace = regexp.MustCompile(`\s+`)
var reNonSpace = regexp.MustCompile(`\S`)

func processLine(line string) string {
	parts := strings.Split(line, ";")
	phrase := reSpace.ReplaceAllString(parts[0], " ")
	rePhrase := ".*?(" + strings.Join(strings.Fields(parts[1]), `).*?\s.*?(`) + ").*"
	re := regexp.MustCompile(rePhrase)
	positions := re.FindStringSubmatchIndex(phrase)
	if positions == nil || len(positions) < 4 {
		return "I cannot fix history"
	}
	
	phraseRedact := reNonSpace.ReplaceAllString(phrase, "_")
	newPhrase := ""
	var s, e, c int
	for i := 2; i < len(positions); i += 2 {
		s = positions[i]
		e = positions[i+1]
		newPhrase += phraseRedact[c:s]
		newPhrase += phrase[s:e]
		c = e
	}
	newPhrase += phraseRedact[c:]
	return newPhrase
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
			fmt.Println(processLine(line))
		}
	}
}
