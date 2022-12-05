package year2022

import (
	"testing"
)

func TestDay_03_Part2(t *testing.T) {
	const correct = 70
	filePath := constructDataPath("03", true)
	if ans := Day_03_Part2(filePath); ans != correct {
		t.Errorf("❌ Got %d; correct is %d", ans, correct)
	}
}
