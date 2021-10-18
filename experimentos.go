package main

import (
	"asignacion_mod/funciones"
	"fmt"
)


func main(){

	// ===== Experimento "n = 200" =====

	// a)
	A := funciones.RandArray(200, 29, 1, 2048)
	//fmt.Println(A)

	// b)
	Distr := make([]int, 3)

	//funciones.ArrChart(funciones.RandArray(500, 29, 1, 2048))
}
