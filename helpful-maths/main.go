package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	str := sc.Text()

	numbers := strings.Split(str, "+")

	var res []int
	for _, v := range numbers {
		numb, err := strconv.Atoi(v)
		if err != nil {
			return
		}

		res = append(res, numb)
	}
	sort.Ints(sort.IntSlice(res))

	str = ""
	for i, v := range res {
		if i == len(res)-1 {
			str += strconv.Itoa(v)
			break
		}
		str += strconv.Itoa(v) + "+"
	}

	fmt.Println(str)
}
