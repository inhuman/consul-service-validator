package main

import (
	"validator"
	"rules"
	"os"
	"fmt"
)

func main() {

	validator.Init(os.Args[1])
	validator.AddRule(rules.ConsulGenericRule)
	validator.AddRule(rules.ServiceAddr)

	exitCode := validator.Run()

	fmt.Println("Exit code:", exitCode)

	os.Exit(exitCode)
}
