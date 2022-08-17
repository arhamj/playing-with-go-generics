package main

func StructConstraint[T ~string, K Numeric, L ~struct {
	ID      T
	Amounts []K
	SumFn   SumFn[K]
}](l L) {
}
