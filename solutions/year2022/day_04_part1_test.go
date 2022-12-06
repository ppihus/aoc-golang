package year2022

import (
	"testing"
)

func TestDay_04_Part1(t *testing.T) {
	const correct = 2
	filePath := constructDataPath("04", true)
	if ans := Day_04_Part1(filePath); ans != correct {
		t.Errorf("❌ Got %d; correct is %d", ans, correct)
	}
}
