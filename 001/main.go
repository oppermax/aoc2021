package main

import (
	"github.com/oppermax/aoc2021/pkg/utils"
	log "github.com/sirupsen/logrus"
	"strconv"
)

func main()  {

	inputString := utils.OpenFile("input.txt")
	input, err := convertInt(inputString)
	if err != nil {
		log.WithError(err)
		return
	}
	increases1 := countIncreases(input)
	log.Infof("Solution 1: %d", increases1)
	increases2 := countIncreases(buildSets(input))
	log.Infof("Solution 2: %d", increases2)

}

func countIncreases(in []int) int {
	var increases, previous int
	log.Info(in)
	for i, line := range in {
		if i == 0 {
			previous = line
			continue
		}
		if line > previous {
			log.Infof("previous: %d current: %d", previous, line)
			increases++
		}
		previous = line
	}


	return increases
}

func convertInt(in []string) ([]int, error){
	out := make([]int, len(in))
	for i, line := range in {
		intg, err := strconv.Atoi(line)
		if err != nil {
			return out, err
		}
		out[i] = intg
	}
	return out, nil
}

func buildSets(in []int) []int {
	subset := in
	sets := make([]int, len(in))
	for i := range in {
		if len(subset) >= 3 {
			sets[i] = getSum(subset[:3])
		}
		subset = in[i:]
	}
	return sets
}

func getSum(in []int) int {
	var out int
	for _, d := range in {
		out += d
	}
	return out
}