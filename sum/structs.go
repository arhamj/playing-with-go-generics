package main

import "fmt"

type SumFn[T Numeric] func(...T) T

type Ledger[T ~string, K Numeric] struct {
	ID      T
	Amounts []K
	SumFn   SumFn[K]
}

func (l Ledger[T, K]) PrintIDAndSum() {
	fmt.Printf("%s has a sum of %v\n", l.ID, l.SumFn(l.Amounts...))
}
