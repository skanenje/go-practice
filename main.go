package main

import (
	"fmt"
)

// func compare(s1, s2 string) int {
// 	for i := 0; i < min(len(s1), len(s2)); i++ {
// 		if s1[i] < s2[i] {
// 			return -1
// 		} else if s1[i] > s2[i] {
// 			return 1
// 		}
// 	}
// 	if len(s1) < len(s2) {
// 		return -1
// 	} else if len(s1) > len(s2) {
// 		return 1
// 	} else {
// 		return 0
// 	}
// }

// func min(x, y int) int {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }

// func main() {
// 	fmt.Println(compare("Hello!", "Hello!"))
// 	fmt.Println(compare("Salut!", "lut!"))
// 	fmt.Println(compare("Ola!", "Ol"))
// }

// func compare(a, b string)int{

// }
// func caseChecker(r rune)bool{

// }
// func caseConvert(r rune) rune{

// }
// func alhaMir(){
// 	args := os.Args[1:]
// 	if len(args) == 0 {
// 		return
// 	}
// 	s := args[0]
// 	var str string
// 	for _, v := range s {
// 		if (v >= 'A' && v <= 'Z') && (v + 25) > 'Z' {
// 			v = 'A' + ('Z' - v)
// 			str += string(v)
// 		}else if (v >= 'a' && v <= 'z') && (v + 25) > 'z' {
// 			v = 'a' + ('z' - v)
// 			str += string(v)
// 		}else if (v < 'A' || v > 'Z') || v == ' ' {
// 			str += string(v)

// 		}

// 	}
// 	fmt.Println(str)
// }
// func main(){
// 	alhaMir()
// }
// "os"
// "fmt"

// func ReduceInt(a []int, f func(int, int) int) {
// 	var z int
//  if len(a) > 0 {
// 	x := a[0]
// 	y := a[1]
// 	z = f(x, y)
//  }
//   s := Itoa(z)
//   for _, v := range s {
// 	z01.PrintRune(v)
//   }
//   z01.PrintRune('\n')
// }
// func Itoa(n int)string{
// 	var isNeg bool
// 	if n < 0 {
// 		isNeg = true
// 		n = -n
// 	}
// 	result := ""
// 	for n > 0 {
// 		digit := n % 10
// 		result = string(rune('0' + digit)) + result
// 		n /= 10
// 	}
// 	if isNeg {
// 		result = "-"+result
// 	}
// 	return result
// }
// func Atoi(s string)int {
// 	var isneg bool
// 	result := 0
// 	for i, v := range s {
// 		if i == 0 && v == '-' {
// 			isneg = true
// 		}else if (v >= '0' && v <= '9') {
// 			digit := int(v - '0')
// 			result = result*10 + digit
// 		}
// 	}
// 	if isneg {
// 		result = -result
// 	}
// 	return result
// }
// func main(){
// 	s := "-123"
// 	r := Atoi(s)
// 	fmt.Printf("result ni .. %d\n", r)
// }
// func main() {
// 	mul := func(acc int, cur int) int {
// 		return acc * cur
// 	}
// 	sum := func(acc int, cur int) int {
// 		return acc + cur
// 	}
// 	div := func(acc int, cur int) int {
// 		return acc / cur
// 	}
// 	as := []int{500, 2}
// 	ReduceInt(as, mul)
// 	ReduceInt(as, sum)
// 	ReduceInt(as, div)
// }
// func main(){
//  	args := os.Args[1:]
// 	if len(args) == 0{
// 		return
// 	}
// 	s1 := args[0]
// 	res := Rot13(s1)
// 	for _, v := range res {
// 		z01.PrintRune(v)
// 	}
// 	z01.PrintRune('\n')
// }
// func Rot13(s string)string{
// 	var s0 string
// 	for _, v := range s {
// 		if (v >= 'a' && v <= 'z') &&  (v + 13) > 'z'{
// 			v = (v + 13) - 26
// 			s0 += string(v)
// 		}else if (v >= 'a' && v <= 'z') &&  (v + 13) <= 'z' {
// 			v = (v + 13)
// 			s0 += string(v)
// 		}else if (v >= 'A' && v <= 'Z') &&  (v + 13) > 'Z'{
// 			v = (v + 13) - 26
// 			s0 += string(v)
// 		}else if (v >= 'A' && v <= 'Z') &&  (v + 13) <= 'Z' {
// 			v = (v + 13)
// 			s0 += string(v)
// 		}else if (v < 'A' || v > 'Z') || v == ' ' {
// 			s0 += string(v)
// 		}
// 	}
// 	return s0
// }
// package main

// import (
// 	"fmt"
// )

// func Chunk(sl []int, sz int) {
// 	var sli1 [][]int
// 	var sli2 []int
// 	if sz == 0 {
// 		fmt.Println()
// 	}else {
// 		count := 0
// 	for i := 0; i < len(sl); i++ {
// 		sli2 = append(sli2, sl[i])
// 		count++
// 		if count == sz {
// 			count = 0
// 			sli1 = append(sli1, sli2)
// 			sli2 = nil
// 		}

// 	}
// 	if sli2 != nil {
// 		sli1 = append(sli1, sli2)
// 	}
// 	fmt.Println(sli1)
// 	}

// }
// func main() {
// 	Chunk([]int{}, 10)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {

// }
// func sumDig(x int) int {
// 	s := Itoa(x)
// 	sum := 0
// 	for i := 0; i < len(s); i++ {
// 		b := Atoi(string(s[i]))
// 		sum += b
// 	}
// 	return sum
// }
// func digitalRoot(n int)int {
// 	if n < 10 {
// 		return n
// 	}
// 	return digitalRoot(sumDig(n))
// }

// func main() {
// 	args := os.Args[1:]
// 	if len(args) == 0 {
// 		return
// 	}
// 	s1 := args[0]
// 	x := Atoi(s1)
// 	r := digitalRoot(x)
// 	fmt.Printf(" ...ni..%d\n", r)
// }

// func Itoa(n int) string {
// 	var isNeg bool
// 	if n < 0 {
// 		isNeg = true
// 		n = -n
// 	}
// 	result := ""
// 	for n > 0 {
// 		digit := n % 10
// 		result = string(rune('0'+digit)) + result
// 		n /= 10
// 	}
// 	if isNeg {
// 		result = "-" + result
// 	}
// 	return result
// }

// func Atoi(s string) int {
// 	var isneg bool
// 	result := 0
// 	for i, v := range s {
// 		if i == 0 && v == '-' {
// 			isneg = true
// 		} else if v >= '0' && v <= '9' {
// 			digit := int(v - '0')
// 			result = result*10 + digit
// 		}
// 	}
// 	if isneg {
// 		result = -result
// 	}
// 	return result
// }
