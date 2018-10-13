package countword

func Countword(str []string, word []string) int {
	var n int
	for i := range word {
		for j := range str {
			if word[i] == str[j] {
				n++
			}
		}
	}
	return n
}
