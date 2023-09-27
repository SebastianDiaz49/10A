package main

import (
	"fmt"

	md "github.com/SebastianDiaz49/10A/modules"
)

var test1 map[string]string

func main() {
	test1 = make(map[string]string)
	test1["hola"] = " Hola mundo "
	fmt.Println("hola test")
	fmt.Println(md.User)
	//fmt.Println(md.pUser)
}
