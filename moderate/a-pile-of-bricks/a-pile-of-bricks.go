package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
    "sort"
)

type Hole struct {
	P1, P2 []int
}

func (h *Hole) Delta(i int) int {
    return Abs(h.P1[i] - h.P2[i])
}

type Brick struct {
    Index int
	P1, P2 []int
}

func (b *Brick) Delta(i int) int {
    return Abs(b.P1[i] - b.P2[i])
}

func Abs(i int) int {
    if i < 0 {
        return -i
    }
    return i
}

var r = strings.NewReplacer("[", "", "]", "", "(", "", ")", "")

func parseVector(line string) []int {
    strs := strings.Split(r.Replace(line), ",")
    vector := make([]int, len(strs))
    for i, str := range strs {
        vector[i], _ = strconv.Atoi(str)
    }
    return vector
}

func parseLine(line string) (*Hole, []*Brick) {
	parts := strings.Split(line, "|")
	v := strings.Fields(parts[0])
	hole := &Hole{parseVector(v[0]), parseVector(v[1])}
	bStr := strings.Split(parts[1], ";")
    bricks := make([]*Brick, len(bStr))
    for i, b := range bStr {
        parts := strings.Fields(r.Replace(b))
        //fmt.Println(parts)
        index, _ := strconv.Atoi(parts[0])
        bricks[i] = &Brick{index, parseVector(parts[1]), parseVector(parts[2])}
    }

    return hole, bricks
}

func brickFits(b *Brick, h *Hole) bool {
    hDeltas := make([]int, 2)
    bDeltas := make([]int, 3)
    for i := 0; i < 2; i++ {
        hDeltas[i] = h.Delta(i)
    }
    for i := 0; i < 3; i++ {
        bDeltas[i] = b.Delta(i)
    }

    sort.Ints(hDeltas)
    sort.Ints(bDeltas)

    //fmt.Println(hDeltas, bDeltas)

    for i, hD := range hDeltas {
        if hD < bDeltas[i] {
            return false
        }
    }

    return true
}

func processLine(line string) string {
    hole, bricks := parseLine(line)
    result := make([]int, 0)
    for _, b := range bricks {
        if brickFits(b, hole) {
            result = append(result, b.Index)
        }
    }

    if len(result) == 0 {
        return "-"
    }
    sort.Ints(result)
    str := ""
    for _, i := range result {
        str += fmt.Sprintf("%d,", i)
    }
    return strings.Trim(str, ",")
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
