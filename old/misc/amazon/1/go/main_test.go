package amazon

import "testing"

func TestBasic(t *testing.T) {
	input := []string{"bag", "sfe", "cbh", "cbh", "red"}
	valid := int32(3)
	if ans := countFamilyLogins(input); ans != valid {
		t.Fatalf("%v is not %v", ans, valid)
	}

}
