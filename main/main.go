package main

import (
	"github.com/kkcaZ/advent-2024/main/app"
	"os"
)

func main() {
	if err := app.Run(nil); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
