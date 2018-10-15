package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"./countword"
	"./readfile"
)

type files struct {
	name     string
	quantity int
}

/**type revers_index struct{
	name_word string
	name_files []string
}*/

func sort(data1 []files, n int) []files {
	for i := 1; i < n; i++ {
		for j := i; j > 0 && data1[j-1].quantity < data1[j].quantity; j-- {
			data1[j-1], data1[j] = data1[j], data1[j-1]
		}
	}
	return data1
}

func main() {
	var word string
	var words []string
	var data files
	var data1 []files
	files_name := make([]string, 0, 10)
	files_name = os.Args[1:]
	fmt.Printf("%s", files_name)
	fmt.Printf("\n%s", "Введите слово: ")
	buf := bufio.NewScanner(os.Stdin)
	buf.Scan()
	word = buf.Text()
	words = strings.Split(word, " ")
	for i := range files_name {
		data.name = files_name[i]
		str1 := readfile.Readfile(files_name[i])
		data.quantity = countword.Countword(str1, words)
		data1 = append(data1, data)
	}
	data1 = sort(data1, len(files_name))
	for i := range files_name {
		if data1[i].quantity != 0 {
			fmt.Printf("\n%s; совпадений - %d", data1[i].name, data1[i].quantity)
		}
	}
}
