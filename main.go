package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getNewNumber(minX uint8, minY uint8) (uint8, uint8) {
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
						return uint8(x), uint8(y)
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
func loadGame(name string) ([]uint8, []uint8, []byte) {
	player1, e := os.ReadFile("saves/" + name + "/1")
	player2, er := os.ReadFile("saves/" + name + "/2")
	d, err := os.ReadFile("saves/" + name + "/dimension")
	if e != nil || er != nil || err != nil {
		fmt.Println(e)
		fmt.Println(er)
		fmt.Println(err)
		return nil, nil, nil
	}
	return player1, player2, d
}
func main() {
	var xDefault, yDefault uint8 = 6, 7
	var x, y uint8 = xDefault, yDefault
	var boardAnswer = ""
	var column = ""
	var currentplayer uint8 = 1
	dimensionArray := []byte{x, y}
	history := make([][]uint8, 2)

	isColumn := regexp.MustCompile(`^\d+$`)
	isValidAnswer := regexp.MustCompile(`[yn]$`)

	fmt.Println("Do you want load game: (y/n)")
	fmt.Scanln(&boardAnswer)

	for !isValidAnswer.MatchString(boardAnswer) {
		fmt.Println(boardAnswer + " is not valid (y/n)")
		fmt.Scanln(&boardAnswer)
	}
	if boardAnswer == "y" {
		fmt.Println("Type name of saved game")
		fmt.Scanln(&boardAnswer)
		history[0], history[1], dimensionArray = loadGame(boardAnswer)
		for history[0] == nil {
			fmt.Println("Type name of saved game")
			fmt.Scanln(&boardAnswer)
			history[0], history[1], dimensionArray = loadGame(boardAnswer)
		}
		x = dimensionArray[0]
		y = dimensionArray[1]
	} else {
		fmt.Println("Do you want to change board dimensions (6x7 is default): (y/n)")
		fmt.Scanln(&boardAnswer)

		for !isValidAnswer.MatchString(boardAnswer) {
			fmt.Println(boardAnswer + " is not valid (y/n)")
			fmt.Scanln(&boardAnswer)
		}
		if boardAnswer == "y" {
			x, y = getNewNumber(xDefault, yDefault)
		}
	}

	board := make([][]uint8, x)
	for i := range board {
		board[i] = make([]uint8, y)
	}
	if history[0] != nil {
		for i := 0; i < len(history[0])+len(history[1]); i++ {
			emptyRow := checkIfColumnFull(board, int(history[int64(i%2)][int64(i/2)])-1)
			board[emptyRow][int(history[int64(i%2)][int64(i/2)])-1] = uint8(i%2) + 1
		}
	}

	fmt.Println("If u wish to save game do by typing save")

	for !checkIfDraw(board) {

		fmt.Print("Player 1:")
		for i := 0; i < len(history[0]); i++ {
			fmt.Print(" " + strconv.FormatInt(int64(history[0][i]), 10) + " ")
		}
		fmt.Println("")
		fmt.Print("Player 2:")
		for i := 0; i < len(history[1]); i++ {
			fmt.Print(" " + strconv.FormatInt(int64(history[1][i]), 10) + " ")
		}
		fmt.Println("")
		fmt.Println("**History**")
		printBoard(board)
		fmt.Println("Player " + strconv.FormatInt(int64(currentplayer), 10))
		fmt.Scanln(&column)
		save := strings.Fields(column)
		if save[0] == "save" {
			fmt.Println("Give name to save")
			fmt.Scanln(&column)
			err := os.MkdirAll("saves/"+column, 0755)
			check(err)
			err1 := os.WriteFile("saves/"+column+"/1", history[0], 0644)
			check(err1)
			err2 := os.WriteFile("saves/"+column+"/2", history[1], 0644)
			check(err2)

			err3 := os.WriteFile("saves/"+column+"/dimension", dimensionArray, 0644)
			check(err3)
			return
		}
		if !isColumn.MatchString(column) {
			fmt.Println("Invalid input")
			continue
		} else {
			columnNo, err := strconv.Atoi(column)
			if err != nil {
				fmt.Println("Error during conversion")
				continue
			}
			if uint8(columnNo) > y || uint8(columnNo) == 0 {
				fmt.Println(column + "number is bigger than number of columns on board or it is 0" + strconv.FormatInt(int64(y), 10))
				continue
			} else {
				emptyRow := checkIfColumnFull(board, columnNo-1)
				if emptyRow == -1 {
					fmt.Println("That column is full")
					continue
				} else {
					history[currentplayer-1] = append(history[currentplayer-1], uint8(columnNo))
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}
