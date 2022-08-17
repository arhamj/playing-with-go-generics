package main

import "fmt"

type HasID interface {
	~struct {
		ID string
	}
}

type CanGetID interface {
	GetID() string
}

type CanSetID interface {
	SetID(string)
}

type Unique struct {
	ID string
}

func (u *Unique) GetID() string {
	return u.ID
}

type UniqueName struct {
	Unique
	Name string
}

func (u *Unique) SetID(s string) {
	u.ID = s
}

func NewHasID[T HasID]() T {
	var t T
	return t
}

func New[T CanSetID]() T {
	var t T
	return t
}

func main() {
	hasIDImpl := NewHasID[Unique]()
	hasIDImpl.SetID("10")
	fmt.Println(hasIDImpl.GetID())
}
