// Package newmath is a trivial example package.
package newmath

// sqrt returns an approximation to the square root of x.
func Sqrt(x float64) float64{
	z := 0.0
	for i :=0; i<1000; i++{
		z -= (z*z - x)/(2 * x)
	}
	return z
}
