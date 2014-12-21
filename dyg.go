package main

import (
	"bufio"
	"fmt"
	"os"

	otto "github.com/glycerine/dyg/backend"
)

var OttoVm = otto.New()

func ImaCompiledFunc(str string) string {
	return fmt.Sprintf("Hi, you are running a compiled-in-golang function, with input from dyg: '%s'", str)
}

func main() {

	OttoVm.Set("comp", func(call otto.FunctionCall) otto.Value {
		s := ImaCompiledFunc(call.Argument(0).String())
		v, _ := otto.ToValue(s)
		return v
	})

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
		fmt.Printf("%v\n", res)
	}
}
