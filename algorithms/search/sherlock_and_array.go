package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	t, _ := strconv.Atoi(s.Text()) //number of test cases

	for i := 0; i < t; i++ {
		s.Scan()
		n, _ := strconv.Atoi(s.Text()) //number of elements in the array

		a := make([]int, n)
		for j := 0; j < n; j++ {
			s.Scan()
			a[j], _ = strconv.Atoi(s.Text())
		}

		left, right := 0, len(a)-1
		sumLeft, sumRight := 0, 0
		for {
			if left == right {
				if sumRight == sumLeft {
					fmt.Printf("YES\n")
				} else {
					fmt.Printf("NO\n")
				}
				break
			}
			if sumLeft > sumRight {
				sumRight += a[right]
				right--
			} else {
				sumLeft += a[left]
				left++
			}
		}
	}
}
