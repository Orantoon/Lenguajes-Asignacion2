package main

import (
	"asignacion_mod/funciones"
	"fmt"
)

func experimento(n int) {
	/*
		============
		a) Crear un arreglo A de tamaño n, cuyos elementos son números pseudo-aleatorios generados por la función
		que resuelve el ítem 1 (rango de valores 0 .. 53). Los números pseudo-aleatorios deben ser guardados en A
		en el orden que los fue generando su función. Puede usar las tajadas (slices) de Go para esto.
	*/

	A := funciones.RandArray(n, 29, 1, 2048)
	fmt.Println("A: ", A)

	/*
		============
		b) Crear un arreglo Distr de enteros, con índices en el rango 0 .. 53. Inicialice todos los elementos del arreglo
		en 0. Luego, recorra el arreglo A, y cada vez que aparezca un número k ∈ { i : ℕ | 0 ≤ i ≤ 53}, aumente en 1
		el contador del índice k en el arreglo Distr.
	*/

	Distr := make([]int, len(A))

	for i := 0; i < len(A); i++ {
		Distr[i] = A[i] + 1
	}
	fmt.Println("Distr: ", Distr)

	/*
		============
		c) Crear una tajada (slice) de n elementos enteros; llamémosla TS. Insertar los elementos de A, uno a uno,
		mediante su algoritmo de inserción en un arreglo (ítem 3). Recoger estadísticas.
	*/

	TS := []int{}
	iTS := 0

	for i := 0; i < len(A); i++ {
		iTS += funciones.InsertVal(&TS, A[i])
	}
	fmt.Println("TS: ", TS)

	/*
		============
		d) Crear una tajada (slice) con una copia de los n elementos del arreglo A; llamémosla TOS. Ordenar TOS
		usando su implementación del algoritmo de ordenamiento por selección (ítem 4).
	*/

	TOS := A

	funciones.SelectionSort(&TOS)
	fmt.Println("TOS: ", TOS)

	/*
		============
		e) Crear una tajada (slice) con una copia de los n elementos del arreglo A; llamémosla TOQ. Ordenar TOQ
		usando su implementación del algoritmo de ordenamiento Quicksort (ítem 5).
	*/

	TOQ := A

	funciones.QuicksortCall(&TOQ)
	fmt.Println("TOQ: ", TOQ)

	/*
		============
		f) Crear un árbol binario de búsqueda vacío, Abb, congruente con la representación diseñada por su grupo
		(ítem 8). Sobre ese árbol, insertar todos los n elementos del arreglo A, conforme a las indicaciones dadas
		para el ítem 9. Recoger estadísticas.
	*/

	var Abb *funciones.BSTree = new(funciones.BSTree)
	iAbb := 0

	for i := 0; i < len(A); i++ {
		iAbb += Abb.Insert(A[i])
	}

	/*
		============
		g) En relación con el método DSW (ítem 11): crear un árbol binario de búsqueda vacío, AbbDSW.
		Seguidamente, insertar todos los n elementos del arreglo A, conforme a las indicaciones dadas para el ítem
		9. Después de poblar el árbol AbbDSW con esos datos, someterlo a la transformación DSW (ítem 11).
	*/

	// NO SE HACE

	/*
		============
		h) Grafique el arreglo Distr (ítem 2).
	*/

	funciones.ArrChart(Distr) // Se dibuja en resultado.png

	/*
		============
		i) Mediante la función generadora de número pseudo-aleatorios creada como parte de su solución al resolver
		el ítem 1: generar una secuencia de 10000 números enteros pseudo-aleatorios (en el rango 0 .. 53). Para
		cada valor v generado, buscarlo en las estructuras creadas en los pasos c) a g), de esta manera:
	*/

	Arr := funciones.RandArray(10000, 29, 1, 2048)
	v := 0
	bTS, bTOS, bTOQ, bAbb, tmp := 0, 0, 0, 0, 0

	for i := 0; i < len(Arr); i++ {
		v = Arr[i]

		// Buscar v en TS (método del ítem 6). Recoger estadísticas.
		_, tmp = funciones.LinearSearch(TS, v)
		bTS += tmp

		// Buscar v en TOS (método del ítem 7). Recoger estadísticas.
		_, tmp = funciones.BinarySearchCall(TOS, v)
		bTOS += tmp

		// Buscar v en TOQ (método del ítem 7). Recoger estadísticas.
		_, tmp = funciones.BinarySearchCall(TOQ, v)
		bTOQ += tmp

		// Buscar v en Abb (método del ítem 10). Recoger estadísticas.
		_, tmp = Abb.Search(v)
		bAbb += tmp

		// Buscar v en AbbDSW (método del ítem 11). Recoger estadísticas.
		// NO SE HACE
	}

	/*
		============
		j) Al finalizar, obtener estas estadísticas para el punto i), conforme a los alcances correspondientes a su
		grupo:
	*/

	// Altura (máxima) del árbol, para cada uno de los árboles Abb y AbbDSW

	fmt.Println("Altura Máx: ", Abb.MaxHeight(Abb.Root))

	// Densidad del árbol2, para cada uno de los árboles Abb y AbbDSW.

	//Misma que la altura debido a que los nodos del arbol Abb solo tienen hijos derechos.

	// Cantidad total de comparaciones realizadas en las inserciones sobre TS, Abb. (c) (f)

	fmt.Println("Comparaciones en inserciones => TS: ", iTS, " Abb: ", iAbb)

	// Cantidad total de comparaciones realizadas en las búsquedas sobre TS, TOS, TOQ, Abb, AbbDSW. (i)

	fmt.Println("Comparaciones en búsquedas => TS: ", bTS, " TOS: ", bTOS, " TOQ: ", bTOQ, " Abb: ", bAbb)

}

func main() {
	// ===== Experimento "n = 200" =====
	experimento(200)

	// ===== Experimento "n = 400" =====
	//experimento(400)

	// ===== Experimento "n = 600" =====
	//experimento(600)

	// ===== Experimento "n = 800" =====
	//experimento(800)

	// ===== Experimento "n = 1000" =====
	//experimento(1000)
}
