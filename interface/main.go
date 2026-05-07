package main

import "strconv"

type Carro struct {
	Modelo string
	Ano int
}

type Veiculo interface{
	Andar()
}

func main() {
	meuCarro := Carro{Modelo: "Fusca", Ano: 1970}
	CarroAndando(meuCarro)
}

func (c Carro) Andar() {
	println("O carro " + c.Modelo + " do ano " + strconv.Itoa(c.Ano) + " está andando!")
}

func CarroAndando (c Carro) {
	c.Andar()
}
