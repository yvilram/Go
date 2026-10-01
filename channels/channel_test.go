package main

import "testing"


func TestHealth(t *testing.T) {
    if Health() != "OK" {
        t.Errorf("expected OK")
    }
}

func Health()string{
	return "OK"
}