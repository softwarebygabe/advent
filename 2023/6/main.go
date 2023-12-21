package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/softwarebygabe/advent/pkg/util"
)

type race struct {
	time, distance int
}

func newRace(t, d int) race {
	return race{time: t, distance: d}
}

func (r race) possibleWinScenarios() int {
	wins := 0
	calcDistance := func(t0, t1 int) int {
		return (t0 * t1) - (t0 * t0)
	}
	for t := 0; t < r.time; t++ {
		dist := calcDistance(t, r.time)
		if dist > r.distance {
			wins++
		}
	}
	return wins
}

func trimAllSpace(s string) string {
	res := []rune{}
	for _, r := range s {
		if !unicode.IsSpace(r) {
			res = append(res, r)
		}
	}
	return string(res)
}

func cleanEach(s []string) []string {
	res := []string{}
	for _, si := range s {
		if si != "" {
			res = append(res, trimAllSpace(si))
		}
	}
	return res
}

func part1(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lines := util.ReaderToStrings(f)
	times := []int{}
	distances := []int{}
	if len(lines) != 2 {
		panic("expected only 2 lines")
	}
	for idx, line := range lines {
		if idx == 0 {
			timesRaw := strings.TrimPrefix(line, "Time: ")
			timesRawList := cleanEach(strings.Split(timesRaw, " "))
			for _, timeRaw := range timesRawList {
				n, err := strconv.Atoi(timeRaw)
				if err != nil {
					panic(err)
				}
				times = append(times, n)
			}
		} else {
			distancesRaw := strings.TrimPrefix(line, "Distance: ")
			distancesRawList := cleanEach(strings.Split(distancesRaw, " "))
			for _, distanceRaw := range distancesRawList {
				n, err := strconv.Atoi(distanceRaw)
				if err != nil {
					panic(err)
				}
				distances = append(distances, n)
			}
		}
	}
	races := []race{}
	if len(times) != len(distances) {
		panic("expected line lengths to match")
	}
	for i := 0; i < len(times); i++ {
		races = append(races, newRace(times[i], distances[i]))
	}
	result := 0
	for _, race := range races {
		wins := race.possibleWinScenarios()
		if result == 0 {
			result = wins
		} else {
			result *= wins
		}
	}
	fmt.Println("result:", result)
}

func part2(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lines := util.ReaderToStrings(f)
	time := ""
	distance := ""
	if len(lines) != 2 {
		panic("expected only 2 lines")
	}
	for idx, line := range lines {
		if idx == 0 {
			timesRaw := strings.TrimPrefix(line, "Time: ")
			timesRawList := cleanEach(strings.Split(timesRaw, " "))
			for _, timeRaw := range timesRawList {
				time += timeRaw
			}
		} else {
			distancesRaw := strings.TrimPrefix(line, "Distance: ")
			distancesRawList := cleanEach(strings.Split(distancesRaw, " "))
			for _, distanceRaw := range distancesRawList {
				distance += distanceRaw
			}
		}
	}
	t, err := strconv.Atoi(time)
	if err != nil {
		panic(err)
	}
	d, err := strconv.Atoi(distance)
	if err != nil {
		panic(err)
	}
	races := []race{
		newRace(t, d),
	}
	result := 0
	for _, race := range races {
		wins := race.possibleWinScenarios()
		if result == 0 {
			result = wins
		} else {
			result *= wins
		}
	}
	fmt.Println("result:", result)
}

func main() {
	part2("input.txt")
}
