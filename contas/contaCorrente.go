package contas

import "banco/clientes"

type ContaCorrente struct { 
	Titular clientes.Titular 
	NumeroAgencia int
	NumeroConta int
	saldo float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string { 
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) { 
	podeDepositar := valorDoDeposito > 0
	if podeDepositar {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso", c.saldo
	} else { 
		return "Valor inserido é negativo", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) string {
	podeTransferir := valorDaTransferencia > 0 && valorDaTransferencia <= c.saldo
	
	if podeTransferir {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return "Transferência realizada com sucesso"
	} else {
		return "Não foi possível realizar transferência"
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo;
}