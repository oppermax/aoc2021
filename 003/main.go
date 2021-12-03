package main

import (
	"debug/elf"
	"github.com/oppermax/aoc2021/pkg/utils"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type decoder struct {
	width int
	heigth int
	tracker []int
}

func newDecoder(lines []string) *decoder {
	for _, s := range lines {
		strconv.Atoi(s)
	}
	return &decoder{
		width:  len(lines[0]),
		heigth: len(lines),
		tracker: make([]int, len(lines[0])),
	}
}

func (d *decoder) decode() {

}

func (d *decoder) count(line string, i int) {
	if i == d.width {
		return
	}
	s, err := strconv.Atoi(string(line[0]))
	if err != nil {
		log.WithError(err)
		return
	}
	d.tracker[i] += s
	d.count(line, i+1)
}

func main()  {
	lines := utils.OpenFile("input.txt")
	d := newDecoder(lines)

}

