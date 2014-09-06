package main

import (
	"time"
	"fmt"
	"github.com/luckylwk/project-euler-go/lib"
	"github.com/cheggaaa/pb"
)


////////////////////////////////////////////////////////////////////////



////////////////////////////////////////////////////////////////////////


func main(){

	fmt.Printf("Factors of 600851475143 are: %v\n", lib.Factorise_trialdivision(600851475143) ) // 6857

	fmt.Printf("Factors of 600851475143 are: %v\n", lib.Factorise_wheel(600851475143) ) // 6857

	// Start timing.
	time_start := time.Now()
	fmt.Printf("Fibonacci testing (40): %-8.2f -- duration: %-6.2f sec. \n", lib.Fibonacci(40), time.Now().Sub(time_start).Seconds() )
	time_start = time.Now()
	fmt.Printf("Fibonacci testing (40): %-8.2f -- duration: %-6.2f sec. \n", lib.Fibonacci_approx(40), time.Now().Sub(time_start).Seconds() )

	// Test progressbar.
	fmt.Printf("Testing a progressbar by 'cheggaaa")
	count := 10000
    bar := pb.StartNew(count)
    for i := 0; i < count; i++ {
        bar.Increment()
        time.Sleep(time.Millisecond)
    }
    bar.FinishPrint("The End!")

}


