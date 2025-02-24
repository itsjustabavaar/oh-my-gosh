package main

import (
	"bufio"
	"fmt"
	"github.com/itsjustabavaar/oh-my-gosh/cmd/handler"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(vars.Prompt)
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		output, code, err := handler.InputHandler(input)
		if err != nil {
			fmt.Println("Error:", err)
		}
		if output != "" {
			fmt.Println(output)
		}
		if code != nil {
			os.Exit(*code)
		}
	}
}
