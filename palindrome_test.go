package palindrome

import "testing"

type tt struct {
	a         int
	isPalin   bool
	shouldErr bool
}

func TestPalindrome(t *testing.T) {
	table := []tt{
		{
			0,
			true,
			false,
		},
		{
			10,
			false,
			false,
		},
		{
			112211,
			true,
			false,
		},
		{
			332233,
			true,
			false,
		},
		{
			123,
			false,
			false,
		},
		{
			1234567654321,
			true,
			false,
		},
		{
			123456777654321,
			true,
			false,
		},
		{
			-12321,
			false,
			true,
		},
		{
			-123421,
			false,
			true,
		},
	}
	for _, i := range table {
		res, err := Palindrome(i.a)
		if err != nil && i.shouldErr {
			t.Log("Correctly errored with:", err)
			continue
		}
		if err != nil {
			t.Error("Errored with:", err)
		}

		if res == i.isPalin {
			t.Log("Correct result for:", i.a)
			continue
		}

		t.Error("Wrong result for:", i.a)
	}
}
