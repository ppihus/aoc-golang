package year2022

import (
	"testing"
)

func TestDay_02_Part1(t *testing.T) {
	const correct = 15
	filePath := constructDataPath("02", true)

	if ans := Day_02_Part1(filePath); ans != correct {
		t.Errorf("❌ Got %d; correct is %d", ans, correct)
	}
}
