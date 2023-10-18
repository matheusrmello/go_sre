package main

import (
	"fmt"
)

type Pessoa struct {
	nome    string
	idade   int
	salario int
}

func main() {
	//forma menos usual de usar a struct
	pessoa2 := new(Pessoa)
	pessoa2.nome = "Matheus Mello"
	pessoa2.idade = 26
	pessoa2.salario = 200

	//forma mais comum de usar a struct
	pessoa := &Pessoa{
		nome: "Matheus",
		idade: 25,
		salario: 100,
	}
	pessoa.addSalaryPessoa(10)
	fmt.Println(pessoa.salario)

	// name, salario := "Matheus", 100
	// setName(name)
	// getName()
	// newSalary, bonus := addSalary(salario, 10)
	// fmt.Println("Novo salario:", newSalary)
	// fmt.Println("Bonus:", bonus)

}

func (p *Pessoa)addSalaryPessoa(bonus int)  {
	p.salario += bonus
}

func getName() string {
	return "Matheus Mello"
}

func setName(name string) {
	fmt.Println("Hello", name)
}

func addSalary(valorAtual int, bonus int) (int, int) {
	return valorAtual + bonus, bonus
}
