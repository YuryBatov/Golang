package main

import (
	"fmt"
	"os"
	//"io/ioutil"
	//"strconv"
	//"strings"
)

//func readfile()

func main() {
	var word string
	files_name := make([]string, 0, 10)
	files_name = os.Args[1:]

	fmt.Printf("%s", files_name)
	fmt.Scan(&word)
	fmt.Printf("%s", word)
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
