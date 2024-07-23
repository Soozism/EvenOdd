package main

import "fmt"

func main() {
	ns := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, s := range ns {
		if s%2 == 0 {
			fmt.Printf("%v is even \n", s)
		}else{
			fmt.Printf("%v is odd \n", s)
		}
	}

}