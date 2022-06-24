package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsPrime(t *testing.T) {
	control := isPrime(7)
	require.Equal(t, control, true)

	control = isPrime(8)
	require.Equal(t, control, false)

	control = isPrime(-1)
	require.Equal(t, control, false)
}
func TestSumOfPrimeDivisors(t *testing.T) {
	number := sumOfPrimeDivisors(8)
	require.Equal(t, number, 2)

	number = sumOfPrimeDivisors(9)
	require.NotEqual(t, number, 4)
}
func TestFindFibonacci(t *testing.T) {
	fib := findFibonacci(5)
	require.Equal(t, fib, 5)

	fib = findFibonacci(4)
	require.NotEqual(t, fib, 2)

}
func TestFindFibonacciLargerThanInput(t *testing.T) {
	res := findPrimeNUmberLargerThanInput(10)

	require.Equal(t, res, 13)
	res = findPrimeNUmberLargerThanInput(8)
	require.NotEqual(t, res, 11)
}
