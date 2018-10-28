package main

func IsPrime(value int) bool {
	return len(FactorsOf(value)) == 0
}

func IsPalindrome(value []rune) bool {
	var length = len(value)

	for i := 0; i < length; i++ {
		if value[i] != value[length-(i+1)] {
			return false
		}
	}

	return true
}

func FactorsOf(value int) []int {
	var factors []int
	var biggestPossibleOtherFactor = value
	var bigFactors []int

	for i := 2; i < biggestPossibleOtherFactor; i++ {
		var otherFactor = value / i

		if value%i == 0 {
			factors = append(factors, i)
			bigFactors = append(bigFactors, otherFactor)
		} else {
			biggestPossibleOtherFactor = otherFactor + 1
		}
	}

	for i := len(bigFactors) - 1; i >= 0; i-- {
		factors = append(factors, bigFactors[i])
	}

	return factors
}
