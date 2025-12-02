package common

// 1 * 36 = 36
// 2 * 18 = 36
// 3 * 12 = 36
// 4 * 9  = 36
// 6 * 6  = 36
func Divisors(n int) []int {
	small := []int{}
	large := []int{}

	// first part till sqrt(n)
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			small = append(small, i)

			// make sure that there are no duplicates
			if i != n/i {
				large = append(large, n/i)
			}
		}
	}

	for i := len(large) - 1; i >= 0; i-- {
		small = append(small, large[i])
	}

	return small
}

// first half of divisor <= n/2 usefull for some things
func LowerHalfDivisors(n int) []int {
	small := []int{}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			small = append(small, i)
		}
	}

	return small
}

// must be greater than 10
func SameDigits(n int) int {
	if n < 10 {
		return 0
	}
	lastDigit := n % 10
	withoutLast := n / 10

	for withoutLast > 0 {
		if lastDigit != withoutLast%10 {
			return 0
		}

		lastDigit = withoutLast % 10
		withoutLast = withoutLast / 10
	}

	return n
}

func Pow(a, b int) int {
	res := 1

	for b > 0 {
		res *= a
		b--
	}

	return res
}

func Mod(a, b int) int {
	return (a%b + b) % b
}
