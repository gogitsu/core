package config

import (
	"os"
	"testing"
)

type TestConfiguration struct {
	Service struct {
		Group string
		Name  string
	}
}

func init() {
	os.Setenv("ENV", "test")
}

func TestConfig(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("PANIC: %+v\n", r)
		}
	}()

	var c *Configuration

	Config(&c)
	if c == nil {
		t.Error("c is nil")
	}
}

func TestTestConfiguration(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("PANIC: %+v\n", r)
		}
	}()

	var tc *TestConfiguration
	ReConfig(&tc)
	if tc == nil {
		t.Error("tc is nil")
	} else if tc.Service.Group != "test-group" {
		t.Errorf("tc.Service.Group != 'test-group'... is: %s\n", tc.Service.Group)
	}
}
