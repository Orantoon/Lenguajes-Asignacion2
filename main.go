package main

import (
	"math/rand"
	"os"

	"github.com/wcharczuk/go-chart"
)

/**
1. Diseñar y construir una función que genere un arreglo de tamaño n con números enteros pseudo-aleatorios
obtenidos mediante el método de congruencia lineal multiplicativa, a partir de una semilla dada. La semilla
deberá ser un número primo entre 11 y 101. Los valores generados deben ser convertidos al intervalo 0 ..
53. El período debe ser ≥ 2048. n puede ser cualquier número en el intervalo 200 .. 1000.
*/

func randArray(numSeed int64) []int {

	rand.Seed(numSeed)

	min := 200
	max := 1000
	n := rand.Intn(max - min + 1) + max
	arr := make([]int, n)



	return arr
}

/**
2. Diseñar y construir una función que genere un gráfico de barras a partir de un arreglo de n números enteros,
cada uno de los cuales pertenece al intervalo 0 .. 53. Las barras deben ser verticales y de tamaño
directamente proporcional al número que representan.
*/

func arrChart() {
	graph := chart.BarChart{
		Title : "Gráfico de Barras",
		Background: chart.Style{
			Padding: chart.Box{
				Top: 30,
			},
		},
		Height : 256,
		BarWidth: 50,
		Bars: []chart.Value{
			{Value: 10, Label: "1"},
			{Value: 11, Label: "2"},
			{Value: 12, Label: "3"},
			{Value: 13, Label: "4"},
			{Value: 14, Label: "5"},
			{Value: 15, Label: "6"},
		},
	}
	f , _ := os.Create("resultados.png")
	defer f.Close()
	graph.Render(chart.PNG,f)
}

/**
3. Diseñar y construir una función que haga la inserción de un valor entero (la llave) en un arreglo de
números enteros. Si la llave no se encuentra en el arreglo, inserta la llave al final del arreglo. Si la llave ya
se encuentra en el arreglo, no la inserta. La función retorna un número entero, que será la cantidad de
comparaciones realizadas, incluida la que llevó a la inserción.
*/

func insertVal(arr *[]int, key int) int {

	for index, value := range *arr {
		if value == key {
			return index + 1 // Found, returns number of comparisons done [index + 1 because it starts at 0]
		}
	}

	len := len(*arr)

	*arr = append(*arr, key)
	return len // Not found, checked all the values in the original array/slice and because it wasn't, it is the length of it.
}

/*
func main(){

	slice1, slice2 := []int{-12,32,99,67,-83,123}, []int{0,1,2,3,4}

	n1, n2 := 13, 2
	three1, three2 := insertVal(&slice1,n1), insertVal(&slice2, n2)

	fmt.Printf("Array 1 = Key: %d, Comparisons done: %d, Final Slice: %v\nArray 2 = Key: %d, Comparisons done: %d, Final Slice: %v", n1, three1, slice1, n2, three2, slice2)
}*/

/**
4. Diseñar y construir una función que ordene ascendentemente un arreglo de enteros mediante el método de
ordenamiento de selección. Este algoritmo es cuadrático en promedio.
*/

func selectionSort(arr *[]int) {
	arr2 := *arr
	len := len(arr2)

	for currentIndex := 0; currentIndex < len-1; currentIndex++ { // Done to all the indexes in the array
		indexMin := currentIndex

		for i := currentIndex + 1; i < len; i++ { // Get the index of the smallest value from the numbers to the right
			if arr2[i] < arr2[indexMin] {
				indexMin = i
			}
		}

		//Swap numbers
		arr2[currentIndex], arr2[indexMin] = arr2[indexMin], arr2[currentIndex]
	}

	*arr = arr2 // Assign changes to original array
}

/*
func main(){
	arr := []int{-12,32,99,67,-83,123}
	selectionSort(&arr)
	fmt.Println(arr)
}*/

/**
5. Diseñar y construir una función que ordene ascendentemente un arreglo de enteros mediante el método de
ordenamiento conocido como Quicksort. Este algoritmo es O(n log2 (n)) en promedio.
*/

func partition(arr *[]int, low int, high int) int { //
	arrC := *arr
	pivot := arrC[high]

	i := low - 1

	for j := low; j < high; j++ {
		if arrC[j] <= pivot {
			i++
			arrC[i], arrC[j] = arrC[j], arrC[i] //Gets the lesser values to the left of the pivot
		}
	}

	//Swap pivot with the next element to i
	arrC[i+1], arrC[high] = arrC[high], arrC[i+1]

	*arr = arrC // Assign changes to original array

	return i + 1 //new pivot
}

func quicksort(arr *[]int, low int, high int) {
	if low < high { //Divides into subarrays by selecting a pivot and organizes them
		pivot := partition(arr, low, high)
		quicksort(arr, low, pivot-1)
		quicksort(arr, pivot+1, high)
	}
}

func quicksortCall(arr *[]int) {
	quicksort(arr, 0, len(*arr)-1)
}

/*
func main(){
	arr := []int{-12,32,99,67,-83,123}
	quicksortCall(&arr)
	fmt.Println(arr)
}*/

/**
6. Diseñar y construir una función que haga la búsqueda de un valor entero (la llave) en un arreglo que no
esté ordenado, mediante el algoritmo de búsqueda secuencial. La función debe retornar un par ordenado:
un valor booleano que indique si encontró la llave buscada (true = la encontró, false = no la encontró),
junto con un entero que indique la cantidad de comparaciones realizadas hasta determinar si la llave está o
no presente en el arreglo.
*/

func linearSearch(arr []int, key int) (bool, int) {

	for index, value := range arr {
		if value == key {
			return true, index + 1 // Found, index of where it was + 1 [because it starts at 0]
		}
	}
	return false, len(arr) // Not Found, the whole length of the array
}

/*
func main(){
	arr := []int{-12,32,99,67,-83,123}
	fmt.Println(linearSearch(arr, 99))
	fmt.Println(linearSearch(arr, 13))
}*/

/**
7. Diseñar y construir una función que haga la búsqueda de un valor entero (la llave) en un arreglo que esté
ordenado, mediante el algoritmo de búsqueda binaria. La función debe retornar un par ordenado: un valor
booleano que indique si encontró la llave buscada (true = la encontró, false = no la encontró), junto con un
entero que indique la cantidad de comparaciones realizadas hasta determinar si la llave está o no presente
en el arreglo.
*/

func binarySearch(arr []int, key int, left int, right int, n int) (bool, int) {
	if left <= right {
		mid := (left + right) / 2 // Mid value between left & right

		//fmt.Println("Mid: ", arr[mid], " Left: ", arr[left], " Right: ", arr[right])

		if arr[mid] == key { // Found
			return true, n
		} else if arr[mid] > key {
			return binarySearch(arr, key, left, mid-1, n+1) // Key is lesser
		} else {
			return binarySearch(arr, key, mid+1, right, n+1) // Key is greater
		}
	}
	return false, n // Not found
}

func binarySearchCall(arr []int, key int) (bool, int) {
	return binarySearch(arr, key, 0, len(arr)-1, 1)
}

/*
func main() {
	arr := []int{1, 5, 10, 12, 16, 23, 29, 32, 35, 40, 49, 56, 57, 70} // in ascendent order

	fmt.Println(binarySearchCall(arr, 56))
	fmt.Println(binarySearchCall(arr, 24))
}
*/
