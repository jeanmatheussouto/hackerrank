package main

import "fmt"

func main() {
	var n, temp int
	var arr []int
	var x float64

	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &temp)
		arr = append(arr, temp)
	}

	var pos, neg, zer float64 = 0.0, 0.0, 0.0

	for _, v := range arr {
		switch {
		case v == 0:
			zer++
		case v < 0:
			neg++
		case v > 0:
			pos++
		}
	}

	x = float64(n)
	fmt.Printf("%f\n%f\n%f\n", pos/x, neg/x, zer/x)
}
