package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
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

func main() {
	var xDefault, yDefault int = 6, 7
	var x, y int = xDefault, yDefault
	fmt.Println("Do you want to change board dimensions (6x7 is default): (y/n)")
	var boardAnswer = ""
	fmt.Scanln(&boardAnswer)
	for boardAnswer != "y" && boardAnswer != "n" {
		fmt.Println(boardAnswer + " is not valid (y/n)")
		fmt.Scanln(&boardAnswer)
	}
	if boardAnswer == "y" {
		x, y = getNewNumber(xDefault, yDefault)
	}
	fmt.Println(x + y)
}
