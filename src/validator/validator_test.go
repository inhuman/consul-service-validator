package validator_test

import (
	"testing"
	"validator"
	"github.com/stretchr/testify/assert"
	"rules"
)
func init() {

}


func TestInit(t *testing.T) {
	validator.Clear()

	err := validator.Init("tests2")
	assert.Equal(t, "stat tests2: no such file or directory", err.Error())

	err = validator.Init("../../tests")
	assert.Equal(t, nil, err)
}

func TestAddRule(t *testing.T) {
	validator.Clear()

	err := validator.Init("../../tests")
	assert.Equal(t, nil, err)

	rls := validator.GetRules()
	assert.Equal(t, 0, len(rls))

	validator.AddRule(rules.ConsulGenericRule)

	rls = validator.GetRules()
	assert.Equal(t, 1, len(rls))
}

func TestRun(t *testing.T) {
	validator.Clear()

	code := validator.Run()
	assert.Equal(t, 2, code)

	err := validator.Init("../../tests")
	assert.Equal(t, nil, err)

	validator.AddRule(rules.ConsulGenericRule)

	code = validator.Run()
	assert.Equal(t, 0, code)

	validator.AddRule(rules.ServiceAddr)
	code = validator.Run()

	assert.Equal(t, 1, code)
}


