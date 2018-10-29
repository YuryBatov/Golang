package main

// TrimSpace удаляет указанный символ справа слево от исходного(num, err = strconv.Atoi(string.TrimSpace(i)))
// таймер <- times.C функция AfterFunc , вызывает функцию через какое то время?
// contex нужен для коректного прекращения работы канала
// WaitGroup (wg := &sync.WaitGroup{} )
// Mutex
// atomic

//testEq нужен для тестов, сравнивает два объекта (массивы например)
// http (простейший запрос представлен ниже)
// GET (uri) a/1 HTTP
// Content type: (любые данные) image()
//
// - любой текст -

//Ответ от сервера
// HTTP // 200 (OK) - все нормально
// Server: (что то)
// Set Cookie k=v;

// "net" библиотека для подключения к интернету
// "net/http"
//
import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
