package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const duckArt string = `
          __
      ___( o)>  "quack"
      \ <_. )
       ^---'  
` // source https://textart.io/art/OUmq7JexuhjpvBnTpL_HAQeF/duckling-swimming

func main() {

	fmt.Println("Tell the duck your problems.")
	fmt.Println()

	for {

		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("something went wrong... are you speaking duck?")
			continue
		}

		if len(strings.Trim(input, "\n")) == 0 {
			continue
		}

		for i := 0; i < 3; i++ {
			time.Sleep(1 * time.Second)
			fmt.Print(".")
		}
		time.Sleep(2 * time.Second)

		fmt.Print(duckArt)

		time.Sleep(time.Second)

	}

}
