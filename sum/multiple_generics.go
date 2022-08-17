package main

import "fmt"

func PrintSum[T ~string, K Numeric](id T, fn SumFn[K], values ...K) {
	fmt.Printf("%s has a sum of %v\n", id, fn(values...))
}
