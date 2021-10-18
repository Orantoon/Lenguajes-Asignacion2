package funciones

import (
	"fmt"
	"os"

	"github.com/wcharczuk/go-chart"
)

/*
============
1. Diseñar y construir una función que genere un arreglo de tamaño n con números enteros pseudo-aleatorios
obtenidos mediante el método de congruencia lineal multiplicativa, a partir de una semilla dada. La semilla
deberá ser un número primo entre 11 y 101. Los valores generados deben ser convertidos al intervalo 0 ..
53. El período debe ser ≥ 2048. n puede ser cualquier número en el intervalo 200 .. 1000.
*/

func RandArray(n int, seed int, k int, m int) []int {
	
	// Validating "n"
	if n < 200 || n > 1000 {
		fmt.Println("El valor n es incorrecto")
		return nil
	}
	
	// Validating Seed
	if seed < 11 || seed > 101 {
		fmt.Println("El valor de la semilla es incorrecto")
		return nil
	}

	for i := 2; i < seed; i++ {	// Prime Number
		if seed % i == 0 {
			fmt.Println("El valor de la semilla es incorrecto")
			return nil
		}
	}

	// Validating "k"
	if k < 0 {
		fmt.Println("El valor k es incorrecto")
		return nil
	}

	// Validating "m"
	if m < 2048 {
		fmt.Println("El valor m es incorrecto")
		return nil
	}
	
	arr := make([]int, n)
	a := 8*k + 3	// 8k + 5 can also be used
	//first := seed

	for i := 0; i < n; i++ {	// Generating the Array
		num := (a * seed) % m	// Main Algorithm, X = (a * [seed or previous number]) % m
		num = num % 54	// Changed to 0..53

		/*
		if num == first && i != 0{	// We can stop the loop to avoid repeating the pattern
			arr = arr[:i]
			break
		}
		*/

		arr[i] = num
		seed = num	// Seed is now the previous number
	}

	fmt.Println("Resultado: ", arr)

	return arr
}


/*
============
2. Diseñar y construir una función que genere un gráfico de barras a partir de un arreglo de n números enteros,
cada uno de los cuales pertenece al intervalo 0 .. 53. Las barras deben ser verticales y de tamaño
directamente proporcional al número que representan.
*/


func ArrChart(arr []int) {

	arrVal := createValues(arr)
	
	graph := chart.BarChart{
		Title : "Gráfico de Barras",
		Background: chart.Style{
			Padding: chart.Box{
				Top: 30,
			},
		},
		Height : 256,
		BarWidth: 50,

		Bars: arrVal,
	}
	f , _ := os.Create("resultado.png")
	defer f.Close()
	graph.Render(chart.PNG,f)
}

func createValues(arr []int) []chart.Value{	// Creates a Value array from a normal array
	res := []chart.Value{}
	val := chart.Value{}

	for i := 0; i < len(arr); i++{
		val.Value = float64(arr[i])
		res = append(res, val)
	}
	return res
}


/*
============
3. Diseñar y construir una función que haga la inserción de un valor entero (la llave) en un arreglo de
números enteros. Si la llave no se encuentra en el arreglo, inserta la llave al final del arreglo. Si la llave ya
se encuentra en el arreglo, no la inserta. La función retorna un número entero, que será la cantidad de
comparaciones realizadas, incluida la que llevó a la inserción.
*/

func InsertVal(arr *[]int, key int) int {
	for index, value := range *arr {
		if value == key {
			return index + 1 // Found, returns number of comparisons done [index + 1 because it starts at 0]
		}
	}; len := len(*arr)

	*arr = append(*arr, key)
	return len // Not found, checked all the values in the original array/slice and because it wasn't, it is the length of it.
}

/*
func test3(){
	slice1, slice2 := []int{-12,32,99,67,-83,123}, []int{0,1,2,3,4}

	n1, n2 := 13, 2
	three1, three2 := insertVal(&slice1,n1), insertVal(&slice2, n2)

	fmt.Printf("Array 1 = Key: %d, Comparisons done: %d, Final Slice: %v\nArray 2 = Key: %d, Comparisons done: %d, Final Slice: %v", n1, three1, slice1, n2, three2, slice2)
}
*/


/*
============
4. Diseñar y construir una función que ordene ascendentemente un arreglo de enteros mediante el método de
ordenamiento de selección. Este algoritmo es cuadrático en promedio.
*/

