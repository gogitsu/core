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

func TestGetString(t *testing.T) {
	ReadConfig()
	if val := GetString("client.balancing.strategy"); val == "" {
		t.Error("client.balancing.strategy is not valued")
	}
}

func TestGet(t *testing.T) {
	if val := Get("client.timeout"); val == nil {
		t.Error("client.timeout is nil")
	}
}

func TestGetStruct(t *testing.T) {
	type cfg struct {
		Component struct {
			Name string
		}
		Version string
	}
	var sconf *cfg
	if err := GetStruct(&sconf); err != nil {
		t.Error("unable to unmarshal configuration struct")
	} else if sconf.Component.Name != "cmp1" {
		t.Error("Component.Name is not 'cmp1'")
	}
}

func TestGetBkpDb(t *testing.T) {
	type cfg struct {
		BkpDB DB `yaml:"bkp-db"`
	}
	var sconf *cfg
	if err := GetStruct(&sconf); err != nil {
		t.Error("unable to unmarshal configuration struct")
	} else if sconf.BkpDB.Host != "localhost" {
		t.Logf("sconf: %+v\n", sconf)
		t.Logf("sconf.BkpDB.Host: %s\n", sconf.BkpDB.Host)
		t.Error("BkpDB.Host is not 'localhost'")
	}
}
