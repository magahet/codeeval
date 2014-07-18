package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path"
	"strconv"
	"strings"
)

type Point struct {
	X, Y float64
}

type Circle struct {
	Center *Point
	Radius float64
}

func dist(p1, p2 *Point) float64 {
	a := p1.X - p2.X
	b := p1.Y - p2.Y
	return math.Sqrt(a*a + b*b)
}

func (c *Circle) Covers(p *Point) bool {
	if dist(c.Center, p) <= c.Radius {
		return true
	}
	return false
}

var r = strings.NewReplacer("Center: ", "", "Radius: ", "", "Point: ", "")

func processLine(line string) bool {
	parts := strings.Split(r.Replace(line), "; ")
	circleXYStr := strings.Split(parts[0][1:len(parts[0])-1], ", ")
	circleCenterX, _ := strconv.ParseFloat(circleXYStr[0], 64)
	circleCenterY, _ := strconv.ParseFloat(circleXYStr[1], 64)
	circleRadius, _ := strconv.ParseFloat(parts[1], 64)
	circle := &Circle{&Point{circleCenterX, circleCenterY}, circleRadius}

	pointXYStr := strings.Split(parts[2][1:len(parts[2])-1], ", ")
	pointX, _ := strconv.ParseFloat(pointXYStr[0], 64)
	pointY, _ := strconv.ParseFloat(pointXYStr[1], 64)
	point := &Point{pointX, pointY}
	return circle.Covers(point)
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
