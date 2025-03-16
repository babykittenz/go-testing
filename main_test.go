package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number"},
		{"not prime", 8, false, "8 is not prime because it is not divisible by 2"},
		{"negative", -1, false, "Negative numbers are not prime numbers"},
		{"zero", 0, false, "0 is not a prime number"},
		{"one", 1, false, "1 is not a prime number"},
	}
	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true got false", e.name)
		}
		if !e.expected && result {
			t.Errorf("%s: expected false got true", e.name)
		}
		if e.msg != msg {
			t.Errorf("%s: expected %s got %s", e.name, e.msg, msg)
		}
	}
}

func Test_prompt(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to write pipe
	os.Stdout = w

	prompt()

	// close our writer
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of the prompt function from our read pipe
	out, _ := io.ReadAll(r)

	// perform out test
	if string(out) != "->" {
		t.Errorf("incorrect prompt expected -> got %s", string(out))
	}
}

func Test_intro(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to write pipe
	os.Stdout = w

	intro()

	// close our writer
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of the prompt function from our read pipe
	out, _ := io.ReadAll(r)

	// perform out test
	if !strings.Contains(string(out), "Enter a whole number, and we'll tell you if it is prime or not. Enter q to quit.") {
		t.Errorf("intro test not correct got %s", string(out))
	}
}

func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "Please enter a whole number"},
		{name: "not a number", input: "not a number", expected: "Please enter a whole number"},
		{name: "q", input: "q", expected: ""},
		{name: "7", input: "7", expected: "7 is a prime number"},
		{name: "8", input: "8", expected: "8 is not prime because it is not divisible by 2"},
		{name: "negative", input: "-1", expected: "Negative numbers are not prime numbers"},
		{name: "zero", input: "0", expected: "0 is not a prime number"},
		{name: "one", input: "1", expected: "1 is not a prime number"},
	}
	for _, e := range tests {

		input := strings.NewReader(e.input)
		reader := bufio.NewScanner(input)
		res, _ := checkNumbers(reader)
		if !strings.EqualFold(res, e.expected) {
			t.Errorf("%s: expected %s, but got %s", e.name, e.expected, res)
		}
	}
}

func Test_readUserInput(t *testing.T) {
	// to test this function we need a channel and an instance of io.reader
	doneChan := make(chan bool)

	// create a reference to a bytes.buffer
	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}
