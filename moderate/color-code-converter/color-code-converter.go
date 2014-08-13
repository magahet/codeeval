package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"math"
)

type RGB struct {
	R, G, B int
}

func (rgb RGB) String () string {
	return fmt.Sprintf("RGB(%d,%d,%d)", rgb.R, rgb.G, rgb.B)
}

func parseHSL(line string) RGB {
	parts := strings.Split(line[4:len(line)-1], ",")
	h, _ := strconv.ParseFloat(parts[0], 64)
	s, _ := strconv.ParseFloat(parts[1], 64)
	l, _ := strconv.ParseFloat(parts[2], 64)
	h /= 360.0
	s /= 100.0
	l /= 100.0
	
    var fR, fG, fB float64
    if s == 0 {
            fR, fG, fB = l, l, l
    } else {
            var q float64
            if l < 0.5 {
                    q = l * (1 + s)
            } else {
                    q = l + s - s*l
            }
            p := 2*l - q
            fR = hueToRGB(p, q, h+1.0/3)
            fG = hueToRGB(p, q, h)
            fB = hueToRGB(p, q, h-1.0/3)
    }
    var r, g, b int
    r, g, b = float64ToInt(fR), float64ToInt(fG), float64ToInt(fB)

    return RGB{r, g, b}
}



// hueToRGB is a helper function for HSLToRGB.
func hueToRGB(p, q, t float64) float64 {
    if t < 0 {
            t += 1
    }
    if t > 1 {
            t -= 1
    }
    if t < 1.0/6 {
            return p + (q-p)*6*t
    }
    if t < 0.5 {
            return q
    }
    if t < 2.0/3 {
            return p + (q-p)*(2.0/3-t)*6
    }
    return p
}

func float64ToInt(x float64) int {
    if x < 0 {
            return 0
    }
    if x > 1 {
            return 255
    }
    return int(x*255 + 0.5)
}



func parseHSV(line string) RGB {
	parts := strings.Split(line[4:len(line)-1], ",")
	h, _ := strconv.ParseFloat(parts[0], 64)
	h /= 360.0
	s, _ := strconv.ParseFloat(parts[1], 64)
	s /= 100.0
	v, _ := strconv.ParseFloat(parts[2], 64)
	v /= 100.0
	
	var fR, fG, fB float64
    i := math.Floor(h * 6)
    f := h*6 - i
    p := v * (1.0 - s)
    q := v * (1.0 - f*s)
    t := v * (1.0 - (1.0-f)*s)
    switch int(i) % 6 {
    case 0:
            fR, fG, fB = v, t, p
    case 1:
            fR, fG, fB = q, v, p
    case 2:
            fR, fG, fB = p, v, t
    case 3:
            fR, fG, fB = p, q, v
    case 4:
            fR, fG, fB = t, p, v
    case 5:
            fR, fG, fB = v, p, q
    }
    r := int((fR * 255) + 0.5)
    g := int((fG * 255) + 0.5)
    b := int((fB * 255) + 0.5)
    return RGB{r, g, b}
}

func parseCMYK(line string) RGB {
	parts := strings.Split(line[1:len(line)-1], ",")
	c, _ := strconv.ParseFloat(parts[0], 64)
	m, _ := strconv.ParseFloat(parts[1], 64)
	y, _ := strconv.ParseFloat(parts[2], 64)
	k, _ := strconv.ParseFloat(parts[3], 64)
	
	r := 255 * (1-c) * (1-k) + 0.5
	g := 255 * (1-m) * (1-k) + 0.5
	b := 255 * (1-y) * (1-k) + 0.5
	return RGB{int(r), int(g), int(b)}
}

func parseHex(line string) RGB {
	if rgb, err := strconv.ParseUint(line[1:], 16, 32); err == nil {
    	return RGB{int(rgb >> 16), int((rgb >> 8) & 0xFF), int(rgb & 0xFF)}
	}
	return RGB{0, 0, 0}
}

func processLine(line string) RGB {
	var c RGB
	if strings.HasPrefix(line, "HSL") {
		c = parseHSL(line)
	} else if strings.HasPrefix(line, "HSV") {
		c = parseHSV(line)
	} else if strings.HasPrefix(line, "(") {
		c = parseCMYK(line)
	} else if strings.HasPrefix(line, "#") {
		c = parseHex(line)
	}
	return c
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
