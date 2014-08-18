package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
    "strconv"
    "sort"
)

func hammingDist(t1, t2 string) int {
    d := len(t1)
    for i := 0; i < len(t1); i++ {
        if t1[i] == t2[i] {
            d--
        }
    }
    return d
}

func levenshteinDist(s, t string) int {
    d := make([][]int, len(s)+1)
    for i := range d {
        d[i] = make([]int, len(t)+1)
    }
    for i := range d {
        d[i][0] = i
    }
    for j := range d[0] {
        d[0][j] = j
    }
    for j := 1; j <= len(t); j++ {
        for i := 1; i <= len(s); i++ {
            if s[i-1] == t[j-1] {
                d[i][j] = d[i-1][j-1]
            } else {
                min := d[i-1][j]
                if d[i][j-1] < min {
                    min = d[i][j-1]
                }
                if d[i-1][j-1] < min {
                    min = d[i-1][j-1]
                }
                d[i][j] = min + 1
            }
        }

    }
    return d[len(s)][len(t)]
}

func processLine(line string) string {
	parts := strings.Fields(line)
    pattern := parts[0]
    m, _ := strconv.Atoi(parts[1])
    text := parts[2]
    matches := make([][]string, m+1)
    d := 0
    if m == 0 {
        for i := 0; i <= len(text) - len(pattern); i++ {
            if text[i:i+len(pattern)] == pattern {
                matches[d] = append(matches[d], pattern)
            }
        }
    } else {
        for i := 0; i <= len(text) - len(pattern); i++ {
            //d = hammingDist(text[i:i+len(pattern)], pattern)
            d = levenshteinDist(text[i:i+len(pattern)], pattern)
            if d > m {
                continue
            }
            matches[d] = append(matches[d], text[i:i+len(pattern)])
        }
    }
    result := ""
    for _, m := range matches {
        if len(m) == 0 {
            continue
        }
        sort.Strings(m)
        result += strings.Join(m, " ") + " "
    }
    if len(result) == 0 {
        return "No match"
    }
    return strings.Trim(result, " ")
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
			//processLine(line)
		}
	}
}
