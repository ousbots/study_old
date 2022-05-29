package p8_string_to_integer_atoi

import "testing"

func TestBasic(t *testing.T) {
	input := "1"
	valid := 1
	if ans := myAtoi(input); ans != valid {
		t.Fatalf("%v is %d not %d", input, valid, ans)
	}

	input = "100"
	valid = 100
	if ans := myAtoi(input); ans != valid {
		t.Fatalf("%v is %d not %d", input, valid, ans)
	}

	input = "99999"
	valid = 99999
	if ans := myAtoi(input); ans != valid {
		t.Fatalf("%v is %d not %d", input, valid, ans)
	}

	input = "2147483648"
	valid = 2147483647
	if ans := myAtoi(input); ans != valid {
		t.Fatalf("%v is %d not %d", input, valid, ans)
	}

	input = "-2147483649"
	valid = -2147483648
	if ans := myAtoi(input); ans != valid {
		t.Fatalf("%v is %d not %d", input, valid, ans)
	}

	input = " 00100"
	valid = 100
	if ans := myAtoi(input); ans != valid {
		t.Fatalf("%v is %d not %d", input, valid, ans)
	}

	input = "-1"
	valid = -1
	if ans := myAtoi(input); ans != valid {
		t.Fatalf("%v is %d not %d", input, valid, ans)
	}

	input = "   1   "
	valid = 1
	if ans := myAtoi(input); ans != valid {
		t.Fatalf("%v is %d not %d", input, valid, ans)
	}

	input = "   +1ab abcd "
	valid = 1
	if ans := myAtoi(input); ans != valid {
		t.Fatalf("%v is %d not %d", input, valid, ans)
	}

	input = "   -1ab abcd "
	valid = -1
	if ans := myAtoi(input); ans != valid {
		t.Fatalf("%v is %d not %d", input, valid, ans)
	}

	input = "  ab abcd "
	valid = 0
	if ans := myAtoi(input); ans != valid {
		t.Fatalf("%v is %d not %d", input, valid, ans)
	}

	input = "+b"
	valid = 0
	if ans := myAtoi(input); ans != valid {
		t.Fatalf("%v is %d not %d", input, valid, ans)
	}

	input = "-a"
	valid = 0
	if ans := myAtoi(input); ans != valid {
		t.Fatalf("%v is %d not %d", input, valid, ans)
	}
}

func TestProvided(t *testing.T) {
	input := "42"
	valid := 42
	if ans := myAtoi(input); ans != valid {
		t.Fatalf("%v is %d not %d", input, valid, ans)
	}

	input = "    -42"
	valid = -42
	if ans := myAtoi(input); ans != valid {
		t.Fatalf("%v is %d not %d", input, valid, ans)
	}

	input = "4193 with words"
	valid = 4193
	if ans := myAtoi(input); ans != valid {
		t.Fatalf("%v is %d not %d", input, valid, ans)
	}
}

func TestExtra(t *testing.T) {
	input := "words and 987"
	valid := 0
	if ans := myAtoi(input); ans != valid {
		t.Fatalf("%v is %d not %d", input, valid, ans)
	}

	input = "-+12"
	valid = 0
	if ans := myAtoi(input); ans != valid {
		t.Fatalf("%v is %d not %d", input, valid, ans)
	}

	input = ""
	valid = 0
	if ans := myAtoi(input); ans != valid {
		t.Fatalf("%v is %d not %d", input, valid, ans)
	}

	input = "00000-42a1234"
	valid = 0
	if ans := myAtoi(input); ans != valid {
		t.Fatalf("%v is %d not %d", input, valid, ans)
	}

	input = "9223372036800400"
	valid = 2147483647
	if ans := myAtoi(input); ans != valid {
		t.Fatalf("%v is %d not %d", input, valid, ans)
	}
}
