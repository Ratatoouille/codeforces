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

	count, err := strconv.Atoi(text)
	if err != nil {
		log.Fatalln(err)
		return
	}

	strs := make([]string, count, count)
	for i := 0; i < count; i++ {
		sc.Scan()
		strs = append(strs, sc.Text())
	}

	for _, str := range strs {
		if len(str) <= 10 {
			fmt.Println(str)
		} else {
			fmt.Printf("%c%v%c\n", str[0], len(str)-2, str[len(str)-1])
		}
	}
}
