package main

import (
	"events/framework/utils"
	"fmt"
	"time"
)

type ASDF struct {
	Date time.Time
}

func main() {
	// future := time.Now().Add(time.Hour * 24 * 7)
	// asdf, _ := time.Parse(time.RFC3339, "2022-05-03T13:21:36.639Z")
	// fmt.Println(time.Until(asdf) <= time.Hour*24)
	// fmt.Println(future)

	now1 := ASDF{Date: time.Now()}
	now2 := ASDF{}

	fmt.Println(utils.MergeObjects(&now1, now2))
	fmt.Println(now1)
}
