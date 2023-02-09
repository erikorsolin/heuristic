package main 

import (
	"fmt"
	"math"
	"sort"
)

// onde os itens serão colocados 
type caixa struct{
	largura float64
	altura float64
	profundidade float64
	items []*item // items que foram empacotados 

}

type RotationType int

type Axis int

type pivo [3]float64

type Dimension [3]float64

type item struct{
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

// Gets
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


// a funcao PutItem tenta pôr o item no pivo p da caixa c
func (c *caixa) PutItem(item *item, p pivo) (fit bool){




}














