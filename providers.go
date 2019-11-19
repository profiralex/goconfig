package goconfig

import (
	"fmt"
	"os"
)

// Provider provides data from a certain source
type Provider interface {
	Lookup(string) (string, error)
}

// EnvProvider loads values from environment
type EnvProvider struct {
}

// Lookup looks up for value in the environment
func (p *EnvProvider) Lookup(key string) (string, error) {
	value, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("Value not found in the environment")
	}

	return value, nil
}
