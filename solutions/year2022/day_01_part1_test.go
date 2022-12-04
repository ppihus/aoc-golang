package year2022

import (
	"testing"
)

func TestDay_01_Part1(t *testing.T) {
	const correct = 24000
	filePath := constructDataPath("01", true)

	if ans := Day_01_Part1(filePath); ans != correct {
		t.Errorf("❌ Got %d; correct is %d", ans, correct)
	}
}
