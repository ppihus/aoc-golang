package year2022

import (
	"testing"
)

func TestDay_04_Part2(t *testing.T) {
	const correct = 4
	filePath := constructDataPath("04", true)
	if ans := Day_04_Part2(filePath); ans != correct {
		t.Errorf("❌ Got %d; correct is %d", ans, correct)
	}
}
