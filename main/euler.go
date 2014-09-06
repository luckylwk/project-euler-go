package main

import (
	"fmt"
	"math"
	"strconv"
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

func fn_euler_four( n int ) float64 {
	// Find the largest palindrome made from the product of two 3-digit numbers.
	// Define an upper and lower limit to search between.
	n_min := math.Pow( 10, float64(n-1) )
	n_max := math.Pow( 10, float64(n) ) - 1.0
	// From the upper limit construct all possible palindromes?
	palindrome_max := 0.0
	// Then cycle back using 3 digit numbers...
	for x := n_max; x > n_min; x-- {
		for y := n_max; y > n_min; y-- {
			p := x * y
			z := strconv.FormatFloat(p, 'f', 0, 64)
			if z == lib.Reverse(z) {
				if p > palindrome_max {
					palindrome_max = p
				}
			}
		}
	}
	return palindrome_max
}


// func fn_euler_five() int {
// 	// 2520 is the smallest number that can be divided by 
// 	// each of the numbers from 1 to 10 without any remainder.
// 	// What is the smallest positive number that is evenly 
// 	// divisible by all of the numbers from 1 to 20?
// 	var i int = 1
// 	var number int = 0
// 	var output int = 0
// 	for output == 0 {
// 		number = i * 20
// 		// Check if it is divisible by 19,18,...

// 		i++
// 	}

// }


func fn_euler_six( n int ) int {
	// Find the difference between the sum of the squares of the first 
	// one hundred natural numbers and the square of the sum.
	var sum int = 0 
	var sumsq int = 0
	for i := 1; i <= n; i++ {
		sum += i
		sumsq += i * i
	}
	return sum * sum - sumsq
}

func fn_euler_seven( n int ) int {
	// What is the 10001-st prime number?
	var counter int = 0
	for i := 1; i > 0; i++ {
		if lib.MillerRabin( uint32(i) ) == true {
			counter++
			if counter == n {
				return i
			}
		}
	}
	return 0
}

func fn_euler_twentyfive( n int ) int {
	// What is the first term in the Fibonacci sequence to contain 1000 digits?
	golden_ratio := 1.61803398875
	return int( math.Ceil( ( math.Log2(10)*(float64(n-1)) + math.Log2(5.0)/2.0 ) / math.Log2(golden_ratio) ) )
}


////////////////////////////////////////////////////////////////////////


func main(){

	// fmt.Println( "Euler Problem One: ", fn_euler_pOne( 1000 ) ) // 233168

	// fmt.Println( "Euler Problem Two: ", fn_euler_pTwo( 4000000 ) ) // 4613732

	// fmt.Println( "Euler Problem Three: ", lib.Factorise_wheel(600851475143) ) // 6857

	// fmt.Println( "Euler Problem Four: ", fn_euler_four( 3 ) ) // 906609

	// five

	// fmt.Println( "Euler Problem Six: ", fn_euler_six( 100 ) ) // 25164150

	// fmt.Println( "Euler Problem Seven: ", fn_euler_seven(10001) ) // 104743

	//

	//

	fmt.Println( "Euler Problem Twentyfive: ", fn_euler_twentyfive(1000) ) // 4782

}


