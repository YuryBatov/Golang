package main

import (
	"bufio"
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
	//fmt.Println(str)
	str1 = strings.Split(str, " ")
	//fmt.Println(str1[2])
	return str1
}

//func sort1(data []files, len1 int) []files {

//}

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
	fmt.Println(words)
	for i := range files_name {
		data.name = files_name[i]
		str1 := readfile(files_name[i])
		data.quantity = countword(str1, words)
		data1 = append(data1, data)
	}
	for i := 1; i < len(files_name); i++ {
		for j := i; j > 0 && data1[j-1].quantity < data1[j].quantity; j-- {
			data1[j-1], data1[j] = data1[j], data1[j-1]
		}
	}
	for i := range files_name {
		if data1[i].quantity != 0 {
			fmt.Printf("\n%s; совпадений - %d", data1[i].name, data1[i].quantity)
		}
	}
}
