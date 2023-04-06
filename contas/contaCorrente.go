package contas

import "bancoAlura/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (conta *ContaCorrente) Sacar(valorDoSaque float64) bool {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= conta.saldo

	if podeSacar {
		conta.saldo -= valorDoSaque
		return true
	} else {
		return false
	}

}

func (conta *ContaCorrente) Depositar(valorDeposito float64) bool {

	if valorDeposito > 0 {
		conta.saldo += valorDeposito
		return true
	} else {
		return false
	}
}

func (conta *ContaCorrente) Transferencia(valorTransferencia float64, contaDestino *ContaCorrente) bool {

	if conta.Sacar(valorTransferencia) {
		contaDestino.Depositar(valorTransferencia)
		return true
	} else {
		return false
	}
}

func (conta *ContaCorrente) ObterSaldo() float64 {
	return conta.saldo
}
