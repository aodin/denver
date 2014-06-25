package parsing

import (
	"testing"
)

// TODO These functions would be perfect for table-based tests
func TestConcatIfNotEmpty(t *testing.T) {
	x := "Hello World"
	output := ConcatIfNotEmpty(" ", "Hello", "", "World")
	if x != output {
		t.Errorf("Unexpected output from ConcatIfNotEmpty: %s", output)
	}
}

func TestParseIntOrDefault(t *testing.T) {
	out := ParseIntOrDefault("4", 0)
	if out != 4 {
		t.Errorf("Unexpected output from ParseIntOrDefault: %s", out)
	}
	out = ParseIntOrDefault("", 5)
	if out != 5 {
		t.Errorf("Unexpected output from ParseIntOrDefault default: %s", out)
	}
}

func TestParseYesOrNo(t *testing.T) {
	if !ParseYesOrNo("YES") {
		t.Errorf("ParseYesOrNo returned false when it should return true")
	}
	if !ParseYesOrNo("yes") {
		t.Errorf("ParseYesOrNo returned false when it should return true")
	}
}
