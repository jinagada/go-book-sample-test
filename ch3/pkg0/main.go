package main

import (
	"bytes"
	"fmt"
	"image/color"
	"math"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

func checkValue(v interface{}) {
	//if v { // 빌드 오류
	if v != nil { // 정상
		fmt.Printf("value is %v\n", v)
		return
	}
	fmt.Println("value is nil")
}

func printNumber() {
	a := 365   // 10진수
	b := 0555  // 8진수
	c := 0x16D // 16진수
	fmt.Println("===")
	fmt.Println(a, b, c)
	fmt.Println("===")
}

func printAscii() {
	var ch1 byte = 65   // 10진수
	var ch2 byte = 0101 // 8진수
	var ch3 byte = 0x41 // 16진수
	fmt.Println("===")
	fmt.Printf("%c %c %c\n", ch1, ch2, ch3)
	fmt.Println("===")
}

func printUnicode() {
	var ch4 rune = 44032   // 10진수
	var ch5 rune = 0126000 // 8진수
	var ch6 rune = 0xAC00  // 16진수
	fmt.Println("===")
	fmt.Printf("%c %c %c\n", ch4, ch5, ch6)
	fmt.Println("===")
}

func printChar() {
	var ch1 byte = 'A'
	var ch2 rune = '가'
	fmt.Println("===")
	fmt.Printf("%c %c\n", ch1, ch2)
	fmt.Println("===")
}

func printComplex() {
	c1 := 1 + 2i            // complex128
	c2 := complex64(3 + 4i) // complex64
	c3 := complex(5, 6)     // complex128
	fmt.Println("===")
	fmt.Println(c1, real(c1), imag(c1))
	fmt.Println(c2, real(c2), imag(c2))
	fmt.Println(c3, real(c3), imag(c3))
	fmt.Println("===")
}

func printOperation() {
	i := 100000
	j := int16(10000)
	k := uint8(100)
	fmt.Println("===")
	fmt.Println(i + int(j))
	fmt.Println(i + int(k))
	fmt.Println(j > int16(k))
	fmt.Println(int16(i))
	fmt.Println(uint8(i))
	fmt.Println("===")
}

func intToUint8(i int) (uint8, error) {
	if 0 <= i && i <= math.MaxUint8 {
		return uint8(i), nil
	}
	return 0, fmt.Errorf("%d cannot convert to uint8 type", i)
}

func printConvert() {
	fmt.Println("===")
	fmt.Println(intToUint8(100))
	fmt.Println(intToUint8(1000))
	fmt.Println("===")
}

func printChar2() {
	var ch int = '\u0041'
	var ch2 int = '\uAC00'
	var ch3 int = '\U00101234'
	fmt.Println("===")
	fmt.Printf("%8d - %8d - %8d\n", ch, ch2, ch3) // 정수
	fmt.Printf("%8c - %8c -%8c\n", ch, ch2, ch3)  // 문자
	fmt.Printf("%8X - %8X - %8X\n", ch, ch2, ch3) // UTF-8 바이트 수
	fmt.Printf("%8U - %8U - %8U\n", ch, ch2, ch3) // UTF-8 코드값
	fmt.Println("===")
}

func printChar3() {
	s := "hello"
	fmt.Println("===")
	fmt.Println(s[0])
	fmt.Println(s[len(s)-1])
	fmt.Println("===")
}

func printChar4() {
	s1 := "hello"
	s2 := "안녕하세요"
	fmt.Println("===")
	fmt.Printf("s1: %c %c %c %c %c\n", s1[0], s1[1], s1[2], s1[3], s1[4])
	fmt.Printf("s2: %c %c %c %c %c\n", s2[0], s2[1], s2[2], s2[3], s2[4])
	fmt.Println("===")
}

func printChar5() {
	s1 := "hello"
	s2 := "안녕하세요"
	r1 := []rune(s1)
	r2 := []rune(s2)
	fmt.Println("===")
	fmt.Printf("s1: %c %c %c %c %c\n", r1[0], r1[1], r1[2], r1[3], r1[4])
	fmt.Printf("s2: %c %c %c %c %c\n", r2[0], r2[1], r2[2], r2[3], r2[4])
	fmt.Println("===")
}

func printChar6() {
	s1 := "hello"
	s2 := "안녕하세요"
	fmt.Println("===")
	for i, c := range s1 {
		fmt.Printf("%c(%d)\t", c, i)
	}
	fmt.Println()
	for i, c := range s2 {
		fmt.Printf("%c(%d)\t", c, i)
	}
	fmt.Println("===")
}

func printChar7() {
	s1 := "hello"
	fmt.Println("===")
	fmt.Println([]rune(s1))
	fmt.Println([]byte(s1))
	fmt.Println(string([]byte{104, 101, 108, 108, 111}))
	fmt.Println(string([]rune{104, 101, 108, 108, 111}))
	s2 := "안녕하세요"
	fmt.Println([]rune(s2))
	fmt.Println([]byte(s2))
	fmt.Println(string([]rune{50504, 45397, 54616, 49464, 50836}))
	fmt.Println(string([]byte{236, 149, 136, 235, 133, 149, 237, 149, 152, 236, 132, 184, 236, 154, 148}))
	fmt.Println(string(104))
	fmt.Println(string(236))
	fmt.Println(string(50504))
	fmt.Println(string([]byte{236, 149, 136}))
	fmt.Println("===")
}

func printLength() {
	s1 := "hello"
	s2 := "안녕하세요"
	fmt.Println("===")
	fmt.Println(len(s1))
	fmt.Println(len(s2))
	fmt.Println(utf8.RuneCountInString(s2))
	fmt.Println(len([]rune(s2)))
	fmt.Println("===")
}

func printChar8() {
	s := "hello"
	fmt.Println("===")
	fmt.Println(s[1:2])
	fmt.Println(s[1:])
	fmt.Println(s[:2])
	s = "안녕하세요"
	fmt.Println(s[1:2])
	fmt.Println(s[1:])
	fmt.Println(s[:2])
	fmt.Println("===")
}

func printDiff() {
	s1 := "hello"
	s2 := "안녕하세요"
	fmt.Println("===")
	fmt.Println(s1 == s2)
	fmt.Println(s1 != s2)
	fmt.Println(s1 > s2)
	fmt.Println(s1 < s2)
	fmt.Println("===")
}

func printStr1() {
	text := "Go is an open source programming language" +
		" that makes it easy to build simple, reliable, and efficient software."
	text += " Go is expressive, concise, clean, and efficient."
	fmt.Println("===")
	fmt.Println(text)
	fmt.Println("===")
}

func printStr2() {
	var str string
	for i := 0; i < math.MaxUint8; i++ {
		if s, ok := nextString(i); ok {
			str += s
		}
	}
	fmt.Println("===")
	fmt.Print(str, "\n")
	fmt.Println("===")
}

func nextString(i int) (s string, ok bool) {
	if unicode.IsLetter(rune(i)) {
		return string(i), true
	}
	return "", false
}

func printStr3() {
	var strArr []string
	for i := 0; i < math.MaxUint8; i++ {
		if s, ok := nextString(i); ok {
			strArr = append(strArr, s)
		}
	}
	fmt.Println("===")
	fmt.Println(strings.Join(strArr, ""))
	fmt.Println("===")
}

func printStr4() {
	var buffer bytes.Buffer
	for i := 0; i < math.MaxUint8; i++ {
		if s, ok := nextString(i); ok {
			buffer.WriteString(s)
		}
	}
	fmt.Println("===")
	fmt.Println(buffer.String())
	fmt.Println("===")
}

func printArray1() {
	var a [5]int
	b := [3]int{1, 2, 3}
	c := [3]int{1, 2}
	d := [...]int{4, 5, 6, 7, 8}
	e := [3][3]int{
		{1, 2, 3},
		{3, 4, 5},
	}
	fmt.Println("===")
	fmt.Printf("%-10T %d %v\n", a, len(a), a)
	fmt.Printf("%-10T %d %v\n", b, len(b), b)
	fmt.Printf("%-10T %d %v\n", c, len(c), c)
	fmt.Printf("%-10T %d %v\n", d, len(d), d)
	fmt.Printf("%-10T %d %v\n", e, len(e), e)
	fmt.Println("===")
}

func printSlice1() {
	var a []int
	b := []int{}
	c := []int{1, 2, 3}
	d := [][]int{
		{1, 2},
		{3, 4, 5},
	}
	e := make([]int, 0)
	f := make([]int, 3, 5)
	fmt.Println("===")
	fmt.Printf("%-10T %d %d %v\n", a, len(a), cap(a), a)
	fmt.Printf("%-10T %d %d %v\n", b, len(b), cap(b), b)
	fmt.Printf("%-10T %d %d %v\n", c, len(c), cap(c), c)
	fmt.Printf("%-10T %d %d %v\n", d, len(d), cap(d), d)
	fmt.Printf("%-10T %d %d %v\n", e, len(e), cap(e), e)
	fmt.Printf("%-10T %d %d %v\n", f, len(f), cap(f), f)
	fmt.Println("===")
}

func printForLoop1() {
	numbers := []int{3, 4, 5, 7, 8, 4, 6, 8, 32, 4}
	fmt.Println("===")
	for index, value := range numbers {
		fmt.Println(index, value)
	}
	fmt.Println("===")
}

func printForLoop2() {
	numbers := []int{3, 4, 5, 7, 8, 4, 6, 8, 32, 4}
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	fmt.Println("===")
	fmt.Println(sum)
	fmt.Println("===")
}

func printForLoop3() {
	numbers := []int{3, 4, 5, 7, 8, 4, 6, 8, 32, 4}
	sum := 0
	for i := range numbers {
		numbers[i] *= 2
		sum += numbers[i]
	}
	fmt.Println("===")
	fmt.Println(numbers)
	fmt.Println(sum)
	fmt.Println("===")
}

func printSlice2() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("===")
	fmt.Println(s, "≡", s[:3], s[3:5], s[5:])
	fmt.Println("===")
}

