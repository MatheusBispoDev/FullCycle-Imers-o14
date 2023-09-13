package main

import (
	"fmt"
	"time"
)

// T1
func main() {
	channel := make(chan int)

	for i := 0; i < 10; i++ {
		go worker(i, channel) //Thread 2 ficara ouvindo as altearcoes, sempre que o canal for preenchido, ela lera ele
	}

	for i := 0; i < 10; i++ {
		channel <- i // atribuindo valor no canal, so e possivel atribuir quando o canal estiver vazio. E assim que ele for lido, ele ser esvaziado
	}

	fmt.Println(<-channel)
}

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

/*
func threads() {
	// go routine - Cria uma thread gerenciada pelo RunTime da Go ao inves de usar o Sistema Operacional para gerenciar, ganhando performance
	// Ex: 1 thread somente no SO consome 1MG em media, com o gerenciamento do Go 1 thread consome 2kbs de memoria
	go counter() // Thread 1
	go counter() // Thread 2
	go counter() // Thread 3
}

func counter() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}


type Course struct {
	Name        string `json:"course"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func (c Course) GetFullInfo() string {
	return fmt.Sprintf("Name: %s, Description: %s, Price: %d", c.Name, c.Description, c.Price)
}

func structs() {
	course := Course{
		Name:        "Golang",
		Description: "Golang Course",
		Price:       100,
	}

	fmt.Println(course.GetFullInfo())
}

func webapi() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	course := Course{
		Name:        "Golang",
		Description: "Golang Course",
		Price:       100,
	}

	json.NewEncoder(w).Encode(course)
	w.Write([]byte("Hello World Web"))
}

func fmtBiblioteca() {
	a := 10
	fmt.Println("Hello World!", a)
}

func entendoGo() {
	a := "Hello World!" //:= significa que estou declarando e atribuindo um valor
	println(a)

	println(soma(1, 2))

	resultado, status := soma(20, 20)

	println(resultado)
	println(status)
}

func soma(x int, y int) (int, bool) {
	if x > 10 {
		return x + y, true
	}

	return x + y, false
}
*/
