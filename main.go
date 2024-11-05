package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

const (
	maxLakeSize   = 8
	maxWallsCells = 5
)

const (
	LAND = iota
	WATER
	WALL
	OUTBOUNDS
)

func buildMap(lakes int, buildings int) [][]int {
	array := make([][]int, 10)

	for i := range array {
		array[i] = make([]int, 10)
	}

	for i := 0; i < lakes; i++ {
		size := rand.IntN(maxLakeSize)
		buildLake(array, size)
	}
	for i := 0; i < buildings; i++ {
		buildWalls(array)
	}
	return array
}

func buildLake(battlefield [][]int, size int) {
	x := rand.IntN(len(battlefield))
	y := rand.IntN(len(battlefield[0]))

	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	for size > 0 {
		battlefield[x][y] = WATER
		indexDir := rand.IntN(3)

		for {
			nextDir := directions[indexDir]
			nextX, nextY := x+nextDir[0], y+nextDir[1]
			if nextX < len(battlefield) && nextX >= 0 && nextY < len(battlefield[0]) && nextY >= 0 {
				x = nextX
				y = nextY
				break
			} else if indexDir >= 3 {
				indexDir = 0
			} else {
				indexDir++
			}
		}
		size--

	}
}

func buildWalls(battlefield [][]int) {
	x := rand.IntN(len(battlefield))
	y := rand.IntN(len(battlefield[0]))

	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	battlefield[x][y] = WALL
	indexDir := rand.IntN(3)

	for i := 1; i < maxWallsCells; i++ {
		nextX := x + directions[indexDir][0]
		nextY := y + directions[indexDir][1]

		if nextX >= 0 && nextX < len(battlefield) && nextY >= 0 && nextY < len(battlefield[0]) {
			battlefield[nextX][nextY] = WALL
			x = nextX
			y = nextY
		} else {
			break
		}
	}
}

func drawMap(battlefield [][]int) error {
	if len(battlefield) == 0 {
		return errors.New("Not correct shape of battlefield")
	}

	for i := 0; i < len(battlefield)*2+3; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("\n")
	for row := 0; row < len(battlefield); row++ {
		fmt.Printf("| ")
		for column := 0; column < len(battlefield[row]); column++ {
			fmt.Printf("%d ", battlefield[row][column])
		}
		fmt.Printf("|\n")
	}

	for i := 0; i < len(battlefield)*2+3; i++ {
		fmt.Printf("-")
	}
	return nil
}

func checkNextPosition(battlefield [][]int, nextX int, nextY int) int {
	// Will return number of next cell, 3 if it is out of bounds.
	if nextX >= 0 && nextX < len(battlefield) && nextY >= 0 && nextY < len(battlefield[0]) {
		return battlefield[nextX][nextY]
	} else {
		return OUTBOUNDS
	}
}

func main() {
	field := buildMap(3, 2)
	drawMap(field)

}
