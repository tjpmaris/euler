package main

func IsPrime(value int) bool {
	return len(FactorsOf(value)) == 0
}

func IsPalindrome(value []rune) bool {
	var length = len(value)

	for i, v := 0, length-1; i < length/2; i, v = i+1, v-1 {
		if value[i] != value[v] {
			return false
		}
	}

	return true
}

func Largest(value []int) int {
	var length = len(value)
	var largest = value[0]

	for i := 1; i < length; i++ {
		if value[i] > largest {
			largest = value[i]
		}
	}

	return largest
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

func IsDivisibleByAll(num int, values []int) bool {
	for _, element := range values {
		if num%element != 0 {
			return false
		}
	}

	return true
}

func IsDivisibleByAny(num int, values []int) bool {
	for _, element := range values {
		if num%element == 0 {
			return true
		}
	}

	return false
}
