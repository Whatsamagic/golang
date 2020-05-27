package main

import (
	"fmt"
	"math/rand"
	"time")

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(time.Now())
	fmt.Println(time.Now().UTC())
	fmt.Println(time.Now().UTC().UnixNano())

	fmt.Println("random nr is", rand.Intn(10))
}
