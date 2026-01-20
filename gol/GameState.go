package gol

import (
	"fmt"
	"time"
)

func ColorReset() string { return "\033[0m" }
func Red() string        { return "\033[31m" }
func Green() string      { return "\033[32m" }
func Purple() string     { return "\033[35m" }
func Cyan() string       { return "\033[36m" }

type GameState struct {
	PreviousGrid          Grid   		// Previous generation
	NextGrid              Grid    		// New generation
	alives                [][]int 		// Alives cells coordinatees
	Size, Generation, lag int  			// Generation number, Lag frame/milliseconds
	debug                 bool    		// Set to true to display debug logs
}

func NewGameState(n, lag int) *GameState {
	return &GameState{
		PreviousGrid: *NewSeed(n),
		NextGrid:     *NewGrid(n),
		Size:         n,
		Generation:   0,
		alives:       [][]int{},
		lag:          lag,
		debug:        false,
	}
}

// --------------------------------------------

func (gs *GameState) Init(){
	gs.initAlives()
	gs.Show()
}

func (gs *GameState) initAlives(){
	alives := [][]int{}
	for i := range gs.PreviousGrid {
		for j := range gs.PreviousGrid[i] {
			cell := gs.PreviousGrid.GetCell(i, j)
			if (cell.IsAlive()){
				alives = append(alives, []int{i, j})
			}
		}
	}
	gs.SetAlives(alives)
}

// --------------------------------------------

func (gs *GameState) Update() {
	gs.updateGen()
	gs.prepareNextGen()
}

// Updates and display next grid state with all logical process
func (gs *GameState) updateGen(){
	gs.updateGenNumber()
	alives :=  [][]int{}	
	for i := range gs.alives {

		// Update cell and its adjacents
		cellCord := gs.alives[i]
		cell := gs.NextGrid.UpdateCell(cellCord[0], cellCord[1], &gs.PreviousGrid)

		if cell.IsAlive() {
			alives = append(alives, []int{cellCord[0], cellCord[1]})
		}
		
	}
	gs.SetAlives(alives)
}

// Prepare next iteration
func (gs *GameState) prepareNextGen() {
	gs.transfertPreviousToNextGrid()
	gs.sleepDelay()
}

func (gs *GameState) sleepDelay(){
	// Game speed
	delay := time.Duration(gs.GetLag()) * time.Millisecond
	time.Sleep(delay)
}

func (gs *GameState) transfertPreviousToNextGrid() {
	gs.PreviousGrid = gs.NextGrid
	gs.NextGrid = *NewGrid(gs.Size)
}

// --------------------------------------------

func (gs *GameState) Show() {
	clearShow()
	gs.showHeader()
	gs.PreviousGrid.Show()
}

func clearShow() {
	fmt.Print("\033[H\033[2J")
}

func (gs *GameState) showHeader() {
	fmt.Println(Purple(), "------------------", ColorReset())
	fmt.Println(Cyan(), "  Generation:", gs.Generation, ColorReset())
	fmt.Println(Cyan(), "  Population:", len(*gs.GetAlives()), ColorReset())
	fmt.Println(Purple(), "------------------", ColorReset())
}

// --------------------------------------------

func (gs *GameState) updateGenNumber() {
	gs.Generation++
}

func (gs *GameState) GetAlives() *[][]int {
	return &gs.alives
}

func (gs *GameState) GetLag() int {
	return gs.lag
}

func (gs *GameState) GetDebug() bool {
	return gs.debug
}

func (gs *GameState) SetAlives(alives [][]int) {
	gs.alives = alives
}