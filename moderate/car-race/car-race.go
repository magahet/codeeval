package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
	"sort"
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
	A, D, Vmax, Time float64
}

func (c Car) String() string {
	return fmt.Sprintf("%d %.2f", c.Num, c.Time)
}

type ByTime []Car

func (c ByTime) Len() int           { return len(c) }
func (c ByTime) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ByTime) Less(i, j int) bool { return c[i].Time < c[j].Time }

func parseCars(lines []string) []Car {
	cars := make([]Car, 0)
	var num int
	var vMax, a, d float64
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) < 4 {
			continue
		}
		num, _ = strconv.Atoi(parts[0])
		vMax, _ = strconv.ParseFloat(parts[1], 64)
		vMax /= 3600.0
		a, _ = strconv.ParseFloat(parts[2], 64)
		d, _ = strconv.ParseFloat(parts[3], 64)
		cars = append(cars, Car{num, vMax/a, vMax/d, vMax, 0.0})
	}

	return cars
}

func calcDistAndTime(v1, v2, a float64) (d, t float64) {
	t = (v2 - v1) / a
	d = (v1 + v2) * t / 2
	return
}

func timeCar(track Track, car Car) float64 {
	var totalTime, segDist, t, d float64
	for _, segment := range track {
		segDist = segment.D
		d, t = calcDistAndTime(segment.V1 * car.Vmax, car.Vmax, car.A)
		totalTime += t
		segDist -= d
		d, t = calcDistAndTime(segment.V2 * car.Vmax, car.Vmax, car.D)
		totalTime += t
		segDist -= d
		totalTime += segDist / car.Vmax
	}
	return totalTime
}

func race(track Track, cars []Car) []Car {
	for i, car := range cars {
		cars[i].Time = timeCar(track, car)
	}
	sort.Sort(ByTime(cars))
	return cars
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

	for _, car := range race(track, cars) {
		fmt.Println(car)
	}
}