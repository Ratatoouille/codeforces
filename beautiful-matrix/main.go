package main

import (
	"fmt"
	"math"
	"os"
)

func main() {

	var ii, jj int
	var t []int
	var a [][]int
	var numb int
	for i := 0; i < 5; i++ {
		t = t[:0]
		for j := 0; j < 5; j++ {
			fmt.Fscan(os.Stdin, &numb)
			if numb == 1 {
				ii = i
				jj = j
			}
			t = append(t, numb)
		}
		a = append(a, t)
	}

	fmt.Println(math.Abs(float64(ii-2)) + math.Abs(float64(jj-2)))
}
