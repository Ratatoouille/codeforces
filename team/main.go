package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()

	str := sc.Text()
	count, err := strconv.Atoi(str)
	if err != nil {
		return
	}

	res := 0
	matrix := make([][]int, count, count)
	for i := 0; i < count; i++ {
		matrix[i] = make([]int, 3, 3)
		for j := 0; j < 3; j++ {
			sc.Scan()
			str := sc.Text()
			number, err := strconv.Atoi(str)
			if err != nil {
				return
			}
			matrix[i][j] = number
		}
		if matrix[i][0]+matrix[i][1]+matrix[i][2] >= 2 {
			res += 1
		}
	}
	fmt.Println(res)
}