func printSlice3() {
	ns1 := []int{1, 2, 3}
	ns2 := []int{6, 7, 8}
	ns3 := []int{8, 9, 10, 11}
	fmt.Println("===")
	ns1 = append(ns1, 4, 5)
	fmt.Println(ns1)
	ns1 = append(ns1, ns2...)
	fmt.Println(ns1)
	ns1 = append(ns1, ns3[1:3]...)
	fmt.Println(ns1)
	fmt.Println("===")
}

func printSlice4() {
	s := make([]int, 0, 3)
	fmt.Println("===")
	for i := 0; i < 9; i++ {
		s = append(s, i)
		fmt.Printf("len: %d, cap: %d, %v\n", len(s), cap(s), s)
	}
	fmt.Println("===")
}

func insert(s, new []int, index int) []int {
	return append(s[:index], append(new, s[index:]...)...)
}

func printSlice5() {
	s := []int{1, 2, 3, 4, 5}
	s = insert(s, []int{-3, -2}, 0)
	fmt.Println("===")
	fmt.Println(s)
	s = insert(s, []int{0}, 2)
	fmt.Println(s)
	s = insert(s, []int{6, 7}, len(s))
	fmt.Println(s)
	fmt.Println("===")
}

func insert2(s, new []int, index int) []int {
	result := make([]int, len(s)+len(new))
	position := copy(result, s[:index])
	position += copy(result[position:], new)
	copy(result[position:], s[index:])
	return result
}

