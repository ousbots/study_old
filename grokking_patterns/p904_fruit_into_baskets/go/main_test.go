package fruit_into_baskets

import "testing"

type TestCase struct {
	input  []int
	output int
}

func TestProvided(t *testing.T) {
	cases := []TestCase{
		{input: []int{1, 2, 1}, output: 3},
		{input: []int{0, 1, 2, 2}, output: 3},
		{input: []int{1, 2, 3, 2, 2}, output: 4},
	}

	for _, test := range cases {
		ans := totalFruit(test.input)
		if ans != test.output {
			t.Fatalf("expected %d not %d", test.output, ans)
		}
	}
}

func TestBasic(t *testing.T) {
	cases := []TestCase{
		{input: []int{0}, output: 1},
		{input: []int{0, 1}, output: 2},
		{input: []int{1, 2, 2, 3, 3, 3}, output: 5},
		{input: []int{3, 3, 3, 2, 2, 1}, output: 5},
	}

	for _, test := range cases {
		ans := totalFruit(test.input)
		if ans != test.output {
			t.Fatalf("expected %d not %d", test.output, ans)
		}
	}
}
