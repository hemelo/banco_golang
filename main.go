package main

import(
	"fmt"
	"banco/contas"
	"banco/clientes"
)

func testaCriacaoStructs() { 
	henrique := clientes.Titular{Nome: "Henrique", CPF: "123.456.768-90", Profissao: "Dev."}

	contaDoHenrique := contas.ContaCorrente{
		Titular: henrique,
		NumeroAgencia: 589,
		NumeroConta: 123456,
	}

	marcela := clientes.Titular{Nome: "Marcela", CPF: "123.456.768-90", Profissao: "Professora"}

	contaDaMarcela := contas.ContaCorrente{
		Titular: marcela,
		NumeroAgencia:589,
		NumeroConta: 123457,
	}

	var contaDaCris *contas.ContaCorrente
	contaDaCris = new(contas.ContaCorrente)
	contaDaCris.Titular.Nome = "Cris"
	contaDaCris.NumeroConta = 123458
	
	fmt.Println(contaDoHenrique, contaDaMarcela, *contaDaCris)
}

func testaFuncoesBancarias() { 
	henrique := clientes.Titular{Nome: "Henrique", CPF: "123.456.768-90", Profissao: "Dev."}

	contaDoHenrique := contas.ContaCorrente{
		Titular: henrique,
		NumeroAgencia: 589,
		NumeroConta: 123456,
	}

	marcela := clientes.Titular{Nome: "Marcela", CPF: "123.456.768-90", Profissao: "Professora"}

	contaDaMarcela := contas.ContaCorrente{
		Titular: marcela,
		NumeroAgencia: 589,
		NumeroConta: 123456,
	}
	contaDoHenrique.Depositar(1000)
	contaDaMarcela.Depositar(2000)
	fmt.Println(contaDoHenrique.Sacar(500))
	fmt.Println(contaDoHenrique.Depositar(200))
	fmt.Println(contaDoHenrique.Transferir(200, &contaDaMarcela))

	PagarBoleto(&contaDoHenrique, 60)
}

type verificarConta interface{
	Sacar(valor float64) string
}

func PagarBoleto(conta verificarConta, valorBoleto float64){
	conta.Sacar(valorBoleto)	
}

func main() { 
	//testaCriacaoStructs()
	testaFuncoesBancarias()
}