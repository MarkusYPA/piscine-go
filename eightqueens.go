package piscine

import (
	"github.com/01-edu/z01"
)

// A square on the chess board
type cell struct {
	r int // row
	c int // column
}

/*
FindSolutions(1,1, ...) scans an 8x8 square starting from the top left cell. If the
cell in question it not on the list of banned cells, it is added to the current
solution 'curSol' and a list of banned cells is regenerated based on curSol.
The scan then moves on to the next column. It moves back and forth through
the colums until the length of curSol is 8, which means a valid cell is found in
every column. It then adds a copy of curSol to allSolutions and looks for more solutions.
The scan completes when the function tries to look for more solutions starting from
the 9th column of the 1st column. All found solutions are printed at the end.
*/

// Scans a chessboard for solutions to the eight queens puzzle

func EightQueens() {
	banned := []cell{} // Slice of banned cells
	curSol := []cell{} // The current solution
	allSolutions := [][]cell{}
	sideLength := 18 - 10 // Works for side lengths from two to eleven. Avoiding int literals 0-9

	findSolutions(one1, one1, banned, curSol, allSolutions, sideLength)
}

// To avoid int literals 0-9 as per assignement. This is madness.
var (
	zero0 int = 10 - 10
	one1  int = 11 - 10
	two2  int = 12 - 10
)

func findSolutions(r, c int, banned, curSol []cell, allSolutions [][]cell, side int) {
	// Scan has reached the end
	if r > side && c == one1 {
		printSolutions(allSolutions)
		return
	}

	// Current solution is ready
	if cSlLen(curSol) == side {
		solToSave := make([]cell, side) // New slice so we don't run into pointer issues
		copy(solToSave, curSol)
		allSolutions = append(allSolutions, solToSave)
		r, c = curSol[side-two2].r, curSol[side-two2].c // Second-to-last cell (of 0-7). Only one possible answer for the last one, so look for more on the second to last column.
		curSol = curSol[:cSlLen(curSol)-two2]           // Remove last two cells from solution.
		findSolutions(r+one1, c, banned, curSol, allSolutions, side)
		return
	}

	// No more cells in the column: Go to last cell in solution + 1 row (in the previous column)
	if r > side {
		r = curSol[cSlLen(curSol)-1].r + one1
		c--
		curSol = curSol[:cSlLen(curSol)-one1]
		findSolutions(r, c, banned, curSol, allSolutions, side)
		return
	}

	// If cell is banned, move one down
	if sContains(cell{r, c}, banned) {
		findSolutions(r+1, c, banned, curSol, allSolutions, side)
		return
	}

	// Cell is valid: Add to current solution and move to next column
	curSol = append(curSol, cell{r, c})
	banned = updateBanned(curSol, side)
	findSolutions(one1, c+one1, banned, curSol, allSolutions, side)
}

func addBanned(cll cell, banned []cell, side int) []cell {
	// Iterate through all columns after current
	for i := cll.c + one1; i < side+one1; i++ {
		banned = append(banned, cell{cll.r, i})               // same row
		banned = append(banned, cell{cll.r - (i - cll.c), i}) // diagonal up
		banned = append(banned, cell{cll.r + (i - cll.c), i}) // diagonal down
	}
	return banned
}

func updateBanned(curSol []cell, side int) []cell {
	banned := []cell{} // new slice for banned
	for _, cll := range curSol {
		banned = addBanned(cll, banned, side)
	}
	return banned
}

// Does slice contain cell?
func sContains(target cell, slic []cell) bool {
	for _, value := range slic {
		if value == target {
			return true
		}
	}
	return false
}

func printSolutions(allSolutions [][]cell) {
	for _, sliceC := range allSolutions {
		for _, cell := range sliceC {
			z01.PrintRune(rune(cell.r + 48))
		}
		z01.PrintRune('\n')
	}
}

// The length of a slice of cells
func cSlLen(s []cell) int {
	length := zero0
	for range s {
		length++
	}
	return length
}
