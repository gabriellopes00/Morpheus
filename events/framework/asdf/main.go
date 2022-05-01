package main

import (
	"fmt"

	"github.com/imdario/mergo"
)

type a struct {
	CoverUrl string
	Email    string
	Name     string
}

func main() {
	a1 := &a{Name: "A", Email: "aemail", CoverUrl: "first"}
	a2 := &a{CoverUrl: "second", Email: "", Name: ""}

	mergo.Merge(a1, a2, mergo.WithOverride)

	fmt.Println(a1)
}
