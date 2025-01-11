package Battlefield

import (
	"errors"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
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

type Position struct{ X, Y int }

func NewBattleField(lakes int, buildings int) [][]int {
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
			if CheckNextPosition(battlefield, nextX, nextY) != 3 {
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

		if CheckNextPosition(battlefield, nextX, nextY) != 3 {
			battlefield[nextX][nextY] = WALL
			x = nextX
			y = nextY
		} else {
			break
		}
	}
}

func DrawBattleField(battlefield [][]int) (string, error) {
	var finalBattlefield string
	if len(battlefield) == 0 {
		return "", errors.New("Not correct shape of battlefield")
	}

	for i := 0; i < len(battlefield)*2+3; i++ {
		finalBattlefield = finalBattlefield + fmt.Sprintf("-")
	}
	finalBattlefield = finalBattlefield + fmt.Sprintf("\n")
	for row := 0; row < len(battlefield); row++ {
		finalBattlefield = finalBattlefield + fmt.Sprintf("| ")
		for column := 0; column < len(battlefield[row]); column++ {
			finalBattlefield = finalBattlefield + fmt.Sprintf("%d ", battlefield[row][column])
		}
		finalBattlefield = finalBattlefield + fmt.Sprintf("|\n")
	}

	for i := 0; i < len(battlefield)*2+3; i++ {
		finalBattlefield = finalBattlefield + fmt.Sprintf("-")
	}
	return finalBattlefield, nil
}

func LogBattlefield(battlefield [][]int) {
	var battleString string
	for _, row := range battlefield {
		for _, column := range row {
			battleString = battleString + fmt.Sprintf("%d", column)
		}
		battleString = battleString + "\n"
	}
	battleByte := []byte(battleString)

	err := os.WriteFile("./BattlefieldLog", battleByte, 0644)

	if err != nil {
		log.Fatalln(err)
	}

}

// Will return number of next cell, 3 if it is out of bounds.
func CheckNextPosition(battlefield [][]int, x, y int) int {
	if x < 0 || y < 0 || x >= len(battlefield) || y >= len(battlefield[0]) {
		return 3
	}
	return battlefield[y][x]
}
