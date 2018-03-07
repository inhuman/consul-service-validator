# Consul service config validator
Validator consul service json files

[![Build Status](https://travis-ci.org/inhuman/consul-service-validator.svg?branch=master)](https://travis-ci.org/inhuman/consul-service-validator)


[![Coverage Status](https://coveralls.io/repos/github/inhuman/consul-service-validator/badge.svg?branch=master)](https://coveralls.io/github/inhuman/consul-service-validator?branch=master)
**Usage**

_Fail_
```bash
$ validator /fail/config
Validator initialized with consul config dir: tests/
Consul generic rule: Configuration is valid!
Service address rule: For service api-tmt validation failed. Address https://172.29.10.77 is incorrect.
Exit code: 1
```

_Success_
```bash
$ validator /success/config
Validator initialized with consul config dir: tests/
Consul generic rule: Configuration is valid!
Service address rule: Service api-tmt address is valid ip.
Configuration is valid!
Exit code: 0
```

**Own rules example**

1. Implement rule; it takes as argument config file or dir and return exit code 
```go
func ServiceAddr(configDirPath string) int {

	fmt.Print("Service address rule: ")

	b, err := config.NewBuilder(config.Flags{ConfigFiles: []string{configDirPath}})

	rt, err := b.Build()

	code := 0

	for _, service := range rt.Services {

		if !validIP4(service.Address) {
			fmt.Println("For service " + service.Name + " validation failed. Address " + service.Address + " is incorrect.")
			code = 1
		} else {
			fmt.Println("Service " + service.Name + " address is valid ip.")
		}
	}

	if err != nil {
		fmt.Sprintf("Config build failed: %v", err.Error())
		code = 1
	}

	if err != nil {
		fmt.Sprintf("Config validation failed: %v", err.Error())
		code = 1
	}

	if code == 0 {
		fmt.Println("Configuration is valid!")
	}

	return code
}
```

2. Add rule to validator
```go
validator.AddRule(rules.ServiceAddr)
``` 

3. PROFIT!

P.S. If you want help project - pull requests are welcome )
