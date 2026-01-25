package gol

import (
	"fmt"
	"math/rand"
)

type Grid [][]Cell

func NewGrid(n int) *Grid {
	m := make(Grid, n)
	for i := range n {
		m[i] = make([]Cell, n)
		for j := range m[i] {
			m[i][j] = *NewCell(i, j)
		}
	}
	return &m
}

// Creates a randomly generated grid
func NewSeed(n int) *Grid {
	m := *NewGrid(n)
	for range n * 5 {
		m[rand.Intn(n)][rand.Intn(n)].SetAlive(true)
	}
	return &m
}

// --------------------------------------------

// Updates a given cell
func (newGrid *Grid) UpdateCell(i, j int, oldGrid *Grid) *Cell {
	oldGrid.updateCellAdjs(i, j)

	cell := oldGrid.GetCell(i, j)
	cell.UpdateState()

	newGrid.SetCell(i, j, *cell)
	return cell
}

// Updates all adjacents of a given cell
func (oldGrid *Grid) updateCellAdjs(i, j int) {
	adjs := make([]Cell, 0, 8)
	for _, position := range getAdjacentsPos() {
		if oldGrid.isValidPosition(i+position[0], j+position[1]) {
			if adjCell := oldGrid.GetCell(i+position[0], j+position[1]); adjCell.IsAlive() {
				adjs = append(adjs, *adjCell)
			}
		}
	}
	oldGrid.GetCell(i, j).SetAdjacents(adjs)
}

// --------------------------------------------

func (m Grid) Show() {
	for i := range m {
		for j := range m[i] {
			displayCell(m.GetCell(i, j))
		}
		fmt.Println(" ")
	}
}

func displayCell(cell *Cell) {
	fmt.Print(cell.ToString(), " ")
}

func (m Grid) GetCell(i, j int) *Cell {
	return &m[i][j]
}

func (m Grid) SetCell(i, j int, c Cell) {
	m[i][j] = c
}

// All possible adjacent position
func getAdjacentsPos() [8][2]int {
	return [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1}, // upper level
		{0, -1}, {0, 1}, // left and right level
		{1, -1}, {1, 0}, {1, 1}} // bottom level
}

func (m Grid) isValidPosition(i, j int) bool {
	return i >= 0 && i < len(m) && j >= 0 && j < len(m[0])
}
