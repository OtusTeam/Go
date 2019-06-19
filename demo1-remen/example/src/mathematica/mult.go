package mathematica

//Function Mult returns multiplication of several numbers
func Mult(args ...int) int {
	acc := int(1)
	for _, a := range args {
		acc *= a
	}
	return acc
}
