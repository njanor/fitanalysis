package main

// GetFactorsExcluding returns the factors (if any) of the number, excluding the number itself
func GetFactorsExcluding(integer uint) []uint {
	factors := make([]uint, 0, 10)
	for i := uint(1); i < integer; i++ {
		if integer%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}