func SelectionSort(arr *[]int) {
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
func test4(){
	arr := []int{-12,32,99,67,-83,123}
	selectionSort(&arr)
	fmt.Println(arr)
}
*/


/*
============
5. Diseñar y construir una función que ordene ascendentemente un arreglo de enteros mediante el método de
ordenamiento conocido como Quicksort. Este algoritmo es O(n log2 (n)) en promedio.
*/

func Partition(arr *[]int, low int, high int) int { //
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

func Quicksort(arr *[]int, low int, high int) {
	if low < high { //Divides into subarrays by selecting a pivot and organizes them
		pivot := Partition(arr, low, high)
		Quicksort(arr, low, pivot-1)
		Quicksort(arr, pivot+1, high)
	}
}

func QuicksortCall(arr *[]int) {
	Quicksort(arr, 0, len(*arr)-1)
}

/*
func test5(){
	arr := []int{-12,32,99,67,-83,123}
	quicksortCall(&arr)
	fmt.Println(arr)
}
*/


/*
============
6. Diseñar y construir una función que haga la búsqueda de un valor entero (la llave) en un arreglo que no
esté ordenado, mediante el algoritmo de búsqueda secuencial. La función debe retornar un par ordenado:
un valor booleano que indique si encontró la llave buscada (true = la encontró, false = no la encontró),
junto con un entero que indique la cantidad de comparaciones realizadas hasta determinar si la llave está o
no presente en el arreglo.
*/

func LinearSearch(arr []int, key int) (bool, int) {
	for index, value := range arr {
		if value == key {
			return true, index + 1 // Found, index of where it was + 1 [because it starts at 0]
		}
	}
	return false, len(arr) // Not Found, the whole length of the array
}

/*
func test6(){
	arr := []int{-12,32,99,67,-83,123}
	fmt.Println(linearSearch(arr, 99))
	fmt.Println(linearSearch(arr, 13))
}
*/


/*
============
7. Diseñar y construir una función que haga la búsqueda de un valor entero (la llave) en un arreglo que esté
ordenado, mediante el algoritmo de búsqueda binaria. La función debe retornar un par ordenado: un valor
booleano que indique si encontró la llave buscada (true = la encontró, false = no la encontró), junto con un
entero que indique la cantidad de comparaciones realizadas hasta determinar si la llave está o no presente
en el arreglo.
*/

func BinarySearch(arr []int, key int, left int, right int, n int) (bool, int) {
	if left <= right {
		mid := (left + right) / 2 // Mid value between left & right

		//fmt.Println("Mid: ", arr[mid], " Left: ", arr[left], " Right: ", arr[right])

		if arr[mid] == key { // Found
			return true, n
		} else if arr[mid] > key {
			return BinarySearch(arr, key, left, mid-1, n+1) // Key is lesser
		} else {
			return BinarySearch(arr, key, mid+1, right, n+1) // Key is greater
		}
	}
	return false, n // Not found
}

func BinarySearchCall(arr []int, key int) (bool, int) {
	return BinarySearch(arr, key, 0, len(arr)-1, 1)
}

/*
func test7() {
	arr := []int{1, 5, 10, 12, 16, 23, 29, 32, 35, 40, 49, 56, 57, 70} // in ascendent order

	fmt.Println(binarySearchCall(arr, 56))
	fmt.Println(binarySearchCall(arr, 24))
}
*/


/*
============
8. Diseñar una representación en Go para los árboles binarios de búsqueda ordinarios, donde cada nodo 
almacene un número entero.
*/

type Node struct {
	key int
	left *Node
	right *Node
}

type BSTree struct {
	root *Node
	count int
}

func (n *Node) Print() {
	if n != nil {
		fmt.Println(n.key)
	} else {
		fmt.Println("Nil Node")
	}
}


/*
============
9. Diseñar y construir una función que haga la inserción de un valor entero (la llave) en un árbol binario de 
búsqueda ordinario. Si la llave no se encuentra en el árbol, crea un nuevo nodo con la llave. Si la llave ya
se encuentra en el árbol, no la inserta. La función retorna un número entero, que será la cantidad de 
comparaciones realizadas, incluida la que llevó a la inserción.
*/

func (b *BSTree) InsertNode(value int, node *Node) *Node{
	b.count++ //New iteration

	if node == nil {
		node = &Node{key: value}
		return node
	} else if node.key == value {
		return node
	} else if value < node.key {
		node.left = b.InsertNode(value, node.left)
	} else {
		node.right = b.InsertNode(value, node.right)
	}; return node
}

func (b *BSTree) Insert(key int) int{
	b.count = 0
	b.root = b.InsertNode(key, b.root)
	return b.count
}


/*
============
10. Diseñar y construir una función que haga la búsqueda de un valor entero (la llave) en un árbol binario de 
búsqueda ordinario, sin hacer inserciones. La función debe retornar un par ordenado: un valor booleano 
que indique si encontró la llave buscada (true = la encontró, false = no la encontró), junto con un entero 
que indique la cantidad de comparaciones realizadas hasta determinar si la llave está o no presente en el 
árbol.
*/

func (b *BSTree) SearchNode(value int, node *Node) bool {
	b.count++

	if node == nil {
		return false
	} else if node.key == value {
		return true
	} else if value < node.key {
		b.SearchNode(value, node.left)
	} else {
		b.SearchNode(value, node.right)
	}; return false
}

func (b *BSTree) Search(key int) (bool, int){
	b.count = 0
	return b.SearchNode(key, b.root), b.count
}

/*
func testBSTree(){
	var b *BSTree = new(BSTree) //var b BSTree //

	b.insert(5)
	b.insert(7)
	b.insert(4)
	b.insert(8)
	b.insert(10)

	fmt.Println("Num: ", b.insert(0))

	b.root.print()
	b.root.right.print()
	b.root.left.print()
	b.root.right.right.print()
	b.root.right.right.right.print()

	_, a := b.search(5)

	fmt.Println("Search: ", a)
}
*/