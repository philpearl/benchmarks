package benchmarks

func keyExtract(key string, separator rune, field int) string {
	seps := 0
	start := -1
	if field == 0 {
		start = 0
	}
	var last rune
	for i, c := range key {
		if last == separator {
			seps++
			if seps == field {
				start = i
			}
		}
		if c == separator {
			if start != -1 {
				return key[start:i]
			}
		}
		last = c
	}

	if start == -1 {
		return ""
	}

	return key[start:]
}
