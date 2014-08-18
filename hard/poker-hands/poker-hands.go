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

var r = strings.NewReplacer("T", "10", "J", "11", "Q", "12", "K", "13", "A", "14")

func parseCard(cardString string) (int, rune) {
	v, _ := strconv.Atoi(r.Replace(cardString[:1]))
	return v, rune(cardString[1])
}

type Hand struct {
	Cards []int
	IsFlush bool
	IsStraight bool
}


func (h Hand) GetCardFreq() (map[int]int, []int) {
	f := make(map[int]int)
	for _, i := range h.Cards {
		f[i]++
	}
	counts := make([]int, len(f))
	i := 0
	for _, c := range f {
		counts[i] = c
		i++
	}
	sort.Ints(counts)
	return f, counts
}

func (h Hand) GetRankValue() int {
	if h.IsStraight && h.IsFlush {
		if h.Cards[0] == 10 {
			return 9
		} else {
			return 8
		}
	}
	
	_, counts := h.GetCardFreq()
	fmt.Println(counts)
	
	if counts[len(counts)-1] == 4 {
		return 7
	}
	
	if counts[len(counts)-1] == 3 {
		if counts[0] == 2 {
			return 6
		} else {
			return 3
		}
	}
	
	if h.IsFlush {
		return 5
	}
	
	if h.IsStraight {
		return 4
	}
	
	if counts[len(counts)-1] == 2 {
		return counts[len(counts)-2]
	}
		
	return 0
}
	
func IsConsecutive(seq []int) bool {
	for i := 0; i < len(seq); i++ {
		if seq[i] != seq[0] + i {
			return false
		}
	}
	return true
}

func newHand(cardStrings []string) Hand {
	cards := make([]int, len(cardStrings))
	suits := map[rune]int{'H': 0, 'D': 0, 'C': 0, 'S': 0}
	var suit rune
	for i, cardStr := range cardStrings {
		cards[i], suit = parseCard(cardStr)
		suits[suit]++
	}
	sort.Ints(cards)
	isFlush := false
	for _, c := range suits {
		if c == 5 {
			isFlush = true
		}
	}
	return Hand{cards, isFlush, IsConsecutive(cards)}
}

func processLine(line string) string {
	parts := strings.Fields(line)
	leftHand := newHand(parts[:5])
	rightHand := newHand(parts[5:])
	
	fmt.Println(leftHand, rightHand)
	fmt.Println(leftHand.GetRankValue(), rightHand.GetRankValue())
	/*
	leftValue := leftHand.GetValue()
	rightValue := leftHand.GetValue()
	
	if leftValue > rightValue {
		return "left"
	} else if leftValue < rightValue {
		return "right"
	}
	*/
	
	return "none"
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
