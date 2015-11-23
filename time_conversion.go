package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var t string

	fmt.Scanf("%s", &t)
	a := t[len(t)-2:]
	t = t[:len(t)-2]
	s := strings.Split(t, ":")
	c, _ := strconv.Atoi(s[0])

	if c <= 1 || c >= 12 {
		s[0] = "00"
	}

	if a == "PM" {
		c, _ := strconv.Atoi(s[0])
		c = c + 12
		s[0] = strconv.Itoa(c)
	}
	fmt.Println(strings.Join(s, ":"))
}
