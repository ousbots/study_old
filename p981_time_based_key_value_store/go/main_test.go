package time_based_key_value_store

import "testing"

func TestProvided(t *testing.T) {
	tm := Constructor()
	tm.Set("foo", "bar", 1)

	if ans := tm.Get("foo", 1); ans != "bar" {
		t.Fatalf("%s is not bar", ans)
	}

	if ans := tm.Get("foo", 3); ans != "bar" {
		t.Fatalf("%s is not bar", ans)
	}

	tm.Set("foo", "bar2", 4)

	if ans := tm.Get("foo", 4); ans != "bar2" {
		t.Fatalf("%s is not bar2", ans)
	}

	if ans := tm.Get("foo", 5); ans != "bar2" {
		t.Fatalf("%s is not bar2", ans)
	}
}

func TestExtra(t *testing.T) {
	tm := Constructor()
	tm.Set("love", "high", 10)
	tm.Set("love", "low", 20)

	if ans := tm.Get("love", 5); ans != "" {
		t.Fatalf("%s is not [empty string]", ans)
	}

	if ans := tm.Get("love", 10); ans != "high" {
		t.Fatalf("%s is not high", ans)
	}

	if ans := tm.Get("love", 15); ans != "high" {
		t.Fatalf("%s is not high", ans)
	}

	if ans := tm.Get("love", 20); ans != "low" {
		t.Fatalf("%s is not low", ans)
	}

	if ans := tm.Get("love", 25); ans != "low" {
		t.Fatalf("%s is not low", ans)
	}
}
