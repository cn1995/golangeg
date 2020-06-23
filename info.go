package golangeg

import "fmt"

var info = "golang program example"

func GetInfo() string {
	return info
}

func PrintMe() {
	fmt.Println(info)
}
