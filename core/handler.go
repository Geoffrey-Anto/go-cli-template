package core

import (
	"time"

	"github.com/briandowns/spinner"
	"github.com/geoffrey-anto/go-cli-template/parser"
	"github.com/geoffrey-anto/go-cli-template/utils"
)

func Convert(args parser.Args) error {
	s := utils.GetSpinner(spinner.CharSets[9], 2*time.Second, "  xxxxxxxxxxxxxxxxxxxxxxxxxxx\n", "  xxxxxxxxxxxxxxxxxxxxxxxxxxx\n")
	s.Start()

	time.Sleep(4 * time.Second)

	s.Stop()

	return nil
}
