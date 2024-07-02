package main

import (
	"fmt"
)
// func isPowerof(x, exp int)bool{
// 	for x > 0{
// 		if x%exp == 0{
// 			return true
// 		}
// 		x /= exp
// 	}
// 	return false
// }
// func main(){
// 	res := isPowerof(128, 2)
// 	fmt.Println(res)
// }
// func main(){
// 	args := os.Args[1:]
// 	if len(args) != 2 {
// 		return
// 	}
// 	s1 := args[0]
// 	// s2 := args[1]
// 	var letters1 map[rune]bool = make(map[rune]bool)
// 	for _, v := range s1 {
// 		letters1[v] = true
		
// 	}
	// var letters2 map[rune]bool = make(map[rune]bool)
	// for _, v := range s2 {
	// 	if letters1[v] && !letters2[v]{
	// 		letters2[v] = true
	// 	}
	// }
// 	for _, v := range s1 {
// 		if letters1[v]{
// 			z01.PrintRune(v)
// 			letters1[v] = false
// 		}
// 	}
// 	z01.PrintRune('\n')
// }
// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 2 {
// 		return
// 	}
// 	s1 := args[0]
// 	s2 := args[1]
// 	var letters map[rune]bool = make(map[rune]bool)
// 	for _, v := range s1 {
// 		letters[v] = true
// 	}
// 	res := ""
// 	for _, v := range s2 {
// 		if letters[v] {
// 			letters[v] = false
// 			res += string(v)
// 		}
// 	}
// 	fmt.Println(res)
// }

// func main(){
// args := os.Args[1:]
// if len(args) != 2 {
// 	return
// }
// s1:=args[0]
// s2:=args[1]
// var letters map[rune]bool= make(map[rune]bool)

// for _, v := range s1 {
// 	letters[v] = true
// }
// var letters2 map[rune]int = make(map[rune]int)
// for _, v := range s2 {
// 	if letters[v]{
// 		letters2[v]++
// 	}
// }
// for _, v := range s1 {
// 	if count, ok := letters2[v] ; ok && count > 0 {
// 		z01.PrintRune(v)
// 		letters2[v]--
// 	}
// }
// z01.PrintRune('\n')
// }
// @@@@@@@@#########################################################################3
// func main(){
// 	args := os.Args[1:]
// 	if len(args) != 3 {
// 		return
// 	}
// 	search := args[1]
// 	replace := args[2]
// 	var s rune
// 	if len(search) == 1 {for _, v := range search {s = v}}
// 	var r rune
// 	if len(replace) == 1 {for _, v := range replace {r = v}}
// 	res := ""
// 	if contains(args[0], search){
// 		for _, v := range args[0] {
// 			if v == s {
// 				v = r
// 			}
// 			res += string(v)
// 		}
// 		fmt.Println(res)
// 	}else {
// 		fmt.Println(args[0])
// 	}

// }
// func contains(s string, sub string)bool{
// 	for _, v := range s {
// 		if sub == string(v) {
// 			return true
// 		}
// 	}
// 	return false
// }
//###################################################################################333
// func FoldInt(f func(int, int)int, sl []int, ac int){
// 	res := ac
// 	for _, n := range sl {
// 		res = f(res, n)
// 	}
// 	fmt.Println(res)
// }
// func main() {
// 	table := []int{1, 2, 3}
// 	ac := 93
// 	FoldInt(Add, table, ac)
// 	FoldInt(Mul, table, ac)
// 	FoldInt(Sub, table, ac)
// 	fmt.Println()

// 	table = []int{0}
// 	FoldInt(Add, table, ac)
// 	FoldInt(Mul, table, ac)
// 	FoldInt(Sub, table, ac)
// }
// func Add(a, b int)int{
// 	return a + b
// }
// func Mul(a, b int)int{
// 	return a * b
// }
// func Sub(a, b int)int {
// 	return a - b
// }
// @##############################################################################
// func main(){
// 	args := os.Args[1:]
// 	if len(args) == 0 {
// 		return
// 	}
// 	if len(args) > 3 {
// 		return
// 	}
// 	if args[1] != "+" && args[1] != "-" && args[1] != "/" && args[1] != "*" && args[1] != "%"  {
// 		return
// 	}
// 	x,_ := atoi(args[0])
// 	if x >= 9223372036854775807 || x <= -9223372036854775807 {
// 		return
// 	}
// 	y,_ := atoi(args[2])
// 	 if y >= 9223372036854775807 || y <= -9223372036854775807 {
// 		return
// 	}
// 	switch args[1]{
// 	case "+":
// 		result := x + y
// 		_,_ = os.Stdout.Write([]byte(itoa(result)+ "\n"))
// 	case "-":
// 		result := x - y
// 		_,_ = os.Stdout.Write([]byte(itoa(result) + "\n"))
// 	case "*":
// 		result := x * y
// 		_,_ = os.Stdout.Write([]byte(itoa(result)+ "\n"))
// 	case "/":
// 		if y == 0 {
// 			_,_ = os.Stdout.Write([]byte("No division by 0\n"))
// 		}else {
// 			result := x / y
// 		_,_ = os.Stdout.Write([]byte(itoa(result)+ "\n"))
// 		}
// 	case "%":
// 		if y == 0 {
// 			_,_ = os.Stdout.Write([]byte("No modulo by 0\n"))
// 		}else {
// 			result := x % y
// 		_,_ = os.Stdout.Write([]byte(itoa(result)+ "\n"))
// 		}

