package smart

// Ported from
// http://stackoverflow.com/questions/667803/what-is-the-best-algorithm-for-arbitrary-delimiter-escape-character-processing

const (
	stateParsing = false
	stateEscaped = true
)

func Split(input, delimiter, escape string) []string {
	state := stateParsing
	found := make([]string, 0)
	parsed := ""

	for _, c := range input {
		c := string(c)
		if state == stateParsing {
			if c == delimiter {
				found = append(found, parsed)
				parsed = ""
			} else if c == escape {
				state = stateEscaped
			} else {
				parsed += c
			}
		} else {
			parsed += c
			state = stateParsing
		}
	}

	if parsed != "" {
		found = append(found, parsed)
	}

	return found
}
