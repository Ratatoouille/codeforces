package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	firstStr := sc.Text()

	sc.Scan()
	secondStr := sc.Text()

	firstStr = strings.ToLower(firstStr)
	secondStr = strings.ToLower(secondStr)

	if firstStr > secondStr {
		fmt.Println(1)
		return
	}
	if firstStr < secondStr {
		fmt.Println(-1)
		return
	}
	fmt.Println(0)
}
