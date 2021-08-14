// result
//      *
//     ***
//    *****
//   *******
//  *********

package main

import "fmt"

func main() {
	result := ""

	for i := 0; i < 5; i++ {
		for j := 4 - i; j > 0; j-- {
			result += " "
		}
		for k := 4 - i; k < 5; k++ {
			result += "*"
		}
		for l := 0; l < i; l++ {
			result += "*"
		}
		fmt.Println(result)
		result = ""

	}
}
