package palindrome

import "testing"

func TestIsPalindromeASCIITrue(t *testing.T) {
	if !IsPalindrome("abcba") {
		t.Errorf("IsPalindrome(%q) = false; want true", "abcba")
	}
}

func TestIsPalindromeASCIIFalse(t *testing.T) {
	if IsPalindrome("hello") {
		t.Errorf("IsPalindrome(%q) = true; want false", "hello")
	}
}

func TestIsPalindromeChinese(t *testing.T) {
	if !IsPalindrome("上海海上") {
		t.Errorf("IsPalindrome(%q) = false; want true", "上海海上")
	}
}

func TestIsPalindromeEmpty(t *testing.T) {
	if !IsPalindrome("") {
		t.Errorf("IsPalindrome(%q) = false; want true", "")
	}
}

func TestIsPalindromeSingleRune(t *testing.T) {
	if !IsPalindrome("x") {
		t.Errorf("IsPalindrome(%q) = false; want true", "x")
	}
}

func TestIsPalindromeEmoji(t *testing.T) {
	if !IsPalindrome("🙂a🙂") {
		t.Errorf("IsPalindrome(%q) = false; want true", "🙂a🙂")
	}
}
