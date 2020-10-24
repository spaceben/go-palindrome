package palindrome

import "errors"

func reverse(n int) (rv int) {
	for n > 0 {
		remainder := n % 10
		rv *= 10
		rv += remainder
		n /= 10
	}
	return
}

// Palindrome return true if an integeres is palindrome (aka the same as read
// from left to right as well as right to left). Example: 12021, 111303111, 0
func Palindrome(a int) (bool, error) {
	if a < 0 {
		return false, errors.New("cannot accept negative integers")
	}
	return a == reverse(a), nil
}
