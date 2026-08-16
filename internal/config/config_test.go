package config

import "testing"

func TestLoad(t *testing.T) {
	if Load().AppName == "" {
		t.Fatal("empty app name")
	}
}
