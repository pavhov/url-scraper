package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(t *testing.M) {
	fmt.Println()
	// setup()
	exitVal := t.Run()
	if exitVal == 0 {
		// teardown()
	}
	os.Exit(exitVal)
}
