package config

import "testing"

func TestConfigValidation(t *testing.T) {
	if Validate(Default()) != nil {
		t.Fatal("default invalid")
	}
}
