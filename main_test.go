package main

import "testing"

func Test_isPrime(t *testing.T) {
	result, msg := isPrime(0)
	if result {
		t.Errorf("Expected false, got %t", result)
	}
	if msg != "0 is not a prime number" {
	}

	result, msg = isPrime(1)
	if result {
		t.Errorf("Expected false, got %t", result)
	}
	if msg != "1 is not a prime number" {
	}

	result, msg = isPrime(2)
	if !result {
		t.Errorf("Expected true, got %t", result)
	}
	if msg != "2 is a prime number" {
	}
}
