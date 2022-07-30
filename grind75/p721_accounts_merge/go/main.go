package p721_accounts_merge

import "sort"

type Account struct {
	Name   string
	Parent *Account
	Emails map[string]bool
}

func (a *Account) TopParent() *Account {
	tmp := a
	for tmp.Parent != nil {
		tmp = tmp.Parent
	}

	return tmp
}

func (a *Account) Merge(other *Account) {
	for email := range other.Emails {
		a.Emails[email] = true
	}
}

// Using the disjoint-set merge is faster than DFS, O(log(n)) vs DFS O(n).
func accountsMerge(accounts [][]string) [][]string {
	seen := make(map[string]*Account)
	cache := make(map[*Account]bool)

	for _, line := range accounts {
		account := Account{Name: line[0], Emails: make(map[string]bool, len(line[1:]))}

		var parent *Account
		for _, email := range line[1:] {
			account.Emails[email] = true

			if prev, exists := seen[email]; exists {
				if parent == nil {
					parent = prev.TopParent()
				} else {
					other := prev.TopParent()
					if parent != other {
						other.Parent = parent
					}
				}

			} else {
				seen[email] = &account
			}
		}

		if parent == &account {
			account.Parent = nil
		} else {
			account.Parent = parent
		}

		cache[&account] = true
	}

	parents := make(map[*Account]bool)
	for acc := range cache {
		if acc.Parent == nil {
			parents[acc] = true
		} else {
			acc.TopParent().Merge(acc)
		}
	}

	output := [][]string{}
	for acc := range parents {
		line := []string{acc.Name}
		for email := range acc.Emails {
			line = append(line, email)
		}

		sort.Slice(line[1:], func(i, j int) bool { return line[i+1] < line[j+1] })

		output = append(output, line)
	}

	return output
}

func accountsMergeDFS(accounts [][]string) [][]string {
	emails := make(map[string]map[string]bool)
	names := make(map[string]string)

	for _, line := range accounts {
		name := line[0]

		for i := 1; i < len(line); i++ {
			names[line[i]] = name

			email := line[i]
			if _, exists := emails[email]; !exists {
				emails[email] = make(map[string]bool)
			}

			if i == 1 {
				continue
			}

			emails[email][line[i-1]] = true
			emails[line[i-1]][email] = true
		}
	}

	output := [][]string{}
	visited := make(map[string]bool)

	for key := range emails {
		if visited[key] {
			continue
		}

		que := []string{key}
		addrs := []string{}

		for len(que) != 0 {
			email := que[len(que)-1]
			que = que[:len(que)-1]

			if !visited[email] {
				visited[email] = true
				addrs = append(addrs, email)

				for other := range emails[email] {
					if !visited[other] {
						que = append(que, other)
					}
				}
			}
		}

		sort.Slice(addrs, func(i, j int) bool { return addrs[i] < addrs[j] })
		output = append(output, append([]string{names[key]}, addrs...))
	}

	return output
}
