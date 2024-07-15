package main

import (
	// "fmt"
	"os"
	"github.com/01-edu/z01"
)
// ######################################## PRINT HEX#########################################	
// var hex_map map[int]string = map[int]string{
// 	0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9",
// 	10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f",
// }

// func toHex(x int) string {
// 	if x == 0 {
// 		return "0"
// 	}
// 	result := ""
// 	for x > 0 {
// 		rem := x % 16
// 		result = hex_map[rem] + result
// 		x /= 16
// 	}
// 	return result
// }
// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 1 {
// 		return
// 	}
// 	res := ""
// 	n, err := atoi(args[0])
// 	if err != nil {
// 		fmt.Print(err.Error())
// 	} else {
// 		res = toHex(n)

// 	}
// 	fmt.Printf("%s\n", res)

// }
// func atoi(s string) (int, error) {
// 	isneg := false
// 	result := 0
// 	if s == "0" {
// 		return 0, nil
// 	}
// 	for i, v := range s {
// 		if i == 0 && v == '-' {
// 			isneg = true
// 		} else if v < '0' || v > '9' {
// 			return -1, fmt.Errorf("ERROR")
// 		} else if v >= '0' && v <= '9' {
// 			digit := int(v - '0')
// 			result = result*10 + digit
// 		}

// 	}
// 	if isneg {
// 		result = -result
// 	}
// 	return result, nil
// }
// @####################################################### REPEAT ALPHA #########################################
// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 1 {
// 		return
// 	} else if args[0] == "" {
// 		return
// 	} else {
// 		s := repeatAlpha(args[0])
// 		fmt.Println(s)
// 	}

// }
// func repeatAlpha(s string) string {
// 	result := ""
// 	// index := 0
// 	for _, v := range s {
// 		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') {
// 			index := int(v - 'a')
// 			if v >= 'A' && v <= 'Z' {
// 				index = int(v - 'A')
// 			}
// 			for i := 0; i <= index; i++ {
// 				result += string(v)
// 			}
// 		} else {
// 			result += string(v)
// 		}

// 	}
// 	return result
// }
// @@@@################################################ REVERSE BITS ##########################
// func reverse(oct byte)byte{
// 	var result byte
// 	for i := 0 ; i < 8 ; i++{
// 		bit := (oct >> i) & 1

// 		result |= bit << (7-i)
// 	}
// 	return result
// }
// func main(){
// 	var b byte= 0b00100110
// 	reversed := reverse(b)
// 	fmt.Printf("reversed bytes is \n %0b\n", reversed)
// }

// ####################### UNION ################################
// func main(){
// 	args := os.Args[1:]
// 		if len(args) != 2 {
// 			return
// 		}
//  s := args[0] + args[1]
//   union(s)
// }
// func union(s string){
// 	var noDoubles map[rune]bool = map[rune]bool{}

// 	for _, v := range s {
// 		if !noDoubles[v]{
// 			z01.PrintRune(v)
// 			noDoubles[v] = true
// 		}
// 	}
// 	z01.PrintRune('\n')
// }
// ####################################3 SWAP BITS @##############################
// func SwapBits(octet byte) byte {
// 	// Extract the upper and lower halves
// 	upperHalf := octet >> 4    // Shift right by 4 bits to get upper half
// 	lowerHalf := octet & 0x0F  // Mask with 0x0F (00001111) to get lower half

// 	// Swap and combine them
// 	return (lowerHalf << 4) | upperHalf
// }

// func main() {
// 	var octet byte = 0b01000001
// 	swapped := SwapBits(octet)

// 	fmt.Printf("Original byte: %08b\n", octet)
// 	fmt.Printf("Swapped byte:   %08b\n", swapped)
// }
// ################################## PIG LATIN ####################################
// func main(){
// 	args := os.Args[1:]
// 	if len(args) != 1 {
// 		return
// 	}
// 	s := args[0]
// 	if containsVow(s){
// 		res := piglatin(s)
// 	fmt.Println(res)
// 	}else {
// 		fmt.Println("No Vowels")
// 	}
	
// }
// func piglatin(s string)string{
// 	res := ""
// 	for i, v := range s {
// 		if i == 0 && isVowel(v) {
// 			res += s + "ay"
// 		}else if isVowel(v){
// 			res += s[i:] + s[:i] + "ay"
// 		}

// 	}
// 	return res
// }
// func isVowel(r rune)bool{
// 	if r == 'a' ||  r == 'e' ||  r == 'i' ||  r == 'o' ||  r == 'u' || r == 'A' ||  r == 'E' ||  r == 'I' ||  r == 'O' ||  r == 'U' {
// 		return true
// 	}
// 	return false
// }
// func containsVow(s string)bool{
// 	var vowels map[string]bool = map[string]bool{
// 		"a": true ,
// 		"e": true ,
// 		"i": true ,
// 		"o": true ,
// 		"u": true ,
// 		"A": true ,
// 		"E": true ,
// 		"I": true ,
// 		"O": true ,
// 		"U": true ,
// 	}
// 	for _, v := range s {
// 		if vowels[string(v)]{
// 			return true
// 		}
// 	}
// 	return false
// }
// ########################   ROMAN NUMERALS ##################################
// type romanNumerals []struct {
// 	value   int
// 	numeral string
// 	calc    string
// }
// func toRoman(x int) (string, string) {
// 	romanNumerals := romanNumerals{
// 		{1000, "M", "M"},
// 		{900, "CM", "(M-C)"},
// 		{500, "D", "D"},
// 		{400, "CD", "(D-C)"},
// 		{100, "C", "C"},
// 		{90, "XC", "(X-C)"},
// 		{50, "L", "L"},
// 		{40, "XL", "(L-X)"},
// 		{10, "X", "X"},
// 		{9, "IX", "(X-I)"},
// 		{5, "V", "V"},
// 		{4, "IV", "(V-I)"},
// 		{1, "I", "I"},
// 	}

