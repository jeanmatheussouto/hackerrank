package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	// t  number of test cases.
	// var n int // students in the class
	// var k int // the cancelation threshold
	// var a int // the arrival times for each student

	t, _, _ := in.ReadLine()
	line, _ := in.ReadString('\n')

	fmt.Println(t, line)
	// fmt.Println("---")
	// fmt.Println(t, n, k, a)
}
