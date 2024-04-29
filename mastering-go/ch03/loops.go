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

}
