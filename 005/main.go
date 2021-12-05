package main

import (
	"fmt"
	"github.com/oppermax/aoc2021/pkg/utils"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strconv"
	"strings"
)

type board struct {
	entries map[position]int
	width   int
	height  int
}

type row struct {
}

type entry struct {
	pos   *position
	count *int
}

type position struct {
	x int
	y int
}

func (p *position) getSum() int {
	return p.x + p.y
}

type vent struct {
	start position
	end   position
}

func createVent(in string) vent {
	r := regexp.MustCompile(`\d+,\d+`)
	bb := r.FindAll([]byte(in), 2)
	out := []int{}
	for _, e := range bb {
		positions := strings.Split(string(e), ",")
		for _, pos := range positions {
			p, err := strconv.Atoi(pos)
			if err != nil {
				log.WithError(err).Errorf("Could not parse %s to int", pos)
				return vent{}
			}
			out = append(out, p)
		}
	}
	sumStart := out[0] + out[1]
	sumEnd := out[2] + out[3]
	if sumStart <= sumEnd {
		return vent{
			start: position{
				x: out[0],
				y: out[1],
			},
			end: position{
				x: out[2],
				y: out[3],
			},
		}
	}
	return vent{
		start: position{
			x: out[2],
			y: out[3],
		},
		end:   position{
			x: out[0],
			y: out[1],
		},
	}
}

func (b *board) addVent(v vent) {
	w := 0
	if v.start.x != v.end.x && v.start.y == v.end.y {
		for i := v.start.x; i <= v.end.x; i++ {
			b.entries[position{
				x: v.start.x + w,
				y: v.start.y,
			}] += 1
			w++
		}
	} else if v.start.x == v.end.x && v.start.y != v.end.y {
		for i := v.start.y; i <= v.end.y; i++ {
			b.entries[position{
				x: v.start.x,
				y: v.start.y + w,
			}] += 1
			w++
		}
	} else if v.start.x != v.end.x && v.start.y != v.end.y { // if line is diagonal
		poss := getDiagonalPositions(v.start, v.end)
		for _, p := range poss {
			b.entries[p]++
		}
	}
}

func getStartEndSum(a, b position) (start, end position) {
	if a.getSum() <= b.getSum() {
		return a, b
	} else {
		return b, a
	}
}

func getStartEndX(a, b position) (start, end position) {
	if a.x < b.x {
		return a, b
	} else {
		return b, a
	}
}

func getDiagonalPositions(a, b position) []position {
	out := []position{}
	w := 0

	if a.getSum() != b.getSum() { // top left to bottom right (away from 0,0)
		start, end := getStartEndSum(a, b)
		for i := start.x; i <= end.x; i++ {
			out = append(out, position{
				x: start.x + w,
				y: start.y + w,
			})
			w++
		}

	} else { // bottom left to top right (sum is always the same)
		start, end := getStartEndX(a, b)
		for i := start.x; i <= end.x; i++ {
			out = append(out, position{
				x: start.x + w,
				y: start.y - w,
			})
			w++
		}
	}
	return out
}

func (b *board) drawBoard() {
	for pos, _ := range b.entries {
		if pos.x > b.width {
			b.width = pos.x
		}
		if pos.y > b.height {
			b.height = pos.y
		}
	}
	for y := 0; y <= b.height; y++ {
		for x := 0; x <= b.width; x++ {
			if x == 0 {
				fmt.Print("\n")
			}
			fmt.Print(b.entries[position{
				x: x,
				y: y,
			}])
		}

	}
}

func (v *vent) isNotDiagonal() bool {
	return v.start.x == v.end.x || v.start.y == v.end.y
}

func main() {
	lines := utils.OpenFile("input.txt")
	b1 := board{
		entries: make(map[position]int),
	}
	b2 := board{
		entries: make(map[position]int),
	}
	for _, line := range lines {
		vent := createVent(line)
		if vent.isNotDiagonal() {
			b1.addVent(vent)
		}
		b2.addVent(vent)

	}
	overlaps1 := 0
	for _, count := range b1.entries {
		if count >= 2 {
			overlaps1++
		}
	}
	overlaps2 := 0
	for _, count := range b2.entries {
		if count >= 2 {
			overlaps2++
		}
	}
	log.Infof("Part one: %d", overlaps1)
	log.Infof("Part two: %d", overlaps2)
}
