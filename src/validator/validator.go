package validator

import (
	"os"
	"fmt"
)

// The rule function require path to service config folder and return exit code.
// If validation failed function must say it itself and return non zero exit code.
// This is so that you can see the result of all the rules in one run.
type Rule func(string) int

var rules []Rule
var configDirPath string

func Init(path string) error {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("Config direcory does not exists:", path)
		return err
	}
	configDirPath = path
	fmt.Println("Validator initialized with consul config dir:", path)

	return nil
}

func AddRule(rule Rule) {
	rules = append(rules, rule)
}

func GetRules() []Rule {
	return rules
}

func Run() int {

	if configDirPath == "" {
		fmt.Println("You must run init method first")
		return 2
	}

	if len(rules) >= 1 {
		for _, fn := range rules {
			code := fn(configDirPath)
			if code != 0 {
				return code
			}
		}
	} else {
		fmt.Println("You must add at least one rule")
		return 1
	}

	return 0
}

func Clear() {
	rules = []Rule{}
	configDirPath = ""
}