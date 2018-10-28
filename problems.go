package main

import (
	"fmt"
	"strconv"
)

/*
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5 below 1000.
*/
func problem1() int {
	var value = 0

	for i := 3; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			value += i
		}
	}

	return value
}

/*
Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
*/
func Problem2() int {
	var value = 0

	for i, v := 0, 1; v < 4000000; {
		var sum = i + v

		fmt.Printf("%d, %d, %d\n", i, v, sum)

		if sum%2 == 0 {
			value += sum
			fmt.Printf("value: %d\n", value)
		}

		i, v = v, sum
	}

	return value
}

/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/
func Problem3() int {
	var value = 0
	var num = 600851475143

	for _, element := range FactorsOf(num) {
		fmt.Printf("%d, %t\n", element, IsPrime(element))
	}

	return value
}

/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

func Problem4() int {
	var palindromes []int
	var max = 999

	for i, v := 0, 0; i <= max; v++ {
		if v == max {
			v = i
			i++
			continue
		}

		var product = i * v
		if IsPalindrome([]rune(strconv.Itoa(product))) {
			palindromes = append(palindromes, product)
		}
	}

	return Largest(palindromes)
}

/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/
func Problem5() int {
	var values = []int{20, 19, 18, 17, 16, 15, 14, 13, 12, 11}

	var i = Largest(values)

	for true {
		if IsDivisibleBy(i, values) {
			return i
		}

		i++
	}

	return 0
}
