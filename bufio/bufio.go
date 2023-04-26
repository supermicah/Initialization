package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := strings.Split(input.Text(), " ")
		fmt.Println(len(arr[len(arr)-1]))
	}
}
