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

type submarineGen1 struct {
	loc *location
}

func (l *location) getResult() int{
	return l.x * l.y
}

func main()  {
	lines := utils.OpenFile("input.txt")
	s := submarineGen1{&location{}}
	for _, ins := range lines {
		dir, dis := findInstruction(ins)
		s.move(dir, dis)
	}
	log.Infof("Solution 1: %d", s.loc.getResult())
}

func (s *submarineGen1) moveVer(instruction int) {
	s.loc.y += instruction
}

func (s *submarineGen1)moveHor(instruction int) {
	s.loc.x += instruction
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
func (s *submarineGen1) move(dir string, dis int) {
	switch dir {
	case "up":
		s.moveVer(dis)
	case "down":
		s.moveVer(-dis)
	case "forward":
		s.moveHor(dis)
	case "backwards":
		s.moveHor(-dis)
	}
}