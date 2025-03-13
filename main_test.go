package main

import "testing"

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
