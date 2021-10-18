package main

import (
	"asignacion_mod/funciones"
	"fmt"
)

func test3(){
	slice1, slice2 := []int{-12,32,99,67,-83,123}, []int{0,1,2,3,4}

	n1, n2 := 13, 2
	three1, three2 := funciones.InsertVal(&slice1,n1), funciones.InsertVal(&slice2, n2)

	fmt.Printf("Array 1 = Key: %d, Comparisons done: %d, Final Slice: %v\nArray 2 = Key: %d, Comparisons done: %d, Final Slice: %v", n1, three1, slice1, n2, three2, slice2)
}

func test4(){
	arr := []int{-12,32,99,67,-83,123}
	funciones.SelectionSort(&arr)
	fmt.Println(arr)
}

func test5(){
	arr := []int{-12,32,99,67,-83,123}
	funciones.QuicksortCall(&arr)
	fmt.Println(arr)
}

func test6(){
	arr := []int{-12,32,99,67,-83,123}
	fmt.Println(funciones.LinearSearch(arr, 99))
	fmt.Println(funciones.LinearSearch(arr, 13))
}

func test7() {
	arr := []int{1, 5, 10, 12, 16, 23, 29, 32, 35, 40, 49, 56, 57, 70} // in ascendent order

	fmt.Println(funciones.BinarySearchCall(arr, 56))
	fmt.Println(funciones.BinarySearchCall(arr, 24))
}

func testBSTree(){
	var b *funciones.BSTree = new(funciones.BSTree) //var b BSTree //

	b.Insert(5)
	b.Insert(7)
	b.Insert(4)
	b.Insert(8)
	b.Insert(10)

	fmt.Println("Num: ", b.Insert(0))

	b.Root.Print()
	b.Root.Right.Print()
	b.Root.Left.Print()
	b.Root.Right.Right.Print()
	b.Root.Right.Right.Right.Print()

	_, a := b.Search(5)

	fmt.Println("Search: ", a)
}