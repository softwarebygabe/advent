package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseInput(filepath string) (line1 []string, line2 []string) {
	results := [][]string{}
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		directionStrings := strings.Split(line, ",")
		results = append(results, directionStrings)
	}

	line1 = results[0]
	line2 = results[1]
	return
}

// Point ...
type Point struct {
	x int
	y int
}

func up(from Point, amt int) []Point {
	results := []Point{}
	for i := 1; i <= amt; i++ {
		newPoint := Point{
			x: from.x,
			y: from.y + i,
		}
		results = append(results, newPoint)
	}
	return results
}

func down(from Point, amt int) []Point {
	results := []Point{}
	for i := 1; i <= amt; i++ {
		newPoint := Point{
			x: from.x,
			y: from.y - i,
		}
		results = append(results, newPoint)
	}
	return results
}

func left(from Point, amt int) []Point {
	results := []Point{}
	for i := 1; i <= amt; i++ {
		newPoint := Point{
			x: from.x - i,
			y: from.y,
		}
		results = append(results, newPoint)
	}
	return results
}

func right(from Point, amt int) []Point {
	results := []Point{}
	for i := 1; i <= amt; i++ {
		newPoint := Point{
			x: from.x + i,
			y: from.y,
		}
		results = append(results, newPoint)
	}
	return results
}

func getPointsForLine(line []string, origin Point) []Point {
	linePoints := []Point{}
	localOrigin := origin
	for _, instruction := range line {
		split := strings.SplitAfterN(instruction, "", 2)
		direction := split[0]
		amount, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		switch direction {
		case "U":
			linePoints = append(linePoints, up(localOrigin, amount)...)
		case "D":
			linePoints = append(linePoints, down(localOrigin, amount)...)
		case "L":
			linePoints = append(linePoints, left(localOrigin, amount)...)
		case "R":
			linePoints = append(linePoints, right(localOrigin, amount)...)
		}
		localOrigin = linePoints[len(linePoints)-1]
	}
	return linePoints
}

func isEqual(a, b Point) bool {
	return a.x == b.x && a.y == b.y
}

func findIntersections(lineA, lineB []Point) []Point {
	intersections := []Point{}
	for _, pA := range lineA {
		for _, pB := range lineB {
			if isEqual(pA, pB) {
				intersections = append(intersections, pA)
			}
		}
	}
	return intersections
}

func manhattanDist(p, q Point) float64 {
	dx := math.Abs(float64(p.x - q.x))
	dy := math.Abs(float64(p.y - q.y))
	return dx + dy
}

func lowestDist(origin Point, intersectionPoints []Point) float64 {
	lowestDist := math.Inf(1)
	for _, ip := range intersectionPoints {
		dist := manhattanDist(origin, ip)
		if dist < lowestDist {
			lowestDist = dist
		}
	}
	return lowestDist
}

func distToIntersectionPoint(linePoints []Point, intersectionPoint Point) float64 {
	dist := 0
	for _, p := range linePoints {
		dist++
		if isEqual(p, intersectionPoint) {
			break
		}
	}
	return float64(dist)
}

func lowestIntersectionDist(line1, line2 []Point, intersections []Point) float64 {
	minDist := math.Inf(1)
	for _, ip := range intersections {
		line1Dist := distToIntersectionPoint(line1, ip)
		line2Dist := distToIntersectionPoint(line2, ip)
		newDist := line1Dist + line2Dist
		if newDist < minDist {
			minDist = newDist
		}
	}
	return minDist
}

func main() {

	line1, line2 := parseInput("./input.txt")

	origin := Point{
		x: 100,
		y: 100000,
	}

	line1Points := getPointsForLine(line1, origin)
	line2Points := getPointsForLine(line2, origin)

	fmt.Println(line1Points)
	fmt.Println(line2Points)

	interectionPoints := findIntersections(line1Points, line2Points)
	fmt.Println(interectionPoints)

	fmt.Println(lowestIntersectionDist(line1Points, line2Points, interectionPoints))
}
