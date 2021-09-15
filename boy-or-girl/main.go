package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	str := sc.Text()
	res := len(str)

	s := strings.Split(str, "")
	sort.Strings(s)
	str = strings.Join(s, "")

	for i := 0; i < len(s)-1; i++ {
		if str[i] == str[i+1] {
			res--
		}
	}

	if res%2 == 0 {
		fmt.Println("CHAT WITH HER!")
		return
	}
	fmt.Println("IGNORE HIM!")
}
