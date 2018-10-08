package main

import "fmt"
import "io/ioutil"

func main() {
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Ошибка чтения файла")
	}
	int(data)
	for i=1;i<n;i++ {
		for j=i;j>0 && data[j-1]>data[j];j-- {
							a = data[i]
				data[i]=data[j]
				data[j]=a	
			}
	}
	fmt.Print(string(data))
}
