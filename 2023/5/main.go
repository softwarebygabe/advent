package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/softwarebygabe/advent/pkg/util"
)

type mapRange struct {
	srcMins, destMins, l []int
}

type seedRange struct {
	min, l int
}

type almanac struct {
	seedToSoil   *mapRange
	soilToFert   *mapRange
	fertToWater  *mapRange
	waterToLight *mapRange
	lightToTemp  *mapRange
	tempToHum    *mapRange
	humToLoc     *mapRange
}

func newAlmanac() almanac {
	return almanac{
		seedToSoil:   &mapRange{},
		soilToFert:   &mapRange{},
		fertToWater:  &mapRange{},
		waterToLight: &mapRange{},
		lightToTemp:  &mapRange{},
		tempToHum:    &mapRange{},
		humToLoc:     &mapRange{},
	}
}

func addRangeEntries(m *mapRange, src, dest, l int) {
	m.srcMins = append(m.srcMins, src)
	m.destMins = append(m.destMins, dest)
	m.l = append(m.l, l)
}

func getFromMapRange(m *mapRange, k int) int {
	// fmt.Printf("k: %d %+v\n", k, m)
	for i := 0; i < len(m.srcMins); i++ {
		if m.inRange(i, k) {
			delta := k - m.srcMins[i]
			return m.destMins[i] + delta
		}
	}
	return k
}

func (m *mapRange) inRange(idx, k int) bool {
	srcMin := m.srcMins[idx]
	if k < srcMin || (srcMin+m.l[idx]) < k {
		// if k is outside range
		// fmt.Println("outside range!")
		return false
	}
	return true
}

func (alm almanac) seedToLoc(s int) int {
	// fmt.Printf("Seed %d, ", s)
	soil := getFromMapRange(alm.seedToSoil, s)
	// fmt.Printf("soil %d, ", soil)
	fert := getFromMapRange(alm.soilToFert, soil)
	// fmt.Printf("fert %d, ", fert)
	water := getFromMapRange(alm.fertToWater, fert)
	// fmt.Printf("water %d, ", water)
	light := getFromMapRange(alm.waterToLight, water)
	// fmt.Printf("light %d, ", light)
	temp := getFromMapRange(alm.lightToTemp, light)
	// fmt.Printf("temp %d, ", temp)
	hum := getFromMapRange(alm.tempToHum, temp)
	// fmt.Printf("hum %d, ", hum)
	loc := getFromMapRange(alm.humToLoc, hum)
	// fmt.Printf("loc %d\n", loc)
	return loc
}

func part2(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lines := util.ReaderToStrings(f)

	alm := newAlmanac()
	seedRanges := []seedRange{}
	curMap := alm.seedToSoil
	for idx, line := range lines {
		if idx == 0 {
			// first line is seeds
			line = strings.TrimPrefix(line, "seeds: ")
			numsRaw := strings.Split(line, " ")
			for i := 0; i < len(numsRaw)-1; i += 2 {
				n1, err := strconv.Atoi(numsRaw[i])
				if err != nil {
					panic(err)
				}
				n2, err := strconv.Atoi(numsRaw[i+1])
				if err != nil {
					panic(err)
				}
				seedRanges = append(seedRanges, seedRange{min: n1, l: n2})
			}
			continue
		}
		// else let's see if we are in a section or not
		if line == "" {
			continue // skip new lines
		}
		if strings.Contains(line, "map:") {
			// we've reached the start of a new section
			switch strings.Split(line, " ")[0] {
			case "seed-to-soil":
				curMap = alm.seedToSoil
			case "soil-to-fertilizer":
				curMap = alm.soilToFert
			case "fertilizer-to-water":
				curMap = alm.fertToWater
			case "water-to-light":
				curMap = alm.waterToLight
			case "light-to-temperature":
				curMap = alm.lightToTemp
			case "temperature-to-humidity":
				curMap = alm.tempToHum
			case "humidity-to-location":
				curMap = alm.humToLoc
			default:
				panic("map unrecognized")
			}
			continue
		}
		// parse nums
		numsRaw := strings.Split(line, " ")
		vals := []int{}
		for _, nRaw := range numsRaw {
			n, err := strconv.Atoi(nRaw)
			if err != nil {
				panic(err)
			}
			vals = append(vals, n)
		}
		if len(vals) != 3 {
			panic("not enough vals")
		}
		// add vals to almanac
		addRangeEntries(curMap, vals[1], vals[0], vals[2])
	}
	miner := util.NewMiner()
	s := spinner.New(
		spinner.CharSets[70],
		100*time.Millisecond,
		spinner.WithFinalMSG(""),
	)
	s.Prefix = "hold tight! calculating the geometric center of the universe... "
	s.Start()
	defer s.Stop()
	for _, sr := range seedRanges {
		fmt.Printf("working on seed range: %+v\n", sr)
		for s := sr.min; s < sr.min+sr.l; s++ {
			loc := alm.seedToLoc(s)
			// fmt.Println("seed", s, "loc", loc)
			miner.Add(loc)
		}
		fmt.Println("current min:", miner.Get())
	}
	fmt.Println("min loc:", miner.Get())
}

func part1(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lines := util.ReaderToStrings(f)

	alm := newAlmanac()
	seeds := []int{}
	curMap := alm.seedToSoil
	for idx, line := range lines {
		if idx == 0 {
			// first line is seeds
			line = strings.TrimPrefix(line, "seeds: ")
			numsRaw := strings.Split(line, " ")
			for _, nRaw := range numsRaw {
				n, err := strconv.Atoi(nRaw)
				if err != nil {
					panic(err)
				}
				seeds = append(seeds, n)
			}
			continue
		}
		// else let's see if we are in a section or not
		if line == "" {
			continue // skip new lines
		}
		if strings.Contains(line, "map:") {
			// we've reached the start of a new section
			switch strings.Split(line, " ")[0] {
			case "seed-to-soil":
				curMap = alm.seedToSoil
			case "soil-to-fertilizer":
				curMap = alm.soilToFert
			case "fertilizer-to-water":
				curMap = alm.fertToWater
			case "water-to-light":
				curMap = alm.waterToLight
			case "light-to-temperature":
				curMap = alm.lightToTemp
			case "temperature-to-humidity":
				curMap = alm.tempToHum
			case "humidity-to-location":
				curMap = alm.humToLoc
			default:
				panic("map unrecognized")
			}
			continue
		}
		// parse nums
		numsRaw := strings.Split(line, " ")
		vals := []int{}
		for _, nRaw := range numsRaw {
			n, err := strconv.Atoi(nRaw)
			if err != nil {
				panic(err)
			}
			vals = append(vals, n)
		}
		if len(vals) != 3 {
			panic("not enough vals")
		}
		// add vals to almanac
		addRangeEntries(curMap, vals[1], vals[0], vals[2])
	}
	miner := util.NewMiner()
	for _, s := range seeds {
		loc := alm.seedToLoc(s)
		fmt.Println("seed", s, "loc", loc)
		miner.Add(loc)
	}
	fmt.Println("min loc:", miner.Get())
}

func main() {
	part2("input.txt")
}
