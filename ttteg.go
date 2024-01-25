package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	emptyCell = " "
	playerX   = "X"
	playerO   = "O"
)

type gameBoard struct {
	board  [3][3]string
	player string
}

func playMove(gb gameBoard, row, col int) gameBoard {
	gb.board[row-1][col-1] = gb.player
	return gb
}

func ttteg() {
	gb := initializeGame()

	for {
		printBoard(gb)

		if gb.player == playerX {
			// Human player's turn
			fmt.Printf("Player %s, enter your move (row and column, e.g., 1 2): ", gb.player)
			var row, col int
			fmt.Scan(&row, &col)

			if isValidMove(gb, row, col) {
				gb = playMove(gb, row, col)
				if checkWin(gb) {
					printBoard(gb)
					fmt.Printf("Player %s wins!\n", gb.player)
					break
				} else if checkDraw(gb) {
					printBoard(gb)
					fmt.Println("It's a draw!")
					break
				}

				switchPlayer(&gb)
			} else {
				fmt.Println("Invalid move. Try again.")
			}
		} else {
			// Computer player's turn
			fmt.Println("Computer's turn...")
			time.Sleep(time.Duration(float64(0.5) * float64(time.Second))) // Simulate a delay for computer's move

			gb = computerPlay(gb)
			if checkWin(gb) {
				printBoard(gb)
				fmt.Printf("Player %s wins!\n", gb.player)
				break
			} else if checkDraw(gb) {
				printBoard(gb)
				fmt.Println("It's a draw!")
				break
			}

			switchPlayer(&gb)
		}
	}
}

func computerPlay(gb gameBoard) gameBoard {
	// Randomly choose an empty cell for the computer's move
	for {
		var row, col int

		if checksetup(gb) == false {
			row = rand.Intn(3) + 1
			col = rand.Intn(3) + 1
		} else {
			gb = pcblock(gb, &row, &col)
		}

		fmt.Printf("Computer plays at %d %d\n", row, col)
		if isValidMove(gb, row, col) {
			fmt.Printf("Computer plays at %d %d\n", row, col)
			return playMove(gb, row, col)
		}
	}
}

// func pcblock(gb gameBoard, row, col *int) gameBoard {
// 	if checksetup(gb) == true {
// 		if setupcol(gb) == true {
// 			for i := 0; i < 3; i++ {
// 				if gb.board[0][i] != emptyCell && gb.board[0][i] == gb.board[1][i] && gb.board[2][i] == emptyCell {
// 					*row = 3
// 					*col = i + 1
// 					gb.board[*row-1][*col-1] = gb.player
// 					return gb
// 				} else if gb.board[1][i] != emptyCell && gb.board[1][i] == gb.board[2][i] && gb.board[0][i] == emptyCell {
// 					*row = 1
// 					*col = i + 1
// 					gb.board[*row-1][*col-1] = gb.player
// 					return gb
// 				} else if gb.board[2][i] != emptyCell && gb.board[2][i] == gb.board[0][i] && gb.board[1][i] == emptyCell {
// 					*row = 2
// 					*col = i + 1
// 					gb.board[*row-1][*col-1] = gb.player
// 					return gb
// 				}
// 			}
// 		} else if setuprow(gb) == true {
// 			for i, currRow := range gb.board {
// 				if currRow[0] != emptyCell && currRow[0] == currRow[1] && currRow[2] == emptyCell {
// 					*row = i + 1
// 					*col = 3
// 					gb.board[*row-1][*col-1] = gb.player
// 					return gb
// 				} else if currRow[1] != emptyCell && currRow[1] == currRow[2] && currRow[0] == emptyCell {
// 					*row = i + 1
// 					*col = 1
// 					gb.board[*row-1][*col-1] = gb.player
// 					return gb
// 				} else if currRow[2] != emptyCell && currRow[2] == currRow[0] && currRow[1] == emptyCell {
// 					*row = i + 1
// 					*col = 2
// 					gb.board[*row-1][*col-1] = gb.player
// 					return gb
// 				}
// 			}
// 		} else if setupdiagon(gb) == true {
// 			if gb.board[0][0] != emptyCell && gb.board[0][0] == gb.board[1][1] && gb.board[2][2] == emptyCell {
// 				*row = 3
// 				*col = 3
// 				gb.board[*row-1][*col-1] = gb.player
// 				return gb
// 			} else if gb.board[1][1] != emptyCell && gb.board[1][1] == gb.board[2][2] && gb.board[0][0] == emptyCell {
// 				*row = 1
// 				*col = 1
// 				gb.board[*row-1][*col-1] = gb.player
// 				return gb
// 			} else if (gb.board[2][2] != emptyCell && gb.board[2][2] == gb.board[0][0]) && (gb.board[1][1] == emptyCell) {
// 				*row = 2
// 				*col = 2
// 				gb.board[*row-1][*col-1] = gb.player
// 				return gb
// 			} else if gb.board[0][2] != emptyCell && gb.board[0][2] == gb.board[1][1] && gb.board[2][0] == emptyCell {
// 				*row = 3
// 				*col = 1
// 				gb.board[*row-1][*col-1] = gb.player
// 				return gb
// 			} else if gb.board[1][1] != emptyCell && gb.board[1][1] == gb.board[2][0] && gb.board[0][2] == emptyCell {
// 				*row = 1
// 				*col = 3
// 				gb.board[*row-1][*col-1] = gb.player
// 				return gb
// 			}
// 		}
// 	}
// 	// If no blocking move is found, make a random move
// 	*row = rand.Intn(3) + 1
// 	*col = rand.Intn(3) + 1

