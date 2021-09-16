package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	str := sc.Text()
	fmt.Printf("%s%s", string(unicode.ToUpper(rune(str[0]))), str[1:])
}
