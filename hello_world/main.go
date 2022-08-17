package main

import (
	"fmt"
)

type request struct {
	host *string
	port *int
}

func printReq(r request) {
	fmt.Print("request: host=")
	if r.host != nil {
		fmt.Print(*r.host)
	}
	fmt.Print(", port=")
	if r.port != nil {
		fmt.Printf("%d", *r.port)
	}
	fmt.Println()
}

func PtrInt(i int) *int {
	return &i
}

func PtrString(s string) *string {
	return &s
}

func Ptr[T any](value T) *T {
	return &value
}

func main() {
	// the use of pointers to indicate optional fields is the problem here
	printReq(request{
		host: nil, // needs a *string
		port: nil, // needs a *int
	})

	// Solution 1: using temporary variables
	// Cluttered and hard to read
	host, port := "local", 80
	printReq(request{
		host: &host, // needs a *string
		port: &port, // needs a *int
	})

	// Solution 2: typed helper functions
	// One Ptr function per type is needed
	printReq(request{
		host: PtrString("test"), // needs a *string
		port: PtrInt(10),        // needs a *int
	})

	// Solution 3: using generics
	printReq(request{
		host: Ptr("test"),
		port: Ptr(10),
	})
}
