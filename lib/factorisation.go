// Created to test using different packages.
package lib


////////////////////////////////////////////////////////////////////////


// Returns all factors of an integer.
// Uses a technique called Trial-division. Keeps breaking the number
// down untill it can't be broken down further and moves to a higher number.
func Factorise_trialdivision( n int ) []int {
	output := []int{1}
	div := 2
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
func Factorise_wheel( n int ) []int {
	output := []int{1}
	step_up := []int{1,2,2,4,2,4,2,4,6,2,6}
	wheel := 3
	div := 2
	next := 0
	for n > 1 {
		for n % div == 0 {
			output = append(output, div)
			n /= div
		}
		div += step_up[next]
		next++
		if next == len(step_up) {
			next = wheel
		}
	} 
	return output
}


////////////////////////////////////////////////////////////////////////