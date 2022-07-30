package p721_accounts_merge

import (
	"sort"
	"testing"
)

func TestProvided(t *testing.T) {
	input := [][]string{{"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"John", "johnsmith@mail.com", "john00@mail.com"}, {"Mary", "mary@mail.com"}, {"John", "johnnybravo@mail.com"}}
	valid := [][]string{{"John", "johnsmith@mail.com", "john_newyork@mail.com", "john00@mail.com"}, {"Mary", "mary@mail.com"}, {"John", "johnnybravo@mail.com"}}
	for index := range valid {
		sort.Slice(valid[index], func(i, j int) bool { return valid[index][i] < valid[index][j] })
	}
	if ans := accountsMerge(input); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}

	input = [][]string{{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe1@m.co"}, {"Kevin", "Kevin3@m.co", "Kevin5@m.co", "Kevin0@m.co"}, {"Ethan", "Ethan5@m.co", "Ethan4@m.co", "Ethan0@m.co"}, {"Hanzo", "Hanzo3@m.co", "Hanzo1@m.co", "Hanzo0@m.co"}, {"Fern", "Fern5@m.co", "Fern1@m.co", "Fern0@m.co"}}
	valid = [][]string{{"Ethan", "Ethan0@m.co", "Ethan4@m.co", "Ethan5@m.co"}, {"Gabe", "Gabe0@m.co", "Gabe1@m.co", "Gabe3@m.co"}, {"Hanzo", "Hanzo0@m.co", "Hanzo1@m.co", "Hanzo3@m.co"}, {"Kevin", "Kevin0@m.co", "Kevin3@m.co", "Kevin5@m.co"}, {"Fern", "Fern0@m.co", "Fern1@m.co", "Fern5@m.co"}}
	for index := range valid {
		sort.Slice(valid[index], func(i, j int) bool { return valid[index][i] < valid[index][j] })
	}
	if ans := accountsMerge(input); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}
}

func TestBasic(t *testing.T) {
	input := [][]string{{"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"John", "johnsmith@mail.com", "john00@mail.com"}, {"John", "joop@mailoo.com", "john_newyork@mail.com"}, {"John", "john00@mail.com"}, {"Mary", "mary@mail.com"}, {"John", "johnnybravo@mail.com"}}
	valid := [][]string{{"John", "johnsmith@mail.com", "joop@mailoo.com", "john_newyork@mail.com", "john00@mail.com"}, {"Mary", "mary@mail.com"}, {"John", "johnnybravo@mail.com"}}
	for index := range valid {
		sort.Slice(valid[index], func(i, j int) bool { return valid[index][i] < valid[index][j] })
	}
	if ans := accountsMerge(input); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}
}

func TestExtra(t *testing.T) {
	input := [][]string{{"David", "David0@m.co", "David4@m.co", "David3@m.co"}, {"David", "David5@m.co", "David5@m.co", "David0@m.co"}, {"David", "David1@m.co", "David4@m.co", "David0@m.co"}, {"David", "David0@m.co", "David1@m.co", "David3@m.co"}, {"David", "David4@m.co", "David1@m.co", "David3@m.co"}}
	valid := [][]string{{"David", "David0@m.co", "David1@m.co", "David3@m.co", "David4@m.co", "David5@m.co"}}
	for index := range valid {
		sort.Slice(valid[index], func(i, j int) bool { return valid[index][i] < valid[index][j] })
	}
	if ans := accountsMerge(input); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}

	input = [][]string{{"David", "David0@m.co", "David1@m.co"}, {"David", "David3@m.co", "David4@m.co"}, {"David", "David4@m.co", "David5@m.co"}, {"David", "David2@m.co", "David3@m.co"}, {"David", "David1@m.co", "David2@m.co"}}
	valid = [][]string{{"David", "David0@m.co", "David1@m.co", "David2@m.co", "David3@m.co", "David4@m.co", "David5@m.co"}}
	for index := range valid {
		sort.Slice(valid[index], func(i, j int) bool { return valid[index][i] < valid[index][j] })
	}
	if ans := accountsMerge(input); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}
}

func comp(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	for _, curr := range a {
		found := false
	check:
		for _, tmp := range b {
			if len(curr) == len(tmp) {
				for j := range curr {
					if curr[j] != tmp[j] {
						continue check
					}
				}
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}