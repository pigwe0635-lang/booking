package main

func cameltosnake(s string) string {
	res := ""
	for i, r := range s {
		if r >= 'A' && r <= 'Z' {
			if i > 0 {
				res += "_"
			}
			res += string(r + 32)
		} else {
			res += string(r)
		}
	}
	return res
}
