// Created to test using different packages.
package lib


////////////////////////////////////////////////////////////////////////


// Sqrt returns an approximation to the square root of x.
func Sqrt( x float64 ) float64 {
	z := 1.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}


////////////////////////////////////////////////////////////////////////