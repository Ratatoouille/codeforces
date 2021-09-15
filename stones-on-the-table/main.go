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
	str := sc.Text()

	count, err := strconv.Atoi(str)
	if err != nil {
		return
	}

	sc.Scan()
	str = sc.Text()

	res := 0
	for i := 0; i < count-1; i++ {
		if str[i] == str[i+1] {
			res++
		}
	}

	fmt.Println(res)
}
