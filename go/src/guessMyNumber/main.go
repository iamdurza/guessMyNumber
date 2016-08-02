package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	GuessNumber := bufio.NewScanner(os.Stdin)
	var GN int
	fmt.Println("Input a number to be guessed: ")
	GuessNumber.Scan()
	randr := GuessNumber.Text()
	GN, _ = strconv.Atoi(randr)
	//begin game
	for {
		fmt.Println("guess a number: ")
		GuessNumber.Scan()
		rando := GuessNumber.Text()
		rn, _ := strconv.Atoi(rando)
		//if too low
		if rn < GN {
			fmt.Println("Too low, Try Again Noob:")
		}
		//if too high
		if rn > GN {
			fmt.Println("Too high, Try again Noob: ")
		}
		//if is =
		if rn == GN {
			break
		}
	}
	fmt.Println("You win")
}
