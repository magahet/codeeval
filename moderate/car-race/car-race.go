package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

type Track []Segment

type Segment struct {
	V1, V2, D float64
}

func parseTrack(line string) Track {
	parts := strings.Fields(line)
	track := make(Track, len(parts)/2)
	var d, a, v1, v2 float64
	for i, part := range parts {
		if i%2 == 0 {
			d, _ = strconv.ParseFloat(part, 64)
			v1 = v2
		} else {
			a, _ = strconv.ParseFloat(part, 64)
			v2 = 1.0 - (a/180.0)
			track[i/2] = Segment{v1, v2, d}
		}
	}
	return track
}

type Car struct {
	Num int
	A, D, Vmax float64
}

func parseCars(lines []string) []Car {
	cars := make([]Car, len(lines))
	var num int
	var vMax, a, d float64
	for i, line := range lines {
		parts := strings.Fields(line)
		num, _ = strconv.Atoi(parts[0])
		vMax, _ = strconv.ParseFloat(parts[1], 64)
		a, _ = strconv.ParseFloat(parts[2], 64)
		d, _ = strconv.ParseFloat(parts[3], 64)
		cars[i] = Car{num, a, d, vMax}
	}
	return cars
}

func calcDistAndTime(v1, v2, a float64) (d, t float64) {
	t := (v2 - v1) / a
	d := (v1 + v2) * t / 2
	return
}

func timeCar(track Track, car Car) float64 {
	var totalTime, v1, v2, segDist float64
	for _, segment := range track {
		segDist = segment.D
		v1 = segment.V1 * car.Vmax
		v2 = segment.V2 * car.Vmax
		d, t := calcDistAndTime(v1, v2, car.A)
		totalTime += t
		segDist -= d
		d, t := calcDistAndTime(v1, v2, car.D)
		totalTime += t
		segDist -= d
		totalTime += segDist / car.Vmax
	}
	return totalTime
}

func race(track Track, cars []Car) string {
	return fmt.Sprintf("%.2f", timeCar(track, cars[0]))
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage:", path.Base(os.Args[0]), "file")
		os.Exit(1)
	}

	file, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		fmt.Println("error opening file", os.Args[1], ":", err)
		os.Exit(1)
	}
	
	lines := strings.Split(string(file), "\n")

	track := parseTrack(lines[0])
	cars := parseCars(lines[1:])

	fmt.Println(race(track, cars))

}