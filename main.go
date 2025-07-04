package main

import (
	"os"

	"github.com/paullockaby/test-go/cmd"
)

func main() {
	err := cmd.Execute(os.Args[1:])
	if err != nil {
		os.Exit(-1)
	}
}
