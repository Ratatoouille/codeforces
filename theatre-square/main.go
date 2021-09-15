package main

import (
	"fmt"
)

func main() {
	var n, m, a int
	fmt.Scan(&n, &m, &a)

	fmt.Println(((n + a - 1) / a) * ((m + a - 1) / a))
}
