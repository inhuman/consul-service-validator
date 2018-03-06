package main

import (
	"github.com/inhuman/consul-service-validator/src/validator"
	"github.com/inhuman/consul-service-validator/src/rules"
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
