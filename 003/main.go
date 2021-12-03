package main

import (
	"github.com/oppermax/aoc2021/pkg/utils"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type decoder struct {
	width int
	heigth int
	tracker []int
}

func newDecoder(lines []string) *decoder {
	return &decoder{
		width:  len(lines[0]),
		heigth: len(lines),
		tracker: make([]int, len(lines[0])),
	}
}

func (d *decoder) decode(lines []string, i int) {
	if i == len(lines) {
		return
	}
	d.count(lines[i], 0)
	d.decode(lines, i+1)
}

func (d *decoder) count(line string, i int) {
	if i == d.width {
		return
	}
	s, err := strconv.Atoi(string(line[i]))
	if err != nil {
		log.WithError(err)
		return
	}
	d.tracker[i] += s
	d.count(line, i+1)
}

func (d *decoder) getGammaRate() int64 {
	var t string
	for i := range d.tracker {
		if d.tracker[i] > d.heigth / 2 {
			t += "1"
		} else {
			t += "0"
		}
	}
	i, err := strconv.ParseInt(t, 2, 64)
	if err != nil {
		log.WithError(err)
		return 0
	}
	return i
}
func (d *decoder) getEpsilonRate() int64 {
	var t string
	for i := range d.tracker {
		if d.tracker[i] < d.heigth / 2 {
			t += "1"
		} else {
			t += "0"
		}
	}
	i, err := strconv.ParseInt(t, 2, 64)
	if err != nil {
		log.WithError(err)
		return 0
	}
	return i
}

func (d *decoder) getPowerConsumption() int64 {
	return d.getGammaRate() * d.getEpsilonRate()
}

func main()  {
	lines := utils.OpenFile("input.txt")
	d := newDecoder(lines)
	d.decode(lines, 0)

	log.Infof("Power consumption: %d", d.getPowerConsumption())
	log.Infof("Life support rating: %d", getLifeSupportRating(lines))



}

func getLifeSupportRating(lines []string ) int64 {
	return convToInt(getOxigenGeneratorRating(lines, 0)[0])  * convToInt(getCO2ScrubberRating(lines, 0)[0])
}

func convToInt(in string) int64 {
	out, err := strconv.ParseInt(in, 2, 64)
	if err != nil {
		log.WithError(err)
		return 0
	}
	return out
}






func getOxigenGeneratorRating(lines []string, i int) []string {

	if len(lines) == 1 {

		return lines
	}
	c := 0
	ones := []string{}
	zeros := []string{}
	for _, line := range lines {
		if strings.HasPrefix(line[i:], "1") {
			c++
			ones = append(ones, line)
		} else {
			zeros = append(zeros, line)
		}
	}
	if c >= len(lines) - c {
		return getOxigenGeneratorRating(ones, i+1)
	}
	return getOxigenGeneratorRating(zeros, i+1)
}

func getCO2ScrubberRating(lines []string, i int) []string {
	if len(lines) == 1 {

		return lines
	}
	c := 0
	ones := []string{}
	zeros := []string{}
	for _, line := range lines {
		if strings.HasPrefix(line[i:], "1") {
			c++
			ones = append(ones, line)
		} else {
			zeros = append(zeros, line)
		}
	}
	if c < (len(lines) - c) {
		return getCO2ScrubberRating(ones, i+1)
	}
	return getCO2ScrubberRating(zeros, i+1)
}