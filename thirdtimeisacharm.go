package main

func ThirdTimeIsACharm(str string) string {
	res := ""
	for i, r := range str {
		if (i+1)%3 == 0 {
			res += string(r)
		}
	}
	return res + "\n"
}