// 	if isValidMove(gb, *row, *col) {
// 		gb.board[*row-1][*col-1] = gb.player
// 		return gb
// 	}
// 	return gb
// }

func pcblock(gb gameBoard, row, col *int) gameBoard {
	fmt.Println("Inside pcblock loop")
	if checksetup(gb) {
		// Check for blocking move in rows
		for i, currRow := range gb.board {
			if currRow[0] == gb.player && currRow[0] == currRow[1] && gb.board[i][2] == emptyCell {
				*row = i + 1
				*col = 3
				// gb.board[*row-1][*col-1] = gb.player
				return gb
			} else if currRow[1] == gb.player && currRow[1] == currRow[2] && gb.board[i][0] == emptyCell {
				*row = i + 1
				*col = 1
				// gb.board[*row-1][*col-1] = gb.player
				return gb
			} else if currRow[2] == gb.player && currRow[2] == currRow[0] && gb.board[i][1] == emptyCell {
				*row = i + 1
				*col = 2
				// gb.board[*row-1][*col-1] = gb.player
				return gb
			}
		}
		fmt.Println("No blocking rows found")
		printBoard(gb)

		// Check for blocking move in columns
		for i := 0; i < 3; i++ {
			if gb.board[0][i] == gb.player && gb.board[0][i] == gb.board[1][i] && gb.board[2][i] == emptyCell {
				*row = 3
				*col = i + 1
				// gb.board[*row-1][*col-1] = gb.player
				return gb
			} else if gb.board[1][i] == gb.player && gb.board[1][i] == gb.board[2][i] && gb.board[0][i] == emptyCell {
				*row = 1
				*col = i + 1
				// gb.board[*row-1][*col-1] = gb.player
				return gb
			} else if gb.board[2][i] == gb.player && gb.board[2][i] == gb.board[0][i] && gb.board[1][i] == emptyCell {
				*row = 2
				*col = i + 1
				// gb.board[*row-1][*col-1] = gb.player
				return gb
			}
		}
		fmt.Println("No blocking columns found")
		printBoard(gb)

		// Check for blocking move in diagonals
		if gb.board[0][0] == gb.player && gb.board[0][0] == gb.board[1][1] && gb.board[2][2] == emptyCell {
			*row = 3
			*col = 3
			// gb.board[*row-1][*col-1] = gb.player
			return gb
		} else if gb.board[1][1] == gb.player && gb.board[1][1] == gb.board[2][2] && gb.board[0][0] == emptyCell {
			*row = 1
			*col = 1
			// gb.board[*row-1][*col-1] = gb.player
			return gb
		} else if gb.board[2][2] == gb.player && gb.board[2][2] == gb.board[0][0] && gb.board[1][1] == emptyCell {
			*row = 2
			*col = 2
			// gb.board[*row-1][*col-1] = gb.player
			return gb
		} else if gb.board[0][2] == gb.player && gb.board[0][2] == gb.board[1][1] && gb.board[2][0] == emptyCell {
			*row = 3
			*col = 1
			// gb.board[*row-1][*col-1] = gb.player
			return gb
		} else if gb.board[1][1] == gb.player && gb.board[1][1] == gb.board[2][0] && gb.board[0][2] == emptyCell {
			*row = 1
			*col = 3
			// gb.board[*row-1][*col-1] = gb.player
			return gb
		}
	} else {
		// If no blocking move is found, make a random move
		*row = rand.Intn(3) + 1
		*col = rand.Intn(3) + 1
	}
	if gb.board[*row][*col] == emptyCell {
		gb.board[*row-1][*col-1] = gb.player
		return gb
	}
	return gb
}

// Add print statements in checksetup
// func checksetup(gb gameBoard) bool {
// 	fmt.Println("Checking setup...")
// 	tempBoard := gb
// 	result := setupcol(tempBoard) || setupdiagon(tempBoard) || setuprow(tempBoard)
// 	fmt.Println("Result of checksetup:", result)
// 	return result
// }

