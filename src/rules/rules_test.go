package rules_test

import (
	"testing"
	"github.com/hashicorp/consul/testutil"
	"os"
	"path/filepath"
	"io/ioutil"
	"rules"
	"github.com/stretchr/testify/assert"
	"validator"
)


// Other tests for consul rule see in:
// https://github.com/hashicorp/consul/blob/master/command/validate/validate_test.go
func TestConsulGenericRule(t *testing.T) {
	validator.Clear()

	td := testutil.TempDir(t, "consul")
	defer os.RemoveAll(td)

	fp := filepath.Join(td, "config.json")
	err := ioutil.WriteFile(fp, []byte(`{"bind_addr":"10.0.0.1", "data_dir":"`+td+`"}`), 0644)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	code := rules.ConsulGenericRule(fp)

	assert.Equal(t, 0, code)
}

func TestServiceAddr(t *testing.T) {
	validator.Clear()

	td := testutil.TempDir(t, "consul")
	defer os.RemoveAll(td)

	fp := filepath.Join(td, "config.json")
	err := ioutil.WriteFile(fp, []byte(`{"bind_addr":"10.0.0.1", "data_dir":"`+td+`"}`), 0644)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	code := rules.ServiceAddr(fp)

	assert.Equal(t, 0, code)
}