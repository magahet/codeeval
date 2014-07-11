package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"path"
	"strings"
)

func Cmp(u1, u2 *url.URL) bool {
	if u1 == u2 {
		return true
	}

	h1 := strings.Trim(u1.Host, ":80")
	h2 := strings.Trim(u2.Host, ":80")

	if strings.ToLower(h1) != strings.ToLower(h2) {
		return false
	}

	if strings.ToLower(u1.Scheme) != strings.ToLower(u2.Scheme) {
		return false
	}

	if u1.Path != u2.Path {
		return false
	}

	return true
}

const hex = "0123456789ABCDEF"

func cleanUrl(url string) string {
	cleanUrl := ""
	for i, char := range url {
		if char == '%' {
			if i <= len(url)-2 {
				if strings.ContainsAny(string(url[i+1]), hex) && strings.ContainsAny(string(url[i+2]), hex) {
					cleanUrl += string(char)
					continue
				}
			}
			cleanUrl += "%25"
		} else {
			cleanUrl += string(char)
		}
	}

	return cleanUrl
}

func processLine(line string) string {
	rawUrls := strings.Split(line, ";")
	u1, _ := url.Parse(cleanUrl(rawUrls[0]))
	u2, _ := url.Parse(cleanUrl(rawUrls[1]))
	//fmt.Println(u1.Path, u2.Path)
	if Cmp(u1, u2) {
		return "True"
	}
	return "False"
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
