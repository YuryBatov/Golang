package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type files struct {
	name     string
	quantity int
}

func countword(str []string, word []string) int {
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
func readfile(files_name string) []string {
	var str1 []string
	data, err := ioutil.ReadFile(files_name)
	if err != nil {
		fmt.Println("Ошибка чтения файла")
		fmt.Println(err)
	}
	str := string(data)
	str1 = strings.Split(str, " ")
	return str1
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
	fmt.Scan(&word)
	words = strings.Split(word, " ")
	for i := range files_name {
		data.name = files_name[i]
		str1 := readfile(files_name[i])
		data.quantity = countword(str1, words)
		data1 = append(data1, data)
	}
	for i := range files_name {

	}
	//fmt.Printf("%s ", str1)

}
