package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var nk []int

	sc.Scan()
	data := sc.Text()
	str := strings.Split(data, " ")

	for i := 0; i < 2; i++ {
		numb, err := strconv.Atoi(str[i])
		if err != nil {
			return
		}
		nk = append(nk, numb)
	}

	fmt.Println(int(nk[0] * nk[1] / 2))
}
