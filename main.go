package main

import (
	"os"

	"github.com/geoffrey-anto/go-cli-template/app"
	"github.com/geoffrey-anto/go-cli-template/parser"
	"github.com/geoffrey-anto/go-cli-template/utils"
)

func main() {

	args, err := parser.Parse_Arguments(os.Args[1:])

	if err != nil {
		utils.FatalLog(err)
	}

	err = app.Run(args)

	if err != nil {
		utils.FatalLog(err)
	}
}
