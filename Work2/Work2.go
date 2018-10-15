package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"./readfile"
)

type files struct {
	name     string
	quantity int
}

func sort(m map[string][]files, str string) map[string][]files {
	for i := range m[str] {
		for j := i; j > 0 && m[str][j-1].quantity < m[str][j].quantity; j-- {
			m[str][j-1], m[str][j] = m[str][j], m[str][j-1]
		}
	}
	return m
}
func revers_ind(str []string, files_name string, m map[string][]files) map[string][]files {
	//var data []files
	var data1 files
	data1.name = files_name
	data1.quantity = 1
	//var k int
	for i := range str {
		k := 0
		if m[str[i]] != nil {
			for j := range m[str[i]] {
				if m[str[i]][j].name == files_name {
					m[str[i]][j].quantity++
					k++
				}
			}
			if k == 0 {
				m[str[i]] = append(m[str[i]], data1)
			}
		}
		if m[str[i]] == nil {
			m[str[i]] = append(m[str[i]], data1)
		}
		k = 0
		m = sort(m, str[i])
	}

	return m
}

func main() {
	m := make(map[string][]files)
	var word string
	var words []string
	files_name := make([]string, 0, 10)
	files_name = os.Args[1:]
	fmt.Printf("\n%s", "Введите слово: ")
	for i := range files_name {
		str1 := readfile.Readfile(files_name[i])
		m = revers_ind(str1, files_name[i], m)
	}
	buf := bufio.NewScanner(os.Stdin)
	buf.Scan()
	word = buf.Text()
	words = strings.Split(word, " ")
	for i := range words {
		for j := range m[words[i]] {
			fmt.Printf("%s: совпадений %d\n", m[words[i]][j].name, m[words[i]][j].quantity)
		}
	}
}
