package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	arr := strings.Split(text, " ")
	sum := 0
	for _, v := range arr {
		i, _ := strconv.Atoi(v)
		sum += i
	}
	fmt.Println(sum)
}
