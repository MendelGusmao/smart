package smart

// Ported from
// http://stackoverflow.com/questions/667803/what-is-the-best-algorithm-for-arbitrary-delimiter-escape-character-processing

func Split(input, delimiter, escape string) []string {
	parsing := false
	escaped := true

	state := parsing
	found := make([]string, 0)
	parsed := ""

	for _, c := range input {
		c := string(c)
		if state == parsing {
			if c == delimiter {
				found = append(found, parsed)
				parsed = ""
			} else if c == escape {
				state = escaped
			} else {
				parsed += c
			}
		} else {
			parsed += c
			state = parsing
		}
	}

	if parsed != "" {
		found = append(found, parsed)
	}

	return found
}
