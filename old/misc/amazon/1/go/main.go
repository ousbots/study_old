package amazon

import "fmt"

func countFamilyLogins(logins []string) int32 {
	count := int32(0)

	for i := range logins {
		check := logins[i]
		checkRot := rotate(check)

		for j := i + 1; j < len(logins); j++ {
			if check == rotate(logins[j]) {
				count++
				continue
			} else if checkRot == logins[j] {
				count++
				continue
			}

			fmt.Printf("%s and %s do not match\n", check, logins[j])
		}
	}

	return count
}

func rotate(login string) string {
	output := []byte{}

	for _, letter := range login {
		if letter == 'z' {
			output = append(output, 'a')
		} else {
			output = append(output, byte(letter)+1)
		}
	}

	return string(output)
}
