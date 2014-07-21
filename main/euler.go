package main

import (
	"time"
	"fmt"
	"math"
	"github.com/luckylwk/project-euler-go/lib"
)


////////////////////////////////////////////////////////////////////////


// Problem ONE
// Find the sum of all the multiples of 3 or 5 below 1000.
func fn_euler_pOne( n int ) int {
	
	var tmp_sum int = 0

	for i := 0; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			tmp_sum += i
		}
	}

	return tmp_sum

}

// Problem TWO
// By considering the terms in the Fibonacci sequence whose values do 
// not exceed four million, find the sum of the even-valued terms.
func fn_euler_pTwo( max_fib int ) int { 

	var tmp_sum int = 0
	var new_fib, this_fib, prev_fib int = 0, 1, 0

	for this_fib < max_fib {
		new_fib = this_fib + prev_fib
		if new_fib%2 == 0 {
			tmp_sum += new_fib
		}
		prev_fib = this_fib
		this_fib = new_fib
	}

	return tmp_sum

}


func fn_euler_four( n float64 ) int {
	// Define an upper and lower limit to search between.
	limit_lower := math.Pow( 1 * math.Pow(10,n-1), 2 )
	limit_upper := math.Pow( 1 * math.Pow(10,n) - 1 ,2 )
	fmt.Println( "Lower limit: ", limit_lower )
	fmt.Println( "Upper limit: ", limit_upper )
	// From the upper limit construct all possible palindromes?

	// Then cycle back using 3 digit numbers...

	return 1
}


////////////////////////////////////////////////////////////////////////


func main(){

	pOne := fn_euler_pOne( 1000 )
	fmt.Println( "Euler Problem One: ", pOne ) // 233168

	pTwo := fn_euler_pTwo( 4000000 )
	fmt.Println( "Euler Problem Two: ", pTwo ) // 4613732

	pThree := lib.Factorise_trialdivision(600851475143)
	fmt.Printf("Factors of 600851475143 are: %v\n", pThree) // 6857
	pThree_b := lib.Factorise_wheel(600851475143)
	fmt.Printf("Factors of 600851475143 are: %v\n", pThree_b) // 6857

	// Start timing.
	time_start := time.Now()
	fmt.Printf("Fibonacci testing (40): %-8.2f -- duration: %-6.2f sec. \n", lib.Fibonacci(40), time.Now().Sub(time_start).Seconds() )
	time_start = time.Now()
	fmt.Printf("Fibonacci testing (40): %-8.2f -- duration: %-6.2f sec. \n", lib.Fibonacci_approx(40), time.Now().Sub(time_start).Seconds() )

}


