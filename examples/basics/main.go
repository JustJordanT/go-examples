package main

import "fmt"

var str string

func main() {
	fmt.Println("vim-go")
	addStr("Jordan")
}

func addStr(name string) string {
	str = name
	return str
}