func printSlice6() {
	s := []int{1, 2, 3, 4, 5}
	s = insert2(s, []int{-3, -2}, 0)
	fmt.Println("===")
	fmt.Println(s)
	s = insert2(s, []int{0}, 2)
	fmt.Println(s)
	s = insert2(s, []int{6, 7}, len(s))
	fmt.Println(s)
	fmt.Println("===")
}

func printSort1() {
	s := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(s)
	fmt.Println("===")
	fmt.Println(s)
	fmt.Println("===")
}

func printMap1() {
	numberMap := map[string]int{}
	numberMap["one"] = 1
	numberMap["two"] = 2
	numberMap["three"] = 3
	fmt.Println("===")
	fmt.Println(numberMap)
	fmt.Println("===")
}

func printMap2() {
	numberMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println("===")
	fmt.Println(numberMap)
	fmt.Println("===")
}

func printMap3() {
	numberMap := make(map[string]int, 3)
	numberMap["one"] = 1
	numberMap["two"] = 2
	numberMap["three"] = 3
	fmt.Println("===")
	fmt.Println(numberMap)
	fmt.Println("===")
}

func printMap4() {
	numberMap := make(map[string]int)
	numberMap["one"] = 1
	numberMap["two"] = 2
	numberMap["three"] = 3
	fmt.Println("===")
	for k, v := range numberMap {
		fmt.Println(k, v)
	}
	fmt.Println("===")
}

