package contas

import "bancoAlura/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (conta *ContaPoupanca) Sacar(valorDoSaque float64) bool {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= conta.saldo

	if podeSacar {
		conta.saldo -= valorDoSaque
		return true
	} else {
		return false
	}

}

func (conta *ContaPoupanca) Depositar(valorDeposito float64) bool {

	if valorDeposito > 0 {
		conta.saldo += valorDeposito
		return true
	} else {
		return false
	}
}

func (conta *ContaPoupanca) Transferencia(valorTransferencia float64, contaDestino *ContaPoupanca) bool {

	if conta.Sacar(valorTransferencia) {
		contaDestino.Depositar(valorTransferencia)
		return true
	} else {
		return false
	}
}

func (conta *ContaPoupanca) ObterSaldo() float64 {
	return conta.saldo
}
