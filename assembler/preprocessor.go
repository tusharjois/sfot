package assembler

import (
	"strings"
)

func Preprocess(input string) string {
	lines := strings.Split(input, "\n")
	m := make(map[string]string)
	var output string

	for _, line := range lines {
		if index := strings.Index(line, "define"); index >= 0 {
			if len(line) <= index+6 {
				continue
			}

			subline := line[index+6:]
			current := 0

			char := string(subline[current])
			if char != "\t" && char != " " {
				continue
			}

			var toReplace string
			var replaceWith string

			for current < len(subline) {
				char = string(subline[current])
				if char == "\t" || char == " " {
					current++
					char = string(subline[current])
				} else {
					var identifier []byte
					for char != "\t" && char != " " && char != "\n" {
						identifier = append(identifier, char[0])
						current++
						if current >= len(subline) {
							break
						}
						char = string(subline[current])
					}
					if toReplace == "" {
						toReplace = string(identifier)
					} else {
						replaceWith = string(identifier)
						break
					}

				}
			}

			m[toReplace] = replaceWith
		} else {
			output += line
			output += "\n"
		}
	}
	for key, value := range m {
		output = strings.Replace(output, key, value, -1)
	}

	return output
}
