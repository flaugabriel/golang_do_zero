package formas

import (
	"fmt"
	"math"
)

type Retangulo struct {
	Altura  float64
	Largura float64
}

type Circulo struct {
	Raio float64
}

type Forma interface {
	Area() float64
}

// pode ser implementado por varios tipos de situações diferentes
func EscreverArea(f Forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.Area())
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}

func main() {
	r := Retangulo{10, 15}
	EscreverArea(r)
	c := Circulo{10}
	EscreverArea(c)
}
