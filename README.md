# DiceGo

Quick mess-around with the Go programming language.

## Usage

Import this package by running this command in your terminal of choice:

`go get github.com/AlexHolderDeveloper/DiceGo`

Then, you can create code like this in your Go files:

```go
package main

import (
    "fmt"
    DiceGo "github.com/AlexHolderDeveloper/DiceGo"
)

func main() {
    fmt.Println(DiceGo.RollDice(20))
}

```

You may want to run `go mod tidy` before running your code. But that's it! 🎉