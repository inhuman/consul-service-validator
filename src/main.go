package main

import (
	"validator"
	"rules"
	"os"
)

func main() {

	validator.Init(os.Args[1])
	validator.AddRule(rules.ConsulGenericRule)
	validator.AddRule(rules.ServiceAddr)
	validator.Run()
}
