package main

import (
	"fmt"
)

func Sum(args ...int) int {
	var sum int
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func SumInts(args ...interface{}) (*int, error) {
	var sum int
	for _, arg := range args {
		switch arg.(type) {
		case int64:
			sum += int(arg.(int64))
		case int:
			sum += arg.(int)
		default:
			return nil, fmt.Errorf("%T is not supported", arg)
		}
	}
	return &sum, nil
}

func SumGenericOld[T int | int64](args ...T) T {
	var sum T
	for _, arg := range args {
		sum += arg
	}
	return sum
}

type Numeric interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64 |
		~complex64 | ~complex128
}

type ProxyInt64 int64

func SumGeneric[T Numeric](args ...T) T {
	var sum T
	for _, arg := range args {
		sum += arg
	}
	return sum
}
