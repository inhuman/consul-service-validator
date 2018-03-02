package validator

import (
	"os"
	"fmt"
)

type Rule func(string) error

var rules []Rule

var configDirPath string

func Init(path string) {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("Config direcory does not exists:", path)
		os.Exit(1)
	}
	configDirPath = path
}

func AddRule(rule Rule) {
	rules = append(rules, rule)
}

func Run() {
	for _, fn := range rules {
		err := fn(configDirPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
