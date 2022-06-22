package main

import "fmt"

func main() {

	fmt.Println(findPrimeNUmberLargerThanInput(10))
	fmt.Println(sumOfPrimeDivisors(12))

}

func findPrimeNUmberLargerThanInput(prime int) int {
	n := 1
	var nextPrime int
	for findFibonacci(n) < prime { //finds the number of fibonacci value after the entered input
		n++
	}
	nextPrime = findFibonacci(n)

	for isPrime(nextPrime) == false { //checks the primality of the found value
		n++
		nextPrime = findFibonacci(n)
	}
	return nextPrime
}
func findFibonacci(n int) int { //finds the fibonacci value of the entered n value
	a := 0
	b := 1

	for i := 0; i < n; i++ {
		temp := a
		a = b
		b = temp + a
	}
	return a
}

func isPrime(n int) bool { // This func Checks the primality of the entered Fibonacci number
	control := true
	for i := 2; i < n; i++ {
		if n%2 == 0 {
			control = false
		}
	}
	return control
}

func sumOfPrimeDivisors(n int) int { // This func find the sum of input value
	sum := 0
	for i := 2; i < n; i++ {
		if n%i == 0 {
			if isPrime(i) {
				sum += i
			}
		}
	}
	return sum
}
