package main

import (
	"bufio"
	"fmt"
	"github.com/itsjustabavaar/oh-my-gosh/cmd/handler"
	"os"
)

func main() {
	prompt := "$ "
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(prompt)
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		_, output, err := handler.HandleInput(input, &prompt)
		if err != nil {
			fmt.Println("Error:", err)
		}
		if output != "" {
			fmt.Println(output)
		}
	}
}
