package main

func IsPrime(value int) bool {
	var max = value

	for i := 2; i < max; i++ {
		if value%i == 0 {
			return false
		} else {
			max = (value / i) + 1
		}
	}

	return true
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
	var max = value
	var bigFactors []int

	for i := 2; i < max; i++ {
		var otherFactor = value / i

		if value%i == 0 {
			factors = append(factors, i)
			bigFactors = append(bigFactors, otherFactor)
			max = otherFactor
		} else {
			max = otherFactor + 1
		}
	}

	factors = append(factors, bigFactors...)

	return factors
}

func NumFactors(value int) int {
	var max = value
	var num = 0

	for i := 1; i < max; i++ {
		var otherFactor = value / i

		if value%i == 0 {
			num += 2
			max = otherFactor
		} else {
			max = otherFactor + 1
		}
	}

	return num
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

func Contains(array []int, value int) bool {
	for _, element := range array {
		if element == value {
			return true
		}
	}

	return false
}

func CountSteps(value int, dic map[int]int) int {
	if dic[value] == 0 {
		if value%2 == 0 {
			dic[value] = CountSteps(value/2, dic) + 1
		} else {
			dic[value] = CountSteps((value*3)+1, dic) + 1
		}
	}

	return dic[value]
}
