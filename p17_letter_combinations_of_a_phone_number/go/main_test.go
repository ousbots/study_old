package p17_letter_combinations_of_a_phone_number

import (
	"sort"
	"testing"
)

func TestBasic(t *testing.T) {
	valid := []string{}
	if ans := letterCombinations(""); !comp(ans, valid) {
		t.Fatalf("%#v is not %#v", ans, valid)
	}

	valid = []string{"p", "q", "r", "s"}
	if ans := letterCombinations("7"); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}

	valid = []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	if ans := letterCombinations("23"); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}

	valid = []string{"pw", "px", "py", "pz", "qw", "qx", "qy", "qz", "rw", "rx", "ry", "rz", "sw", "sx", "sy", "sz"}
	if ans := letterCombinations("79"); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}
}

func comp(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Strings(a)
	sort.Strings(b)

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
