package godice

import "math/rand"
import "time"


// Returns a random number between 1 and diceSize, inclusive.
// NOTE: A function is only exported for others to use if it begins with a capital letter.
// Usage: 
/*
package main

import (
    "fmt"
    DiceGo "github.com/AlexHolderDeveloper/DiceGo"
)

func main() {
    fmt.Println(DiceGo.RollDice(20))
}

*/
func RollDice(diceSize int) int {
	diceMin := 1
	diceMax := diceSize
	rand.Seed(time.Now().UnixNano())

	
	return rand.Intn(diceMax-diceMin) + diceMin
}
