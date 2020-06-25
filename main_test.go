package main

import "testing"

func TestHello(t *testing.T) {
	str, err := Hello(true)
	if err != nil {
		t.Error(err)
	}

	t.Log(str)
}

