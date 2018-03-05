package rules

import (
	"fmt"
	"github.com/hashicorp/consul/command/validate"
	"github.com/mitchellh/cli"
	"os"
)

func ConsulGenericRule(configDirPath string) int {

	cmd := validate.New(&cli.BasicUi{Writer: os.Stdout, ErrorWriter: os.Stderr})

	args := []string{configDirPath}

	fmt.Print("Consul generic rule: ")

	if code := cmd.Run(args); code != 0 {
		return code
	}

	return 0
}
