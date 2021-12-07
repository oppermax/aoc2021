package main

import (
	"github.com/oppermax/aoc2021/pkg/utils"
	log "github.com/sirupsen/logrus"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	in := utils.OpenFile("input.txt")
	submarines := getSubs(in[0])

	pt1, pt2 := determineLowestFuelCost(submarines)
	log.Infof("Solution for part 1: %d", pt1)
	log.Infof("Solution for part 2: %d", pt2)
}

func getSubs(in string) []int {
	split := strings.Split(in, ",")
	submarines := make([]int, len(split))
	for i, str := range split {
		sub, err := strconv.Atoi(str)
		if err != nil {
			log.WithError(err).Panic()
		}
		submarines[i] = sub
	}
	return submarines
}

func alignAtSimple(level int, in []int) int {
	out := 0
	for _, sub := range in {
		out += int(math.Abs(float64(sub - level)))
	}
	return out
}


func alignAtComplex(level int, in []int) int {
	var out int
	for _, sub := range in {
		if sub < level {
			out += addComplexFuel(level - sub)
		} else if sub > level {
			out += addComplexFuel(sub - level)
		}
	}
	return out
}

func addComplexFuel(dif int) int {
	out := 0
	for i := 1; i <= dif ; i++ {
		out += i
	}
	return out
}

func determineLowestFuelCost(in []int) (pt1 int, pt2 int) {
	sort.Ints(in)
	costs1 := []int{}
	costs2 := []int{}
	for i := in[0]; i < in[len(in)-1]; i++ {
		costs1 = append(costs1, alignAtSimple(i, in))
		costs2 = append(costs2, alignAtComplex(i, in))
	}

	sort.Ints(costs1)
	sort.Ints(costs2)

	return costs1[0], costs2[0]
}
