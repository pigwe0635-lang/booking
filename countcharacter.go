package main

func CountChar(s string, c rune) int {
	n := 0
	for _, r := range s {
		if r == c {
			n++
		}
	}
	return n
}
