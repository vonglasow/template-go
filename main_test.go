package main

import "testing"

func TestMain(t *testing.T) {
	err := ""
	if err != "" {
		t.Logf("error loading config: %v", err)
		t.Fail()
	}
}