// 	if x > 3999 {
// 		fmt.Println("ERROR: cannot convert to roman digit")
// 	}
// 	result := ""
// 	calculation := ""
// 	for _, rn := range romanNumerals {
// 		for x >= rn.value {
// 			result += rn.numeral
// 			if calculation != "" {
// 				calculation += "+"
// 			}
// 			calculation += rn.calc
// 			x -= rn.value
// 		}
// 	}
// 	return result, calculation
// }

// func main() {
// 	args := os.Args[1:]
// 	if len(args) > 1 {
// 		fmt.Println("ERROR: cannot convert to roman digit")
// 	}
// 	num, err := Atoi(args[0])
// 	if err != nil {
// 		fmt.Println("ERROR: cannot convert to roman digit")
// 	} else {
// 		if num > 3999 {
// 			fmt.Println("ERROR: cannot convert to roman digit")
// 		} else {
// 			roman, calculation := toRoman(num)
// 			fmt.Println(calculation)
// 			fmt.Println(roman)
// 		}

// 	}

// }
// ###################   PRINT BITS    #####################################
// func printBits(x int){
// 	res := ""
// 	if x == 1 {
// 		res += string(rune((x) + '0'))
// 	}else if x == 0 {
// 		res += string(rune((x) + '0'))
// 	}else {
// 		for x > 0 {
// 			var n rune = rune((x%2) + '0')
// 			res += string(n)
// 			x /= 2
// 		}
// 	}

// 	for len(res) < 8 {
// 		res += "0"
// 		if len(res) == 8 {
// 			break
// 		}
// 	}
// 	for i := len(res)-1; i >= 0; i-- {
// 		z01.PrintRune(rune(res[i]))
// 	}
// 	z01.PrintRune('\n')

// }
// func checkLength(s string)string{
// 	str := s
// 	if len(s) < 8 {
// 		str += "0"
// 	}
// 	return checkLength(str)
// }
// func main(){

// 	args := os.Args[1:]
// 	if len(args) != 1 {
// 		return
// 	}
// 	res, err := Atoi(args[0])
// 	if err != nil {
// 	  printBits(0)
// 	}
// 	printBits(res)
// resi := checkLength(resu)

// }
// ##########################    IS POWER OF 2 ######################@@@##########@@
// func isPowerof2(x int) bool {
// 	if x == 1 {
// 		return true
// 	}
// 	for x >= 2 {
// 		if x%2 != 0 {
// 			return false
// 		}
// 		x /= 2
// 	}
// 	return x == 1
// }
// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 1 {
// 		return
// 	}
// 	n, err := Atoi(args[0])
// 	if err != nil {
// 		fmt.Printf("%v\n", err)
// 		return
// 	} else {
// 		res := isPowerof2(n)
// 		fmt.Println(res)
// 	}

// }

// ###################################    INTER  ###########################
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
// 	var letters2 map[rune]bool = make(map[rune]bool)
// 	for _, v := range s2 {
// 		if letters1[v] && !letters2[v]{
// 			letters2[v] = true
// 		}
// 	}
// 	for _, v := range s1 {
// 		if letters1[v]{
// 			z01.PrintRune(v)
// 			letters1[v] = false
// 		}
// 	}
// 	z01.PrintRune('\n')
// }
// @@@@@@@@###############################     SEARCH REPLACE ##########################################3
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
//##################################   FOLD INT   #################################################333
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
// @############################     DOOP   ##################################################
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

// @######################################     ROT 13    ###########################################

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
// }

// ########################################   COMPARE    ############################################3
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
// ################################# ALPHA MIRROR   #################################################3
// func alphaMir(){
// 	args := os.Args[1:]
// 	if len(args) == 0 {
// 		return
// 	}
// 	s := args[0]
// 	var str string
// 	for _, v := range s {
// 		if (v >= 'A' && v <= 'Z')  {
// 			v = 'A' + ('Z' - v)
// 		}else if (v >= 'a' && v <= 'z') {
// 			v = 'a' + ('z' - v)
// 		}
// 		str += string(v)
// 	}
// 	fmt.Println(str)
// }
// func main(){
// 	alphaMir()
// }

// ############################################   REDUCE INT ####################################3
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

// ###################################    CHUNK     #############################################3
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

// #############################3 ITOA ATOI #######################
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

// func Atoi(s string) (int, error) {
// 	var isneg bool
// 	result := 0
// 	for i, v := range s {
// 		if i == 0 && v == '-' {
// 			isneg = true
// 		} else if v < '0' || v > '9' {
// 			return 0, fmt.Errorf("error converting; invalid number")
// 		} else if v >= '0' && v <= '9' {
// 			digit := int(v - '0')
// 			result = result*10 + digit
// 		}
// 	}
// 	if isneg {
// 		result = -result
// 	}
// 	return result, nil
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