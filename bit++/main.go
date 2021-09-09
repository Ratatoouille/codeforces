package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	data := sc.Text()
	numb, err := strconv.Atoi(data)
	if err != nil {
		return
	}

	var x = 0
	for i := 0; i < numb; i++ {
		sc.Scan()
		data := sc.Text()
		if data[0] == '-' || data[2] == '-' {
			x--
		}
		if data[0] == '+' || data[2] == '+' {
			x++
		}
	}
	fmt.Println(x)
}
