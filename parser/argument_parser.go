package parser

import (
	"flag"

	"github.com/geoffrey-anto/go-cli-template/utils"
)

type Args struct {
}

func Parse_Arguments(args []string) (Args, error) {

	exampleFlag := flag.String("e", "", "The example flag")

	flag.Parse()

	err := utils.ValidateFlags(exampleFlag)

	if err != nil {
		return Args{}, err
	}

	return Args{}, nil
}
