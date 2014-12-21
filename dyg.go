package main

import (
	"bufio"
	"fmt"
	"os"

	otto "github.com/glycerine/dynamic-go/backend"
)

var OttoVm = otto.New()

func main() {
	repl()
}

func evaluate(line string) string {
	if len(line) == 0 {
		return line
	}
	value, err := OttoVm.Run(line)
	if err != nil {
		return err.Error()
	} else {
		return fmt.Sprintf("%#v", value)
	}
}

func repl() {
	for {
		fmt.Print("dyg> ")
		in := bufio.NewReader(os.Stdin)
		line, err := in.ReadString('\n')
		if err != nil {
			fmt.Printf("error during read of input: %s\n", err)
			continue
		}
		if len(line) == 0 {
			continue
		}
		res := evaluate(line)
		fmt.Printf("result: '%s'\n", res)
	}
}
