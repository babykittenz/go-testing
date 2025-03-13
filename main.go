package main

import "fmt"

func main() {
	n := 2
	_, msg := isPrime(n)
	fmt.Println(msg)
}
func isPrime(n int) (bool, string) {

	// prime numbers do not include 0 or 1
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not a prime number", n)
	}
	// negative numbers are not prime
	if n < 0 {
		return false, fmt.Sprintf("Negative numbers are not prime numbers")
	}
	// use the modulus operator to repeatedly check to see if we have a prime number.
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not prime because it is not divisible by %d", n, i)
		}
	}
	return true, fmt.Sprintf("%d is a prime number", n)
}
