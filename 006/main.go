package main

import (
	"github.com/oppermax/aoc2021/pkg/utils"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

func initFishes(in string) []int {
	inSplit := strings.Split(in, ",")
	fishes := make([]int, 9)
	for i := 0; i < len(inSplit); i++ {
		fish, err := strconv.Atoi(inSplit[i])
		if err != nil {
			log.Panic()
		}
		fishes[fish]++
	}
	return fishes
}

func addDay(fishes []int) []int {
	newFishes := fishes[0] // there are as many new fishes as there are 0 day fishes
	fishes[0] = 0 // reset 0 day fishes
	for j := 1; j < len(fishes); j++ { // we start at 1 so we can substract 1 to reduce the days
		fishes[j-1] = fishes[j] // we reduce the days per fish
	}
	fishes[6] += newFishes // newFishes are also the number of old fishes so we add them
	fishes[8] = newFishes // new fishes start out with 8 days to go
	return fishes
}

func addDays(in string, times int) int{

	fishes := initFishes(in)
	for i := 1; i <= times; i++ {
		fishes = addDay(fishes)
	}

	total := 0
	for _, c := range fishes {
		total += c
	}
	return total
}

func main()  {
	in := utils.OpenFile("input.txt")

	t1 := addDays(in[0], 80)
	t2 := addDays(in[0], 256)


	log.Infof("Part 1: after 80 days, there are %d fishes", t1)
	log.Infof("Part 2: after 256 days, there are %d fishes", t2)

}

