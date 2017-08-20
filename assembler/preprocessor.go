package assembler

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func Preprocess(input string) (string, map[string]string, error) {
	lines := strings.Split(input, "\n")
	m := make(map[string]string)
	re := regexp.MustCompile(`"\s*([^"]*)\s*"`)
	var output string // for assembly

	pos := 0

	for pos < len(lines) {
		line := lines[pos]
		if index := strings.Index(line, ".include"); index >= 0 {
			if len(line) <= index+8 {
				continue
			}

			subline := line[index+8:]
			fileName := re.FindString(subline)
			fileName = fileName[1 : len(fileName)-1]

			filedata, err := ioutil.ReadFile(fileName)

			if err != nil {
				return "", m, fmt.Errorf("in opening file \"%v\":\n\t%v", fileName, err) // TODO
			}

			processed, defines, err := Preprocess(string(filedata))

			if err != nil {
				return "", m, fmt.Errorf("in included file \"%v\":\n\t%v", fileName, err) // TODO
			}

			for k, v := range defines {
				if _, prs := m[k]; prs {
					msg := fmt.Sprintf("redefinition of %v in included file \"%v\"", k, fileName)
					return "", m, errors.New(msg)
				} else {
					m[k] = v
				}
			}

			lines = append(lines, "BRK\n") // Separate included subroutines from main
			lines = append(lines, strings.Split(processed, "\n")...)
			pos++

		} else if index := strings.Index(line, ".define"); index >= 0 {
			if len(line) <= index+7 {
				continue
			}

			subline := line[index+7:]
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

			if _, prs := m[toReplace]; prs {
				msg := fmt.Sprintf("redefinition of %v at line %v", toReplace, pos)
				return "", m, errors.New(msg)
			} else {
				m[toReplace] = replaceWith
			}

			pos++

		} else {
			output += line
			output += "\n"
			pos++
		}
	}
	for key, value := range m {
		output = strings.Replace(output, key, value, -1)
	}

	return output, m, nil
}
