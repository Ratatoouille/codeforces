package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	str := sc.Text()

	count := 0
	for i := 0; i < len(str)-1; i++ {
		count++
		if str[i] != str[i+1] && count < 7 {
			count = 0
			continue
		}
	}
	if count >= 7 || (count == 6 && str[len(str)-2] == str[len(str)-1]) {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}
