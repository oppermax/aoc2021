package main

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"strconv"

	//utils "github.com/oppermax/aoc2021/pkg/utils"
	"os"
)

func readFile(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		num := scanner.Text()

		lines = append(lines, num)
	}
	return lines
}

func main()  {
	file, err := os.Open("input.txt")
	if err != nil {
		log.WithError(err)
		return
	}
	inputString := readFile(file)
	input, err := convertInt(inputString)
	if err != nil {
		log.WithError(err)
		return
	}
	increases := countIncreases(input)
	log.Infof("Solution: %d", increases)

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