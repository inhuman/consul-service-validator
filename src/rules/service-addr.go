package rules

import (
	"fmt"
	"github.com/hashicorp/consul/agent/config"
	"os"
	"strings"
	"regexp"
)

func ServiceAddr(configDirPath string) error {

	fmt.Print("Service adders rule: ")

	b, err := config.NewBuilder(config.Flags{ConfigFiles: []string{configDirPath}})

	rt, err := b.Build()

	for _, service := range rt.Services {

		if !validIP4(service.Address) {
			fmt.Println("For service " + service.Name + " validation failed. Address " + service.Address + " is incorrect.")
		} else {
			fmt.Println("Service " + service.Name + " address is valid ip.")
		}
	}

	if err != nil {
		fmt.Sprintf("Config build failed: %v", err.Error())
		os.Exit(1)
	}

	if err != nil {
		fmt.Sprintf("Config validation failed: %v", err.Error())
		os.Exit(1)
	}

	return nil
}

func validIP4(ipAddress string) bool {
	ipAddress = strings.Trim(ipAddress, " ")

	re, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	if re.MatchString(ipAddress) {
		return true
	}
	return false
}
