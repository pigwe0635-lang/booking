package main

func FirstWord(s string) string {
	firstword := ""
	for _, CH := range s {
		if CH == ' ' {
			break
		}
		firstword += string(CH)
	}
	return firstword + "\n"
}
