// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("========= Código original =========")
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	fmt.Println("========= Exercício 1.1 =========")
	s2, sep2 := "", ""
	for _, arg2 := range os.Args[0:] {
		s2 += sep2 + arg2
		sep2 = " "
	}
	fmt.Println(s2)

	fmt.Println("========= Exercício 1.2 =========")
	for i, arg2 := range os.Args[1:] {
		fmt.Printf("Indice: [%v] | Valor: [%v] \n", i, arg2)
	}
}