// 	}

// }
// func itoa(x int) string{
// 	var isneg bool
// 	result := ""
// 	if x < 0 {
// 		isneg = true
// 		x = -x
// 	}
// 	for x > 0 {
// 	    digit:= x % 10
// 		result = string(rune('0' + digit)) + result
// 		x /= 10
// 	}
// 	if isneg {
// 		result = "-" + result
// 	}
// 	return result
// }
// func atoi(s string)(int, error){
// 	var isneg bool
// 	result := 0
// 	for i, n := range s {
// 		if i == 0 && n == '-'{
// 			isneg = true
// 		}else if (n < '0' || n > '9') {
// 			err := "Invalid number for conversion"
// 			return 0, fmt.Errorf("Error:%s", err)

// 		}else{
// 			digit := int(rune(n - '0'))
// 			result = result*10 + digit
// 		}
// 	}
// 	if isneg {
// 		result = -result
// 	}
// 	return result, nil
// }
// @#################################################################################

// func main() {
// 	s := os.Args[1]
// 	fmt.Printf("%s\n", rrot13(s))
// func rrot13(s string) string {
// 	res := ""
// 	for _, v := range s {
// 		if (v >= 'a' && v <= 'z') && (v+13) > 'z' {
// 			v = (v + 13) - 26

// 		} else if (v >= 'a' && v <= 'z') && (v+13) <= 'z' {
// 			v = (v + 13)

// 		} else if (v >= 'A' && v <= 'Z') && (v+13) > 'Z' {
// 			v = (v + 13) - 26

// 		} else if (v >= 'A' && v <= 'Z') && (v+13) <= 'Z' {
// 			v = (v + 13)

// 		}
// 		res += string(v)

// 	}
// 	return res
// }
// ####################################################################
// func main(){
// 	chunk2([]int{}, 10)
// 	chunk2([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
// 	chunk2([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
// 	chunk2([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
// 	chunk2([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
// }
// func chunk2(sl []int, sz int){
// 	resStr := ""
// 	if sz == 0 {
// 		z01.PrintRune('\n')
// 	}else {
// 		if len(sl) == 0 {
// 			z01.PrintRune('[')
// 			z01.PrintRune(']')
// 			z01.PrintRune('\n')
// 		}else {
// 			for _, n := range sl {
// 				resStr+=itoa(n)
// 			}
// 			if len(resStr) < sz  {
// 				// do something
// 				z01.PrintRune('[')
// 				z01.PrintRune('[')
// 				for i, n := range resStr {
// 					z01.PrintRune(n)
// 					if i == len(resStr)-1{
// 						break
// 					}
// 					z01.PrintRune(' ')
// 				}
// 				z01.PrintRune(']')
// 				z01.PrintRune(']')
// 			}
// 			quotient := len(resStr) / sz
// 			rem := len(resStr) % sz
// 			var asli []string
// 			for quotient > 0 {
// 				asli = append(asli, resStr[:sz])
// 				resStr = resStr[sz:]
// 				quotient--
// 			}
// 			if rem != 0 {
// 				asli = append(asli, resStr[(len(resStr)-rem):])
// 			}
// 			z01.PrintRune('[')
// 			for _, astr := range asli {
// 				z01.PrintRune('[')
// 				for i, n := range astr {
// 					z01.PrintRune(n)
// 					if i == len(astr)-1{
// 						break
// 					}
// 					z01.PrintRune(' ')
// 				}
// 				z01.PrintRune(']')
// 			}
// 			z01.PrintRune(']')
// 			z01.PrintRune('\n')
// 		}

// 	}

// }
// func itoa(x int)string{
// 	var isNeg bool
// 	result := ""
// 	if x < 0 {
// 		isNeg = true
// 		x = -x
// 	}else if x == 0 {
// 		return "0"
// 	}
// 	for x > 0 {
// 		digit := x%10
// 		result =  string(rune('0' + digit))+ result
// 		x /=10
// 	}
// 	if isNeg{
// 		result = "-"+result
// 	}
// 	return result
// }
// ####################################################################################3
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
// ##################################################################################3
// func alphaMir(){
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
// 	alphaMir()
// }
// "os"
// "fmt"
// ################################################################################3
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
// #########################################################################3
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
// ###################################################################################
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
// ################################################################################3
// package main

// import (
// 	"fmt"
// )

// func Chunk(sl []int, sz int) {
// 	var sli1 [][]int
// 	var sli2 []int
// 	if sz == 0 {
// 		fmt.Println()
// 	} else {
// 		count := 0
// 		for i := 0; i < len(sl); i++ {
// 			sli2 = append(sli2, sl[i])
// 			count++
// 			if count == sz {
// 				count = 0
// 				sli1 = append(sli1, sli2)
// 				sli2 = nil
// 			}

// 		}
// 		if sli2 != nil {
// 			sli1 = append(sli1, sli2)
// 		}
// 		fmt.Println(sli1)
// 	}

// }
// func main() {
// 	Chunk([]int{}, 10)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
// }

// ############################################################################3
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
