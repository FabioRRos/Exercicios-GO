package main

import "fmt"

func main() {

	for horas := 0; horas <= 11; horas++ {
		for minutos := 0; minutos <= 59; minutos++ {
			for segundos := 1; segundos <= 59; segundos++ {
				if segundos < 10 {
					a := fmt.Sprint(": 0", segundos)
					fmt.Println(horas, ":", minutos, a)
				} else {
					fmt.Println(horas, ":", minutos, ":", segundos)
				}
			}
		}
	}

	fmt.Println("12 : 00 : 00")
}
