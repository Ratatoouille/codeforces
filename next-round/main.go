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

	sc.Scan()
	data = sc.Text()
	str = strings.Split(data, " ")

	var numbers []int
	for i := 0; i < nk[0]; i++ {
		numb, err := strconv.Atoi(str[i])
		if err != nil {
			return
		}
		numbers = append(numbers, numb)
	}

	t := numbers[nk[1]-1]
	res := 0

	for i := 0; i < len(numbers); i++ {
		if numbers[i] > 0 && numbers[i] >= t {
			res++
		}
	}
	fmt.Println(res)
}
