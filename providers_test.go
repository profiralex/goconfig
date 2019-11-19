package goconfig_test

import (
	"os"
	"github.com/profiralex/goconfig"
	"testing"
)

func TestEnvProviderKeyDoesNotExist(t *testing.T) {
	provider := goconfig.EnvProvider{}
	value, err := provider.Lookup(testStringKey)

	if len(value) > 0 {
		t.Errorf("Unexpected value %s", value)
	}

	if err == nil {
		t.Errorf("Expected error not received")
	}
}

func TestEnvProviderKeyExists(t *testing.T) {
	os.Setenv(testStringKey, testStringValue)
	defer os.Unsetenv(testStringKey)

	provider := goconfig.EnvProvider{}
	value, err := provider.Lookup(testStringKey)

	if value != testStringValue {
		t.Errorf("Expected value %s got %s", testStringValue, value)
	}

	if err != nil {
		t.Errorf("Unexpected error recived %w", err)
	}
}