func printMap5() {
	numberMap := map[string]int{}
	numberMap["zero"] = 0
	numberMap["one"] = 1
	numberMap["two"] = 2
	fmt.Println("===")
	fmt.Println(numberMap["zero"])
	fmt.Println(numberMap["one"])
	fmt.Println(numberMap["two"])
	fmt.Println(numberMap["three"])
	if v, ok := numberMap["three"]; ok {
		fmt.Println("'three' is in numberMap. value: ", v)
	} else {
		fmt.Println("'three' is not in numberMap")
	}
	if _, ok := numberMap["four"]; !ok {
		numberMap["four"] = 4
	}
	fmt.Println(numberMap)
	fmt.Println("===")
}

func printMap6() {
	numberMap := map[int]string{}
	numberMap[1] = "one"
	numberMap[2] = "two"
	fmt.Println("===")
	fmt.Println(numberMap)
	numberMap[3] = "three"
	fmt.Println(numberMap)
	numberMap[3] = "삼"
	fmt.Println(numberMap)
	delete(numberMap, 3)
	fmt.Println(numberMap)
	fmt.Println("===")
}

func printPointer1() {
	var p *int
	i := 1
	p = &i
	fmt.Println("===")
	fmt.Println(i)
	fmt.Println(&i)
	fmt.Println(*p)
	fmt.Println(p)
	fmt.Println("===")
}

func printPointer2() {
	var p *int
	var pp **int
	i := 1
	p = &i
	pp = &p
	fmt.Println("===")
	fmt.Println(i, *p, **pp)
	i += 1
	fmt.Println(i, *p, **pp)
	*p++
	fmt.Println(i, *p, **pp)
	**pp++
	fmt.Println(i, *p, **pp)
	fmt.Println("===")
}

func printPointer3() {
	type rect struct{ w, h float64 }
	//var i int = 1
	//var p *int = &i
	//var s *rect = &rect{1, 2}
	i := 1
	p := &i
	s := &rect{1, 2}
	fmt.Println("===")
	fmt.Println(p)
	fmt.Println(s)
	fmt.Println("===")
}

func printPointer4() {
	p := new(int)
	*p = 1
	fmt.Println("===")
	fmt.Println(p)
	fmt.Println(*p)
	type rect struct{ w, h float64 }
	r := new(rect)
	r.w, r.h = 3, 4
	fmt.Println(r)
	fmt.Println(*r)
	fmt.Println("===")
}

func multiply(numbers []int, factor int) {
	for i := range numbers {
		numbers[i] *= factor
	}
}

func printPointer5() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	multiply(numbers, 5)
	fmt.Println("===")
	fmt.Println(numbers)
	fmt.Println("===")
}

func printPointer6() {
	rect := rect{2, 4, 10, 20, color.RGBA{0xFF, 0, 0, 0xFF}}
	resize(&rect, 10, 10)
	fmt.Println("===")
	fmt.Println(rect)
	fmt.Println("===")
}

type rect struct {
	x0, y0, x1, y1 int
	color.RGBA
}

func resize(rect *rect, width, height int) {
	(*rect).x1 += width
	rect.y1 += height
}

func main() {
	printNumber()
	printAscii()
	printUnicode()
	printChar()
	printComplex()
	printOperation()
	printConvert()
	printChar2()
	printChar3()
	printChar6()
	printChar7()
	printChar4()
	printChar5()
	printLength()
	printChar8()
	printDiff()
	printStr1()
	printStr2()
	printStr3()
	printStr4()
	printArray1()
	printSlice1()
	printForLoop1()
	printForLoop2()
	printForLoop3()
	printSlice2()
	printSlice3()
	printSlice4()
	printSlice5()
	printSlice6()
	printSort1()
	printMap1()
	printMap2()
	printMap3()
	printMap4()
	printMap5()
	printMap6()
	printPointer1()
	printPointer2()
	printPointer3()
	printPointer4()
	printPointer5()
	printPointer6()
}
