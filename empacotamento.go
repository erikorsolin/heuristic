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


// funcao que cria uma nova caixa com os parâmetros passados
func NovaCaixa(largura, altura, profundidade float64) *caixa {
		return &caixa{  		// o operador & serve para retornar o endereço de memória dessa variável
			largura: largura,
			altura: altura,
			profundidade: profundidade,
			items: make([]*item, 0),    

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

