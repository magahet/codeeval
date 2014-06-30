package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
	"strconv"
)

type Point struct {
    x, y int
}

type Rect struct {
    tl, br Point
}

func overlapping(r1, r2 *Rect) bool {
    // r1 tl corner inside r2
    xOverlap := (r1.br.x >= r2.tl.x) && (r1.tl.x <= r2.br.x)
    yOverlap := (r1.br.y <= r2.tl.y) && (r1.tl.y >= r2.br.y)

    //fmt.Printf("%v, %v\n", r1, r2)
    //fmt.Printf("%v, %v\n", xOverlap, yOverlap)
    return xOverlap && yOverlap
}

func processLine(line string) string {
    points := make([]Point, 4)
    var x, y int
    for i, char := range strings.Split(line, ",") {
        if i%2 == 0 {
            x, _ = strconv.Atoi(char)
        } else {
            y, _ = strconv.Atoi(char)
            points[i/2] = Point{x, y}
        }
    }
    rectA := Rect{points[0], points[1]}
    rectB := Rect{points[2], points[3]}
    if overlapping(&rectA, &rectB) {
        return "True"
    } else {
        return "False"
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
		fmt.Println(processLine(line))
	}
}
