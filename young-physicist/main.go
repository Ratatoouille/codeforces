package main

import (
	"fmt"
)

func main() {
	var count int
	fmt.Scan(&count)

	a := make([][]int, count)
	var numb int

	for i := 0; i < count; i++ {
		a[i] = make([]int, 3)

		for j := 0; j < 3; j++ {
			fmt.Scan(&numb)
			a[i][j] = numb
		}
	}

	res := make([]int, 3)
	for i := 0; i < count; i++ {
		res[0] += a[i][0]
		res[1] += a[i][1]
		res[2] += a[i][2]
	}

	if res[0] == 0 && res[1] == 0 && res[2] == 0 {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}
