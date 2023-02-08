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

