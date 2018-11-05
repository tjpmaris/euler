package main

import (
	"fmt"
	"strconv"
)

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

func AddBig(block []string) string {
	var extraStuff = 0
	var sum = ""
	var i = 0

	for len(block) > 0 {
		if i < len(block) && len(block[i]) == 0 {
			fmt.Println("Removing Block")
			RemoveAt(block, i)
			continue
		}

		if i >= len(block) {
			i = 0
			fmt.Println("Setting i to 0")

			var extraStringStuff = strconv.Itoa(extraStuff)
			var len = len(extraStringStuff)

			sum = extraStringStuff[len-1:len] + sum

			extraStringStuff = extraStringStuff[0 : len-1]
			extraStuff, _ = strconv.Atoi(extraStringStuff)
		} else {

			var v = len(block[i]) - 1

			extraStuff += int(block[i][v] - '0')

			i++
		}
	}

	var extraStringStuff = strconv.Itoa(extraStuff)
	sum = extraStringStuff + sum

	return sum
}

func RemoveAt(value []string, index int) {
	copy(value[index:], value[index+1:])
	value = value[:len(value)-1]
}

func MaxLength(value []string) int {
	var maxlen = 0

	for _, element := range value {
		if maxlen < len(element) {
			maxlen = len(element)
		}
	}

	return maxlen
}

func AddStrings(value1 string, value2 string) string {
	var str1 = Reverse(value1)
	var str2 = Reverse(value2)
	var combined = ""

	var len1 = len(str1)
	var len2 = len(str2)

	var added = 0
	var index = 0

	for index < len1 || index < len2 {
		if index < len1 {
			added += int(rune(str1[index]) - '0')
		}

		if index < len2 {
			added += int(rune(str2[index]) - '0')
		}

		combined += Reverse(strconv.Itoa(added))[0:1]
		added /= 10

		index++
	}

	if added > 0 {
		combined += strconv.Itoa(added)
	}

	return Reverse(combined)
}

func Reverse(str string) string {
	var reverse = []rune(str)

	for i, v := 0, len(reverse)-1; i < v; i, v = i+1, v-1 {
		reverse[i], reverse[v] = reverse[v], reverse[i]
	}

	return string(reverse)
}