func checksetup(gb gameBoard) bool {
	checked := false

	if !checked && setupcol(gb) {
		fmt.Println("Checking setupcol...")
		checked = true
	}
	if !checked && setupdiagon(gb) {
		fmt.Println("Checking setupdiagon...")
		checked = true
	}
	if !checked && setuprow(gb) {
		fmt.Println("Checking setuprow...")
		checked = true
	}

	return checked
}

func setuprow(gb gameBoard) bool {
	fmt.Println("Checking setuprow...")
	for _, row := range gb.board {
		if (row[0] != emptyCell && row[0] == row[1]) || (row[1] != emptyCell && row[1] == row[2]) || (row[2] != emptyCell && row[2] == row[0]) {
			return true
		}
	}
	return false
}

// func setupcol(gb gameBoard) bool {
// 	fmt.Println("Checking setupcol...")
// 	for i := 0; i < 3; i++ {
// 		if (gb.board[0][i] != emptyCell && gb.board[0][i] == gb.board[1][i]) || (gb.board[1][i] != emptyCell && gb.board[1][i] == gb.board[2][i]) || (gb.board[2][i] != emptyCell && gb.board[2][i] == gb.board[0][i]) {
// 			return true
// 		}
// 	}
// 	return false
// }

func setupcol(gb gameBoard) bool {
	fmt.Println("Checking setupcol...")
	for i := 0; i < 3; i++ {
		if gb.board[0][i] != emptyCell && gb.board[0][i] == gb.board[1][i] && gb.board[2][i] == emptyCell {
			return true
		} else if gb.board[1][i] != emptyCell && gb.board[1][i] == gb.board[2][i] && gb.board[0][i] == emptyCell {
			return true
		} else if gb.board[2][i] != emptyCell && gb.board[2][i] == gb.board[0][i] && gb.board[1][i] == emptyCell {
			return true
		}
	}
	return false
}

func setupdiagon(gb gameBoard) bool {
	fmt.Println("Checking setupdiagon...")
	if (gb.board[0][0] != emptyCell && gb.board[0][0] == gb.board[1][1]) || (gb.board[1][1] != emptyCell && gb.board[1][1] == gb.board[2][2]) || (gb.board[2][2] != emptyCell && gb.board[2][2] == gb.board[0][0]) {
		return true
	}
	if (gb.board[0][2] != emptyCell && gb.board[0][2] == gb.board[1][1]) || (gb.board[1][1] != emptyCell && gb.board[1][1] == gb.board[2][0]) || (gb.board[2][0] != emptyCell && gb.board[2][0] == gb.board[0][2]) {
		return true
	}
	return false
}

func initializeGame() gameBoard {
	// Initialize the board with empty cells
	var board [3][3]string
	for i := range board {
		for j := range board[i] {
			board[i][j] = emptyCell
		}
	}

	// Player X starts the game
	player := playerX
	return gameBoard{
		board:  board,
		player: player,
	}
}

func printBoard(gb gameBoard) {
	fmt.Println("  1 2 3")
	for i, row := range gb.board {
		fmt.Printf("%d ", i+1)
		for _, cell := range row {
			fmt.Printf("%s ", cell)
		}
		fmt.Println()
	}
	fmt.Println("Current player:", gb.player)
	fmt.Println()
}

func isValidMove(gb gameBoard, row, col int) bool {
	return row >= 1 && row <= 3 && col >= 1 && col <= 3 && gb.board[row-1][col-1] == emptyCell
}

func checkWin(gb gameBoard) bool {
	// Check rows, columns, and diagonals for a win
	return checkRows(gb) || checkColumns(gb) || checkDiagonals(gb)
}

func checkRows(gb gameBoard) bool {
	for _, row := range gb.board {
		if row[0] != emptyCell && row[0] == row[1] && row[1] == row[2] {
			return true
		}
	}
	return false
}

func checkColumns(gb gameBoard) bool {
	for i := 0; i < 3; i++ {
		if gb.board[0][i] != emptyCell && gb.board[0][i] == gb.board[1][i] && gb.board[1][i] == gb.board[2][i] {
			return true
		}
	}
	return false
}

func checkDiagonals(gb gameBoard) bool {
	if gb.board[0][0] != emptyCell && gb.board[0][0] == gb.board[1][1] && gb.board[1][1] == gb.board[2][2] {
		return true
	}
	if gb.board[0][2] != emptyCell && gb.board[0][2] == gb.board[1][1] && gb.board[1][1] == gb.board[2][0] {
		return true
	}
	return false
}

func checkDraw(gb gameBoard) bool {
	// Check if the board is full
	for _, row := range gb.board {
		for _, cell := range row {
			if cell == emptyCell {
				return false
			}
		}
	}
	return true
}

func switchPlayer(gb *gameBoard) {
	if gb.player == playerX {
		gb.player = playerO
	} else {
		gb.player = playerX
	}
}
