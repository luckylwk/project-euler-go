// Created to test using different packages.
package lib

import "math"


////////////////////////////////////////////////////////////////////////


// Returns the n-th Fibonacci number by recursion.
func Fibonacci( n int ) float64 {
	if n == 0 {
		return 0.0
	} else if n == 1 {
		return 1.0
	} else {
		return Fibonacci(n-1) + Fibonacci(n-2)
	}
}

// Function to calculate the Fibonacci number directly.
func Fibonacci_approx( n float64 ) float64 {
	return ( math.Pow( 1 + math.Sqrt(5), n ) - math.Pow( 1 - math.Sqrt(5), n ) ) / ( math.Pow(2,n) * math.Sqrt(5) )
}


////////////////////////////////////////////////////////////////////////