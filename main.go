package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// print a welcome message
	intro()
	// create a channel to indicate when the user wants to quit
	doneChan := make(chan bool)
	// start a go routine to read user input and run program
	go readUserInput(os.Stdin, doneChan)
	// block until the done chan gets a value
	<-doneChan

	// close the channel
	close(doneChan)
	// say goodbye
	fmt.Println("Goodbye")
}

func readUserInput(in io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(in)

	for {
		res, done := checkNumbers(scanner)

		if done {
			doneChan <- true
			return
		}

		fmt.Println(res)
		prompt()
	}
}
func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	// read scanner
	scanner.Scan()

	// check to see if the user wants to quit
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	// try to convert what the user typed into an int
	numToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter a whole number", false
	}

	_, msg := isPrime(numToCheck)
	return msg, false
}
func intro() {
	fmt.Println("Is it Prime?")
	fmt.Println("----------------")
	fmt.Println("Enter a whole number, and we'll tell you if it is prime or not. Enter q to quit.")
	prompt()
}

func prompt() {
	fmt.Print("->")
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
