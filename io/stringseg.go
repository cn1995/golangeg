package main

import (
	"log"
	. "strings"
)

func main() {
	do1()
}

func do1() {
	Count("abcd", "a")
	join := Join([]string{"a", "b", "c"}, "")
	log.Fatal(join)
}
