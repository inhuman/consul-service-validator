package rules

import (
	"fmt"
	"github.com/hashicorp/consul/command/validate"
	"github.com/mitchellh/cli"
	"os"
)

func ConsulGenericRule(configDirPath string) error {

	cmd := validate.New(&cli.BasicUi{Writer: os.Stdout, ErrorWriter: os.Stderr})

	args := []string{configDirPath}

	fmt.Print("Consul generic rule: ")

	if code := cmd.Run(args); code != 0 {
		os.Exit(1)
	}

	return nil
}
