package main

import "fmt"

type contaCorrente struct {
	titular     string
	agencia     int
	numeroConta int
	saldo       float64
}

func main() {

	contaDoDiego := contaCorrente{"diego", 589, 123456, 10.0}

	contaDaBruna := contaCorrente{"bruna", 589, 654321, 8.0}

	fmt.Println(contaDoDiego)
	fmt.Println(contaDaBruna)

}
