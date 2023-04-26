package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	line := input.Text()

	if len(line) == 0 {
		return
	}

	strs := strings.Split(line, " ")

	//2 abc 123456789
	num, err := strconv.Atoi(strs[0])
	if err != nil {
		return
	}
	result := []string{}
	dv := "00000000"
	for i := 1; i <= num; i++ {
		s := strs[i]
		fmt.Println(s)
		l := 0
		for l < len(s) {
			min := int(math.Min(float64(len(s)-l), 8))
			v := s[l : l+min]
			if len(v) < 8 {
				v = v + dv[:8-len(v)]
			}
			result = append(result, v)
			l = l + 8
		}
	}
	sort.Strings(result)
	l := len(result)
	for i, s := range result {
		if i == l-1 {
			fmt.Print(s)
		} else {
			fmt.Printf("%s ", s)
		}
	}
}
