package main

import (
	"fmt"
	//"io/ioutil"
	//"strconv"
	//"strings"
)

//func readfile()

func main() {
	var files_names string
	files_name := make([]string, 0, 10)
	var n int
	fmt.Println("Введите количество файлов: ")
	fmt.Scan(&n)
	//n := 10
	fmt.Printf("Введите названия этих файлов:")
	for i := 1; i <= n; i++ {
		fmt.Scan(&files_names)
		files_name = append(files_name, files_names)

	}
	fmt.Printf("%s", files_name)
	//files_name = append(files_name, "12")
	//fmt.Println("Введите названия файлов")
	//fmt.Printf("%s %d %d", files_name, len(files_name), cap(files_name))
	/**data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Ошибка чтения файла")
		fmt.Println(err)
	}
	str := string(data)
	str1 := strings.Split(str, " ")
	var n []int
	n = make([]int, len(str1))
	for i := 0; i < len(str1); i++ {
		n[i], err = strconv.Atoi(str1[i])
	}
	if err != nil {
		fmt.Println("Ошибка чего то там")
		fmt.Println(err)
	}*/
}
