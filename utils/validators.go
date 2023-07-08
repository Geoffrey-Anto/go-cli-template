package utils

import "fmt"

func CanProceed() bool {
	var canproceed string
	fmt.Scanln(&canproceed)
	return canproceed == "y"
}

func ValidateFlags(exampleFlag *string) error {

	return nil
}
