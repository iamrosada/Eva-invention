package main

import "fmt"

func main() {

	for i := 0; i < 100; i++ {

		if i%20 == 0 {
			continue
		}

		if i == 16 {

			break
		}

		fmt.Println(i, " ")

	}

	fmt.Println()

	i := 10

	for {
		if i < 0 {

			break
		}

		fmt.Print(i, " ")
		i--
	}

	fmt.Println()

}
