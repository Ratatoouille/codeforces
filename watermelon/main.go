package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	text := sc.Text()
	x, err := strconv.Atoi(text)
	if err != nil {
		log.Fatal(err)
	}

	if x != 2 && x%2 == 0 {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}
