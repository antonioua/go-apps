package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	wg.Add(1)

	go updateMessage("epsilon")
	wg.Wait()

	if msg != "epsilon" {
		t.Errorf("Expectd to find epsilon, but it's not there")
	}
}

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "epsilon"
	printMessage()

	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "epsilon") {
		t.Errorf("Expected to find epsilon, but it's not there")
	}
}

func Test_main(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, universe!") {
		t.Errorf("Doesn't find \"Hello, universe!\" phrase in stdout")
	}

	if !strings.Contains(output, "Hello, cosmos!") {
		t.Errorf("Doesn't find \"Hello, cosmos!\" phrase in stdout")
	}

	if !strings.Contains(output, "Hello, world!") {
		t.Errorf("Doesn't find \"Hello, world!\" phrase in stdout")
	}

}
