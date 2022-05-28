package p721_accounts_merge

import "sort"

func accountsMerge(accounts [][]string) [][]string {
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
