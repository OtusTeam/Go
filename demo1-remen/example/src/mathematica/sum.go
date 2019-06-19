package mathematica

//Function Sum returns sum of several numbers
func Sum(args ...int) int {
	acc := int(0)
	for _, a := range args {
		acc += a
	}
	return acc
}
