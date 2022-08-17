package main

import "fmt"

func main() {
	// functions
	fmt.Println(Sum(1, 2))

	sum1, _ := SumInts(int64(1), 10)
	fmt.Println(*sum1)

	sum2 := SumGenericOld(int64(1), 10)
	fmt.Println(sum2)

	sum3 := SumGeneric(int64(1), 10)
	fmt.Println(sum3)

	sum4 := SumGeneric(1i, 2i, complex128(10))
	fmt.Println(sum4)

	sum5 := SumGeneric[complex128](1i, 2i, 10)
	fmt.Println(sum5)

	sum6 := SumGeneric([]ProxyInt64{1, 2, 4}...)
	fmt.Println(sum6)

	// multiple generics
	PrintSum("id-0", SumGeneric[complex128], 1, 2, 3, 4i, 9i)

	// structs
	Ledger[string, complex128]{
		ID:      "id-0",
		Amounts: []complex128{1, 2, 3, 4i, 9i},
		SumFn:   SumGeneric[complex128],
	}.PrintIDAndSum()

	// struct constraints
	StructConstraint[string, int64, Ledger[string, int64]](Ledger[string, int64]{})

	// interface constraints
	PrintLedger[string, complex128, Ledger[string, complex128]](Ledger[string, complex128]{
		ID:      "id-3",
		Amounts: []complex128{1, 2, 3, 4i, 9i},
		SumFn:   SumGeneric[complex128],
	})
}
