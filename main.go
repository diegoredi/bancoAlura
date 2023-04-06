package main

import (
	"bancoAlura/clientes"
	"bancoAlura/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) bool
}

func main() {

	contaDoDiego := contas.ContaCorrente{Titular: clientes.Titular{
		Nome:      "Diego",
		CPF:       "44214590899",
		Profissao: "Escravo",
	}, NumeroAgencia: 234, NumeroConta: 606060}

	contaPoupanca := contas.ContaPoupanca{Titular: clientes.Titular{
		Nome:      "Diego Poupanca",
		CPF:       "44214590899",
		Profissao: "Escravo",
	}, NumeroAgencia: 234, NumeroConta: 606060}

	fmt.Println(contaDoDiego.ObterSaldo())

	contaDoDiego.Depositar(1000)
	contaPoupanca.Depositar(100)

	PagarBoleto(&contaDoDiego, 100)
	PagarBoleto(&contaPoupanca, 20)
	fmt.Println(contaDoDiego.ObterSaldo())
	fmt.Println(contaPoupanca.ObterSaldo())

}
