package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

func dist(p1, p2 *Point) int {
	a := p1.X - p2.X
	b := p1.Y - p2.Y
	return a*a + b*b
}

var r = strings.NewReplacer("[", "", "]", "")

func parseVector(line string) []int {
    strs := strings.Split(r.Replace(line), ",")
    vector := make([]int, len(strs)
    for i, str := range strs {
        vector[i], _ = strconv.Atoi(str)
    }
    return vector
}

func parseLine(line string) *Hole, []*Bricks {
	parts := strings.Split(line, "|")
	v := strings.Fields(parts[0])
	hole = &Hole{parseVector(v[0]), parseVector(v[1])}
	v = strings.Fields(parts[0][1:len(parts[0])-1)
	
}

func processLine(line string) bool {
    hole, bricks := parseLine(line)
	points := make([]*Point, len(parts))
	var x, y int
	for i, p := range parts {
	    xy := strings.Split(r.Replace(p), ",")
	    x, _ = strconv.Atoi(xy[0])
	    y, _ = strconv.Atoi(xy[1])
	    points[i] = &Point{x, y}
	}

	distMap := make(map[int]int)

	for _, p1 := range points {
	    for _, p2 := range points {
	        distMap[dist(p1, p2)]++
	        if len(distMap) > 3 {
	            return false
	        }
	    }
	}

    if len(distMap) != 3 {
        return false
    }

    for d, c := range distMap {
        if d == 0 && c != 4 {
            return false
        } else if d != 0 && c != 4 && c != 8 {
            return false
        }
    }
	//fmt.Println(line)
	//fmt.Println(distMap)

    return true
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
