package util

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRandStringOK(t *testing.T) {
	// arrange
	strLen := 10
	isHex := regexp.MustCompile(`^[A-F0-9]+$`)

	for i := 0; i < 100; i++ {
		t.Run(
			fmt.Sprintf("Test case %d", i),
			func(t *testing.T) {
				// act
				str := RandString(strLen)

				// assert
				if len(str) != strLen {
					t.Errorf("Expected length %d, got %d", len(str), strLen)
				}
				if !isHex.MatchString(str) {
					t.Errorf("Expected valid hexadecimal string, got %s", str)
				}
			},
		)
	}
}

func BenchmarkRandString(b *testing.B) {
	n := 10
	for i := 0; i <= b.N; i++ {
		_ = RandString(n)
	}
}
