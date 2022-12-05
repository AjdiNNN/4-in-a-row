package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func getNewNumber(minX int, minY int) (int, int) {
	var new = ""
	re := regexp.MustCompile(`(\d+)x(\d+)`)
	fmt.Println("Type board dimensions rows must not be less than " + strconv.FormatInt(int64(minX), 10) + " and columns muss not be less than " + strconv.FormatInt(int64(minY), 10))
	fmt.Println("Type in format (rows)x(columns) ex. 6x7 and values of rows and columns must not differ by more than 2")
	for {
		fmt.Scanln(&new)
		if re.MatchString(new) {
			value := re.FindAllStringSubmatch(new, 2)[0]
			x, error1 := strconv.Atoi(value[1])
			if error1 != nil {
				fmt.Println("Error during conversion")
				return 6, 7
			}
			y, error2 := strconv.Atoi(value[2])
			if error2 != nil {
				fmt.Println("Error during conversion")
				return 6, 7
			}
			if x >= 6 {
				if y >= 7 {
					if math.Abs(float64(x-y)) <= 2 {
						return x, y
					} else {
						fmt.Println("Rows and columns must not differ by more than 2")
						continue
					}
				} else {
					fmt.Println(value[2] + " is not greater than or equal " + strconv.FormatInt(int64(minY), 10))
					continue
				}
			} else {
				fmt.Println(value[1] + " is not greater than or equal " + strconv.FormatInt(int64(minX), 10))
				continue
			}
		} else {
			fmt.Println(new + " is invalid format please try with (rows)x(columns)")
		}
	}
}
func printInitBoard(x int, y int) {
	for i := 1; i <= y; i++ {
		fmt.Print(" " + strconv.FormatInt(int64(i), 10) + " ")
	}
	fmt.Println()
	for i := 1; i <= x; i++ {
		fmt.Println(strings.Repeat("{ }", y))
	}
}
func checkIfDraw(board [][]uint8) bool {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 0 {
				return false
			}
		}
	}
	return true
}
func checkIfWon(board [][]uint8, player uint8) bool {
	for c := 0; c < (len(board[0]) - 3); c++ {
		for r := 0; r < len(board); r++ {
			if board[r][c] == player && board[r][c+1] == player && board[r][c+2] == player && board[r][c+3] == player {
				return true
			}
		}
	}
	for c := 0; c < len(board[0]); c++ {
		for r := 0; r < len(board)-3; r++ {
			if board[r][c] == player && board[r+1][c] == player && board[r+2][c] == player && board[r+3][c] == player {
				return true
			}
		}
	}
	for c := 0; c < len(board[0])-3; c++ {
		for r := 1; r <= (len(board) - 3); r++ {
			if board[r][c] == player && board[r+1][c+1] == player && board[r+2][c+2] == player && board[r+3][c+3] == player {
				return true
			}
		}
	}
	for c := 0; c < len(board[0])-3; c++ {
		for r := 3; r < len(board); r++ {
			if board[r][c] == player && board[r-1][c+1] == player && board[r-2][c+2] == player && board[r-3][c+3] == player {
				return true
			}
		}
	}
	return false
}
func checkIfColumnFull(board [][]uint8, column int) int {
	for i := 0; i < len(board); i++ {
		if board[i][column] == 0 {
			return i
		}
	}
	return -1
}
func printBoard(board [][]uint8) {
	for i := 1; i <= len(board[0]); i++ {
		fmt.Print("  " + strconv.FormatInt(int64(i), 10) + "  ")
	}
	fmt.Println()
	for i := len(board) - 1; i >= 0; i-- {
		for j := 0; j < len(board[i]); j++ {
			var symbol = " "
			if board[i][j] == 1 {
				symbol = "◯"
			} else if board[i][j] == 2 {
				symbol = "⬤"
			}
			fmt.Print("{ " + symbol + " }")
		}
		fmt.Println()
	}
}
func main() {
	var xDefault, yDefault int = 6, 7
	var x, y int = xDefault, yDefault
	var boardAnswer = ""
	re := regexp.MustCompile(`^\d+$`)
	fmt.Println("Do you want to change board dimensions (6x7 is default): (y/n)")
	fmt.Scanln(&boardAnswer)
	for boardAnswer != "y" && boardAnswer != "n" {
		fmt.Println(boardAnswer + " is not valid (y/n)")
		fmt.Scanln(&boardAnswer)
	}
	if boardAnswer == "y" {
		x, y = getNewNumber(xDefault, yDefault)
	}
	board := make([][]uint8, x)
	for i := range board {
		board[i] = make([]uint8, y)
	}
	printInitBoard(x, y)
	var column = ""
	var currentplayer uint8 = 1
	for !checkIfDraw(board) {
		printBoard(board)
		fmt.Println("Player " + strconv.FormatInt(int64(currentplayer), 10))
		fmt.Scanln(&column)
		if !re.MatchString(column) {
			fmt.Println("Invalid input")
			continue
		} else {
			columnNo, err := strconv.Atoi(column)
			if err != nil {
				fmt.Println("Error during conversion")
				continue
			}
			if columnNo > y {
				fmt.Println(column + " is bigger than number of columns on board " + strconv.FormatInt(int64(y), 10))
				continue
			} else {
				emptyRow := checkIfColumnFull(board, columnNo-1)
				if emptyRow == -1 {
					fmt.Println("That column is full")
					continue
				} else {
					board[emptyRow][columnNo-1] = currentplayer
				}
			}
		}
		if checkIfWon(board, currentplayer) {
			printBoard(board)
			fmt.Println("Player" + strconv.FormatInt(int64(currentplayer), 10) + " WON!!!")
			return
		}
		if currentplayer == 1 {
			currentplayer = 2
		} else {
			currentplayer = 1
		}
	}
	fmt.Println("YAWN ITS DRAW!!")
}
