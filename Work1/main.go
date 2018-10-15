package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Ошибка чтения файла")
		fmt.Println(err)
		//os.Exit
	}
	str := string(data)
	str1 := strings.Split(str, " ")
	var n []int
	n = make([]int, len(str1))
	for i := 0; i < len(str1); i++ {
		n[i], err = strconv.Atoi(str1[i])
		if err != nil {
			fmt.Println("Ошибка")
			fmt.Println(err)
			//os.Exit
		}
	}
	//n, err := strconv.Atoi(str1)

	for i := 0; i < len(n); i++ {
		for j := i; j > 0 && n[j-1] > n[j]; j-- {
			n[j-1], n[j] = n[j], n[j-1]
		}
	}
	fmt.Print(n)
	//fmt.Print(strconv.Atoi(str1[1]))
}
