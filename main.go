package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	gridSize, err := ParseIntInput("Provide Grid Size (default 50): ")
	if err != nil {
		gridSize = 40
	}
	userFrameRate, err := ParseIntInput("Provide Frame Rate (default 1000 ms): ")
	if err != nil {
		userFrameRate = 1000
	}

	quit := false
	var grid = make([][]bool, gridSize)
	var newGrid = make([][]bool, gridSize)
	for i := 0; i < gridSize; i++ {
		grid[i] = make([]bool, gridSize)
		newGrid[i] = make([]bool, gridSize)
		for j := 0; j < gridSize; j++ {
			// cell_status := 0
			cell_status := rand.Intn(2)
			if cell_status == 1 {
				grid[i][j] = true
				newGrid[i][j] = true
			} else {
				grid[i][j] = false
				newGrid[i][j] = false
			}
		}
	}

	// grid[0][1] = true
	// grid[1][2] = true
	// grid[2][0] = true
	// grid[2][1] = true
	// grid[2][2] = true

	// grid[3][3] = true
	// grid[3][4] = true
	// grid[3][5] = true

	// Main Loop
	for !quit {
		// Controls Framerate
		time.Sleep(time.Duration(userFrameRate) * time.Millisecond)
		// Screen Update
		// fmt.Printf("\033[0;0H")
		fmt.Println("Game Of Life")
		for i := 0; i < gridSize; i++ {
			var line string
			for j := 0; j < gridSize; j++ {
				if grid[i][j] {
					line += "██"
				} else {
					line += "  "
				}
			}
			fmt.Print(line + "\n")
		}

		// Calculate new generation
		for i := 0; i < gridSize; i++ {
			for j := 0; j < gridSize; j++ {
				newGrid[i][j] = UpdateCell(j, i, grid, gridSize)
			}
		}
		// Update Grid
		for i := 0; i < gridSize; i++ {
			for j := 0; j < gridSize; j++ {
				grid[i][j] = newGrid[i][j]
			}
		}
	}
}

func ParseIntInput(prompt string) (int, error) {
	fmt.Println(prompt)
	var input int
	_, err := fmt.Scanln("%d", &input)
	return input, err
}

func UpdateCell(cellX int, cellY int, grid [][]bool, gridSize int) bool {
	numberOfNeighbors := 0
	hasLeftNeighbors := cellX != 0
	hasRightNeighbors := cellX != gridSize-1
	hasTopNeighbors := cellY != 0
	hasBottomNeighbors := cellY != gridSize-1

	// Row with cell
	if hasLeftNeighbors && grid[cellY][cellX-1] {
		numberOfNeighbors += 1
	}
	if hasRightNeighbors && grid[cellY][cellX+1] {
		numberOfNeighbors += 1
	}
	// Row above cell
	if hasTopNeighbors {
		if hasLeftNeighbors && grid[cellY-1][cellX-1] {
			numberOfNeighbors += 1
		}
		if hasRightNeighbors && grid[cellY-1][cellX+1] {
			numberOfNeighbors += 1
		}
		if grid[cellY-1][cellX] {
			numberOfNeighbors += 1
		}
	}
	// Row below cell
	if hasBottomNeighbors {
		if hasLeftNeighbors && grid[cellY+1][cellX-1] {
			numberOfNeighbors += 1
		}
		if hasRightNeighbors && grid[cellY+1][cellX+1] {
			numberOfNeighbors += 1
		}
		if grid[cellY+1][cellX] {
			numberOfNeighbors += 1
		}
	}
	if grid[cellY][cellX] && (numberOfNeighbors == 2 || numberOfNeighbors == 3) {
		return true
	} else if !grid[cellY][cellX] && numberOfNeighbors == 3 {
		return true
	} else {
		return false
	}
}
