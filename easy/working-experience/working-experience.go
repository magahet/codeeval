package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
	"time"
)

func iterDates(startStr, endStr string) <-chan time.Time {
    const dateFormat = "Jan 2006"
    out := make(chan time.Time)
    go func() {
        start, _ := time.Parse(dateFormat, startStr)
        end, _ := time.Parse(dateFormat, endStr)
        for now := start; now != end; now = now.AddDate(0, 1, 0) {
            out <- now
        }
        out <- end
        close(out)
    }()
    return out
}

func processLine(line string) int {
    monthsWorked := make(map[time.Time]bool)
    for _, dateRanges := range strings.Split(line, "; ") {
        dates := strings.Split(dateRanges, "-")
        for date := range iterDates(dates[0], dates[1]) {
            monthsWorked[date] = true
        }
    }
    return len(monthsWorked) / 12
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
