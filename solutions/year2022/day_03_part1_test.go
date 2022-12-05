package year2022

import (
	"testing"
)

func TestDay_03_Part1(t *testing.T) {
	const correct = 157
	filePath := constructDataPath("03", true)
	if ans := Day_03_Part1(filePath); ans != correct {
		t.Errorf("❌ Got %d; correct is %d", ans, correct)
	}
}
