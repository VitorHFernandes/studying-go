package formas

import (
	"math"
)

type Forma interface {
	Area() float64
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

type Retangulo struct {
	Altura  float64
	Largura float64
}

type Circulo struct {
	raio float64
}
