package main

import (
	"fmt"
	"net/http"
)

type Carro struct {
	Nome   string `json:"nome"`
	Modelo string `json:"modelo"`
	Ano    int    `json:"-"`
}

func (c Carro) Andar() {
	fmt.Println("O carro " + c.Nome + " esta andando")
}

func (c Carro) Parar() {
	fmt.Println("O carro " + c.Nome + " esta parado")
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func main() {
	var b int
	b = 22
	fmt.Println(b)
	print("Hello World")
	var c, d string = "hello", " world"
	fmt.Println(c, d)
	a := 1
	print(a)

	carro1 := Carro{"Ford", "Mustang", 1969}

	fmt.Println(carro1.Nome)
	carro1.Andar()
	carro1.Parar()

	http.HandleFunc("/", home)

	http.ListenAndServe("localhost:8080", nil)

}
