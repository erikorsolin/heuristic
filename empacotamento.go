package main 

import (
	"fmt"
	"math"
	"sort"
	"math/rand"
)

type caixa struct{
	largura float64
	altura float64
	profundidade float64
	items []*item 

}

type RotationType int

type Axis int

const (
	WidthAxis Axis = iota
	HeightAxis
	DepthAxis
)

type pivo [3]float64

// x, y, z
var posicao_inicial = pivo{0, 0, 0}

type Dimension [3]float64



type item struct{
	name string
	largura float64
	altura float64
	profundidade float64

	// usado durante packer.Pack()
	RotationType RotationType
	Position pivo
}




// funcao que cria uma nova caixa com os parâmetros passados
func NovaCaixa(largura, altura, profundidade float64) *caixa {
		return &caixa{  		// o operador & serve para retornar o endereço de memória dessa variável
			largura: largura,
			altura: altura,
			profundidade: profundidade,
			items: make([]*item, 0),  // inicializa uma slice com o tipo item de comprimento zero  

		}

}


// Gets da caixa
func (l *caixa) GetLargura() float64{
	return l.largura
}

func (a *caixa) GetAltura() float64{
	return a.altura
}

func (p *caixa) GetProfundidade() float64{
	return p.profundidade
}

func (x *caixa) GetVolume() float64{
	return x.altura * x.largura * x.profundidade
}



// Gets do Item

func (i *item) GetName() string{
	return i.name
}

func (i *item) GetLargura() float64{
	return i.largura
}

func (i *item) GetAltura() float64{
	return i.altura
}

func (i *item) GetProfundidade() float64{
	return i.profundidade
}

func(i *item) GetVolume() float64{
	return i.altura * i.largura * i.profundidade
}



// tenta colocar o item i no pivo p da caixa c
func (c *caixa) ColocarItem(i *item, p pivo) (colocado bool){
	i.Position = p
	if c.largura < p[0]+i.largura || c.altura <  p[1]+i.altura || c.profundidade < p[2]+i.profundidade {
		return
		
	}
	colocado = true

	for _, ib := range c.items {    // verifica se os itens já empacotados não interceptam o novo item
		if ib.intersect(i) {
			colocado = false
			break
		}
	}

	if colocado {
		c.items  = append(c.items, i)
	}

	return
}



// intersect verifica se há uma intersecção entre o item i passado como parâmetro e outro item atual da iteração 

func (i *item) intersect(i2 *item) bool{
	return rectIntersect(i, i2, WidthAxis, HeightAxis) &&
		   rectIntersect(i, i2, HeightAxis, DepthAxis) &&
		   rectIntersect(i, i2, WidthAxis, DepthAxis)
}





