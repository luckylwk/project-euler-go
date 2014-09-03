// Created to test using different packages.
package lib


////////////////////////////////////////////////////////////////////////


// Returns all factors of an integer.
// Uses a technique called Trial-division. Keeps breaking the number
// down untill it can't be broken down further and moves to a higher number.
func Factorise_trialdivision( n uint64 ) []uint64 {
	output := []uint64{1}
	var div uint64 =  2
	for n > 1 {
		for n % div == 0 {
				output = append(output, div)
				n /= div
		}
		if div == 2 {
				div++
		} else {
				div += 2
		}
	}
	return output
}


// Returns all factors of an integer.
// Uses a technique called Wheel factorisation.
func Factorise_wheel( n uint64 ) []uint64 {
	output := []uint64{1}
	step_up := []uint64{1,2,2,4,2,4,2,4,6,2,6}
	var wheel uint64 = 3
	var div uint64 = 2
	var next uint64 = 0
	for n > 1 {
		for n % div == 0 {
			output = append(output, div)
			n /= div
		}
		div += step_up[next]
		next++
		if next == uint64(len(step_up)) {
			next = wheel
		}
	}
	return output
}


////////////////////////////////////////////////////////////////////////