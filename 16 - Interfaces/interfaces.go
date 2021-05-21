package main

import (
	"fmt"
	"math"
)

//flexibilidade com os tipos
//

//interface só tem assinaturas de métodos como eles devem ser

//ela so tem assinatura de método
type forma interface {
	//método que chama area que retorna um float64
	area() float64
}

//funcao que esta recebendo uma interface que pode ser implementada por varios tipos diferentes
//f recebe a interface forma
func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}

//struct retangulo
//       circulo

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	//return math.Pi * (c.raio * c.raio)
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {
	r := retangulo{10, 15}
	escreverArea(r) //nao pode usar o retangulo como tipo forma por nao atender os requisitos para ser uma forma.

	c := circulo{10}
	escreverArea(c)
}
