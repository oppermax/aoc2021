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

type submarineGen2 struct {
	loc *location
	aim int
}

func (l *location) getResult() int{
	return l.x * l.y
}

func (l *location) getCurrentLocation() {
	log.Infof("X %d / Y %d", l.x, l.y)
}

func main()  {
	lines := utils.OpenFile("input.txt")
	s1 := submarineGen1{&location{}}
	s2 := submarineGen2{&location{}, 0}
	for _, ins := range lines {
		dir, dis := findInstruction(ins)
		s1.move(dir, dis)
		s2.move(dir, dis)
	}
	log.Infof("Solution 1: %d", s1.loc.getResult())
	log.Infof("Solution 2: %d", s2.loc.getResult())
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

func (s *submarineGen2) move(dir string, dis int) {
	switch dir {
	case "up":
		s.aim -= dis
	case "down":
		s.aim += dis
	case "forward":
		s.loc.x += dis
		s.loc.y += dis * s.aim
	}
	log.Info(dir," ", dis)
	s.loc.getCurrentLocation()
	s.getAim()
}

func (s *submarineGen2) getAim() {
	log.Infof("Aim %d", s.aim)
}