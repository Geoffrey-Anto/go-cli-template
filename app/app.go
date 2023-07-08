package app

import (
	"fmt"

	"github.com/geoffrey-anto/go-cli-template/core"
	"github.com/geoffrey-anto/go-cli-template/parser"
	"github.com/geoffrey-anto/go-cli-template/utils"
)

func Run(args parser.Args) error {
	colors := utils.Colors{}
	fmt.Println(string(colors.Purple()), " Welcome to go-cli-template! ")
	fmt.Print(string(colors.Reset()), " Press y to continue, or any other key to abort.  ")

	if !utils.CanProceed() {
		fmt.Println("aborted.")
		return fmt.Errorf("aborted")
	} else {
		core.Convert(args)
		return nil
	}
}
