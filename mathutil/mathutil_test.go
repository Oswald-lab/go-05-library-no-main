package mathutil

import "testing"

func TestGCD(t *testing.T) {
	if got := GCD(54, 24); got != 6 {
		t.Fatalf("GCD = %d, want 6", got)
	}
}

func TestIsPrime(t *testing.T) {
	if !IsPrime(13) || IsPrime(15) {
		t.Fatal("IsPrime wrong")
	}
}
