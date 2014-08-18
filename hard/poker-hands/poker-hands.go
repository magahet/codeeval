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

func cardFreqToCountMap(freq map[int]int) map[int][]int {
    c := make(map[int][]int)
    for i := 1; i <= 4; i++ {
        s := make([]int, 0)
        for card, count := range freq {
            if count == i {
                s = append(s, card)
            }
        }
        if len(s) > 0 {
            c[i] = s
        }
    }
    return c
}

func GetRankValue(counts []int, h Hand) int {
	if h.IsStraight && h.IsFlush {
		if h.Cards[0] == 10 {
			return 9
		} else {
			return 8
		}
	}

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

func maxN(s []int, n int) []int {
    sort.Ints(s)
    return s[len(s)-n:]
}

func GetMajorMinorCards(countMap map[int][]int) (major, minor int) {
    var mm []int
    _, ok := countMap[4]
    if ok {
        return countMap[4][0], countMap[1][0]
    }
    _, ok = countMap[3]
    if ok {
        _, ok = countMap[2]
        if ok {
            return countMap[3][0], countMap[2][0]
        } else {
            return countMap[3][0], maxN(countMap[1], 1)[0]
        }
    }
    _, ok = countMap[2]
    if ok {
        if len(countMap[2]) == 2 {
            mm = maxN(countMap[2], 2)
            return mm[1], mm[0]
        } else {
            return countMap[2][0], maxN(countMap[1], 1)[0]
        }
    }

    mm = maxN(countMap[1], 2)
    return mm[1], mm[0]
}

func (h Hand) GetValue() uint64 {
	freq, counts := h.GetCardFreq()
    rankValue := GetRankValue(counts, h)
    countMap := cardFreqToCountMap(freq)
    major, minor := GetMajorMinorCards(countMap)
    //fmt.Println("parts:", rankValue, major, minor)
    vStr := fmt.Sprintf("%d%02d%02d%02d%02d%02d%02d%02d", rankValue, major, minor, h.Cards[4], h.Cards[3], h.Cards[2], h.Cards[1], h.Cards[0])
    //fmt.Println("vstr:", vStr)
    v, _ := strconv.ParseUint(vStr, 10, 64)
    return v
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

	//fmt.Println(leftHand, rightHand)
	leftValue := leftHand.GetValue()
	rightValue := rightHand.GetValue()
	//fmt.Println(leftValue, rightValue)

	if leftValue > rightValue {
		return "left"
	} else if leftValue < rightValue {
		return "right"
	}

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
