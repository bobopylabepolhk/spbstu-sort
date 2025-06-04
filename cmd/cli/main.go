package main

import (
	"flag"
	"fmt"

	"github.com/bobopylabepolhk/spbstu-sort/internal/generator"
	"github.com/bobopylabepolhk/spbstu-sort/internal/sorting"
)

type flags struct {
	size int
	ceil int
}

func main() {
	f := flags{}
	flag.IntVar(&f.size, "size", 1, "size of initial unsorted slice")
	flag.IntVar(&f.ceil, "ceil", 10000, "maximum value of generated int")
	flag.Parse()

	if f.size < 1 {
		msg := fmt.Sprintf("invalid size %v can't be more less than 1", f.size)
		panic(msg)
	}

	// if f.alg != "bubble" && f.alg != "selection" && f.alg != "insertion" {
	// 	msg := fmt.Sprintf("invalid algorhitm %v available values: bubble, selection, insertion. bubble is default", f.size)
	// 	panic(msg)
	// }

	gen := generator.NewDataGenerator(f.ceil)
	dataset1 := gen.GenIntSlice(f.size)
	dataset2 := make([]int, 0)
	dataset2 = append(dataset2, dataset1...)
	var compareCount int
	var shiftCount int

	fmt.Println("список:")
	fmt.Println(dataset1, "\n")

	compareCount, shiftCount = sorting.BubbleSort(&dataset1)

	fmt.Println("пузырьковая сортировка:")
	fmt.Println("сравнения:", compareCount, "перестановки:", shiftCount, "\n")

	compareCount, shiftCount = sorting.InsertionSort(&dataset2)

	fmt.Println("сортировка вставкой:")
	fmt.Println("сравнения:", compareCount, "перестановки:", shiftCount, "\n")

	fmt.Println("отсортированный список:")
	fmt.Println(dataset2)
}
