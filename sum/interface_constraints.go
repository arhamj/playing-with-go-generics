package main

type ILedger[T ~string, K Numeric] interface {
	~struct {
		ID      T
		Amounts []K
		SumFn   SumFn[K]
	}
	PrintIDAndSum()
}

func PrintLedger[T ~string, K Numeric, L ILedger[T, K]](l L) {
	l.PrintIDAndSum()
}
