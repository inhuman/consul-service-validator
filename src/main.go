package main

import (
	"validator"
	"rules"
	"os"
	"fmt"
)

func main() {

	if len(os.Args) > 1 {
		validator.Init(os.Args[1])
		validator.AddRule(rules.ConsulGenericRule)
		validator.AddRule(rules.ServiceAddr)

		os.Exit(
			validator.Run())

	} else {
		fmt.Println(help)
		os.Exit(1)
	}
}

const help = `
Usage: validator FILE_OR_DIRECTORY [OPTIONS]
  Performs a thorough sanity test on Consul configuration files. For each file
  or directory given, the validate command will attempt to parse the contents
  just as the "consul agent" command would, and catch any errors.
  This is useful to do a test of the configuration only, without actually
  starting the agent. This performs all of the validation the agent would, so
  this should be given the complete set of configuration files that are going
  to be loaded by the agent. This command cannot operate on partial
  configuration fragments since those won't pass the full agent validation.
  Returns 0 if the configuration is valid, or 1 if there are problems.

  Also possible to use custom validation rules.
`
