package main

import (
	"celeritas"
	"fmt"
)

func main() {
	fmt.Println(celeritas.TestFunc(1, 2))
	fmt.Println(celeritas.TestFunc2(5, 2))
	fmt.Println(celeritas.TestFunc3(3, 2))
}
