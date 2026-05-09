package reverse

import "testing"

func TestStringASCII(t *testing.T) {
	got := String("abc")
	want := "cba"
	if got != want {
		t.Errorf("String(%q) = %q; want %q", "abc", got, want)
	}
}

func TestStringChinese(t *testing.T) {
	got := String("你好")
	want := "好你"
	if got != want {
		t.Errorf("String(%q) = %q; want %q", "你好", got, want)
	}
}

func TestStringEmpty(t *testing.T) {
	got := String("")
	want := ""
	if got != want {
		t.Errorf("String(%q) = %q; want %q", "", got, want)
	}
}

func TestStringMixed(t *testing.T) {
	got := String("a你b好c")
	want := "c好b你a"
	if got != want {
		t.Errorf("String(%q) = %q; want %q", "a你b好c", got, want)
	}
}
