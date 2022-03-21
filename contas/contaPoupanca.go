package contas

import "banco/clientes"

type ContaPoupanca struct { 
	Titular clientes.Titular 
	NumeroAgencia, NumeroConta, Operacao int
	saldo float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string { 
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) { 
	podeDepositar := valorDoDeposito > 0
	if podeDepositar {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso", c.saldo
	} else { 
		return "Valor inserido é negativo", c.saldo
	}
}

func (c *ContaPoupanca) Transferir(valorDaTransferencia float64, contaDestino *ContaPoupanca) string {
	podeTransferir := valorDaTransferencia > 0 && valorDaTransferencia <= c.saldo
	
	if podeTransferir {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return "Transferência realizada com sucesso"
	} else {
		return "Não foi possível realizar transferência"
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo;
}