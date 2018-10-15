package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type files struct {
	name      string
	name_file []string
	quantity  []int
}

/**type revers_index struct{
	name_word string
	name_files []string
}*/
func Readfile(files_name string) []string {
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
func Countword(str []string, word string) int {
	var n int
	for j := range str {
		if word == str[j] {
			n++
		}
	}
	return n
}
func main() {
	var word string
	var words []string
	var files_name []string
	var f_n []string
	var qu []int
	//var data files
	var data1 []files
	//files_name := make([]string, 0, 10)
	files_name = os.Args[1:]
	fmt.Printf("%s", files_name)
	fmt.Printf("\n%s", "Введите слово: ")
	buf := bufio.NewScanner(os.Stdin)
	buf.Scan()
	word = buf.Text()
	words = strings.Split(word, " ")
	for i := range words {
		data1[i].name = words[i]
		k := 0
		for j := range files_name {
			str := Readfile(files_name[j])
			n := Countword(str, words[i])
			if n != 0 {
				f_n[k] = files_name[j]
				qu[k] = n
				k++
			}
		}
		data1[i].name_file = f_n
		data1[i].quantity = qu
	}
	fmt.Println(data1)
}
