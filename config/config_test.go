package config

import (
	"log"
	"testing"
)

func TestServiceStruct(t *testing.T) {
	log.Println("testing config.Service struct")
	s := Service{
		Group:   "test-group",
		Name:    "test-name",
		Version: "0.0.1",
	}

	if s.Group != "test-group" {
		t.Error("s.Group != 'test-group'")
	} else if s.Name != "test-name" {
		t.Error("s.Name != 'test-name'")
	} else if s.Version != "0.0.1" {
		t.Error("s.Version != '0.0.1'")
	}
}
