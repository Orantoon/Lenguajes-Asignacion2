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