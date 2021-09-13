package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	vowels := "aoyeui"

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	str := sc.Text()

	str = strings.ToLower(str)

	for _, v := range vowels {
		str = strings.ReplaceAll(str, string(v), "")
	}

	var res string
	for _, s := range str {
		res += "." + string(s)
	}
	fmt.Println(res)
}
