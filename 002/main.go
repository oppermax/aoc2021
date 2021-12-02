package main

import (
	"github.com/oppermax/aoc2021/pkg/utils"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type location struct {
	x int
	y int
}

func (l *location) getResult() int{
	return l.x * l.y
}

func main()  {
	lines := utils.OpenFile("input.txt")
	loc := &location{}
	for _, ins := range lines {
		dir, dis := findInstruction(ins)
		move(dir, dis, loc)
	}
	log.Infof("Solution 1: %d", loc.getResult())
}

func moveVer(instruction int, loc *location) {
	loc.y += instruction
}

func moveHor(instruction int, loc *location) {
	loc.x += instruction
}

func findInstruction(line string) (string, int){
	split := strings.Split(line, " ")
	direction := split[0]
	distance, err := strconv.Atoi(split[1])
	if err != nil {
		log.WithError(err)
		return "", 0
	}
	return direction, distance
}
func move(dir string, dis int, loc *location) {
	switch dir {
	case "up":
		moveVer(dis, loc)
	case "down":
		moveVer(-dis, loc)
	case "forward":
		moveHor(dis, loc)
	case "backwards":
		moveHor(-dis, loc)
	}
}