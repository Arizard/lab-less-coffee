package main

import (
	"fmt"
	"os"

	"github.com/arizard/lab-less-coffee/pkg/spc"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	cmd := args[0]

	if cmd == "0" {
		fmt.Println(spc.OverlapNone)
	}
}
