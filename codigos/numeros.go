package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		if i%3 == 0 {
			fmt.Printf("%d é divisível por 3\n", i)
		}
	}

	fmt.Print("===== Parte 2\n =====")

	for i := 0; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Pin Pan")
		} else if i%3 == 0 {
			fmt.Println("Pin")
		} else if i%5 == 0 {
			fmt.Println("Pan")
		} else {
			fmt.Println(i)
		}
	}
}
