package main

import (
	"github.com/oppermax/aoc2021/pkg/utils"
	log "github.com/sirupsen/logrus"
	"strings"
)

type display struct {
	patterns       []string
	output         []string
	patternToValue map[string]int
	valueToPattern map[int]string
}

func readDisplays(in []string) []*display {
	out := make([]*display, len(in))
	for i, line := range in {
		split := strings.Split(line, "|")
		patternsUnsorted := strings.Split(strings.TrimSpace(split[0]), " ")
		patterns := make([]string, len(patternsUnsorted))
		for d, pattern := range patternsUnsorted {
			patterns[d] = utils.SortStringByCharacter(pattern)
		}
		outPutUnsorted := strings.Split(strings.TrimSpace(split[1]), " ")
		output := make([]string, len(outPutUnsorted))
		for d, o := range outPutUnsorted {
			output[d] = utils.SortStringByCharacter(o)
		}
		out[i] = &display{
			patterns:       patterns,
			output:         output,
			patternToValue: make(map[string]int),
			valueToPattern: make(map[int]string),
		}
	}
	return out
}

func setupUniques(displays []*display) {
	for _, display := range displays {
		for _, pattern := range display.patterns {
			switch len(pattern) {
			case 2:
				display.patternToValue[pattern] = 1
				display.valueToPattern[1] = pattern
			case 3:
				display.patternToValue[pattern] = 7
				display.valueToPattern[7] = pattern

			case 4:
				display.patternToValue[pattern] = 4
				display.valueToPattern[4] = pattern

			case 7:
				display.patternToValue[pattern] = 8
				display.valueToPattern[8] = pattern
			}
		}
	}
}

func partOne(displays []*display) int {
	out := 0
	for _, d := range displays {
		for _, output := range d.output {
			if _, ok := d.patternToValue[output]; ok {
				out++
			}
		}
	}
	return out
}

// patterns with len 5: 2,3,5
// patterns with len 6: 0,6,9


func partTwo(displays []*display) int {
	out := 0
	for _, d := range displays {
		for _, pattern := range d.patterns {
			switch len(pattern) {
			case 6:
				// find 9
				found := true
				for _, letter := range d.valueToPattern[4] {
					if !strings.Contains(pattern, string(letter)) {
						found = false
					}
				}
				if found {
					d.patternToValue[pattern] = 9
					d.valueToPattern[9] = pattern
					log.Infof("Found a 9 %s", pattern)
					continue
				}
				// find 0
				found = true
				for _, letter := range d.valueToPattern[1] {
					if !strings.Contains(pattern, string(letter)) {
						log.Info(string(letter), " ", pattern)
						found = false
					}
				}
				if found {
					d.patternToValue[pattern] = 0
					d.valueToPattern[0] = pattern
					log.Infof("Found a 0 %s", pattern)
					continue
				}
				// find 6
				d.patternToValue[pattern] = 6
				d.valueToPattern[6] = pattern
				log.Infof("Found a 6 %s", pattern)
			}
		}
		for _, pattern := range d.patterns {
			switch len(pattern) {
			case 5:
				// find 5
				found := true
				for _, letter := range pattern {
					if !strings.Contains(d.valueToPattern[6], string(letter)) {
						found = false
					}
				}
				if found {
					d.patternToValue[pattern] = 5
					d.valueToPattern[5] = pattern
					log.Infof("Found a 5 %s", pattern)
					continue
				}
				// find 3
				found = true
				for _, letter := range d.valueToPattern[1] {
					if !strings.Contains(pattern, string(letter)) {
						found = false
					}
				}
				if found {
					d.patternToValue[pattern] = 3
					d.valueToPattern[3] = pattern
					log.Infof("Found a 3 %s", pattern)
					continue
				}
				// find 2
				d.patternToValue[pattern] = 2
				d.valueToPattern[2] = pattern
				log.Infof("Found a 2 %s", pattern)
			}
		}
		for i, outp := range d.output {
			switch i {
			case 0:
				out += d.patternToValue[outp] * 1000
			case 1:
				out += d.patternToValue[outp] * 100
			case 2:
				out += d.patternToValue[outp] * 10
			case 3:
				out += d.patternToValue[outp]
			}

		}
	}

	return out
}

func main() {
	lines := utils.OpenFile("input.txt")
	displays := readDisplays(lines)
	setupUniques(displays)
	log.Info(partOne(displays))
	log.Info(partTwo(displays))
}
