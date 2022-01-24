package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
)

type Item struct {
	name     string  "상품 이름"
	price    float64 "상품 가격"
	quantity int     "구매 수량"
}

func (t Item) Cost() float64 {
	return t.price * float64(t.quantity)
}

func printShirt() {
	shirt := Item{"Men's Slim-Fit Shirt", 25000, 3}
	fmt.Println("===")
	fmt.Println(shirt.Cost())
	fmt.Println("===")
}

func printType1() {
	type quantity int
	type dozen []quantity
	var d dozen
	for i := quantity(1); i <= 12; i++ {
		d = append(d, i)
	}
	fmt.Println("===")
	fmt.Println(d)
	fmt.Println("===")
}

func display1(i int) {
	fmt.Println("===")
	fmt.Println(i)
	fmt.Println("===")
}

func display2(i quantity) {
	fmt.Println("===")
	fmt.Println(i)
	fmt.Println("===")
}

type quantity int

func printType2() {
	var q quantity = 3
	display1(int(q))
}

func printType3() {
	i := 3
	display2(quantity(i))
}

type costCalculator func(quantity, float64) float64

func describe1(q quantity, price float64, c costCalculator) {
	fmt.Printf("quantity: %d, price: %0.0f, cost: %0.0f\n", q, price, c(q, price))
}

func printDescribe() {
	var offBy10Percent, offBy1000Won costCalculator
	offBy10Percent = func(q quantity, price float64) float64 {
		return float64(q) * price * 0.9
	}
	offBy1000Won = func(q quantity, price float64) float64 {
		return float64(q)*price - 1000
	}
	fmt.Println("===")
	describe1(3, 10000, offBy10Percent)
	describe1(3, 10000, offBy1000Won)
	fmt.Println("===")
}

type rect struct {
	width  float64
	height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func printRect1() {
	r := rect{3, 4}
	fmt.Println("===")
	fmt.Println("area : ", r.area())
	fmt.Println("===")
}

type shaper interface {
	area() float64
}

func describe2(s shaper) {
	fmt.Println("area : ", s.area())
}

func printRect2() {
	r := rect{3, 4}
	fmt.Println("===")
	describe2(r)
	fmt.Println("===")
}

func (q quantity) greaterThan(i int) bool {
	return int(q) > i
}

func (q *quantity) increment() {
	*q++
}

func (q *quantity) decrement() {
	*q--
}

func printQuantity() {
	q := quantity(3)
	q.increment()
	fmt.Println("===")
	fmt.Printf("Is q(%d) greater than %d? %t \n", q, 3, q.greaterThan(3))
	q.decrement()
	fmt.Printf("Is q(%d) greater than %d? %t \n", q, 3, q.greaterThan(3))
	fmt.Println("===")
}

type numberMap map[string]int

func (m numberMap) add(key string, value int) {
	m[key] = value
}

func (m numberMap) remove(key string) {
	delete(m, key)
}

func printNumberMap() {
	m := make(numberMap)
	m["one"] = 1
	m["two"] = 2
	m.add("three", 3)
	fmt.Println("===")
	fmt.Println(m)
	m.remove("two")
	fmt.Println(m)
	fmt.Println("===")
}

func (rect) new() rect {
	return rect{}
}

func printRect3() {
	r := rect{}.new()
	fmt.Println("===")
	fmt.Println(r)
	fmt.Println("===")
}

func (r *rect) resize(w, h float64) {
	r.width += w
	r.height += h
}

func printRect4() {
	r := rect{3, 4}
	fmt.Println("===")
	fmt.Println("area : ", r.area())
	r.resize(10, 10)
	fmt.Println("area : ", r.area())
	areaFn := rect.area
	resizeFn1 := (*rect).resize
	//resizeFn2 := rect.resize // 에러 발생 : 파라메터가 있는 경우에는 포인터로 작성해야함.
	fmt.Println("area : ", areaFn(r))
	resizeFn1(&r, -10, -10)
	fmt.Println("area : ", areaFn(r))
	//resizeFn2(r, 10, 10)
	//fmt.Println("area : ", areaFn(r))
	fmt.Println("===")
}

func printItem1() {
	shirt := Item{"Men's Slim-Fit Shirt", 25000, 3}
	shoes := Item{"Sports Shoes", 30000, 0}
	dress := Item{"Stripe Shift Dress", 0, 2}
	phone := Item{"Amazon Fire Phone, 32GB", 21900, 1}
	fmt.Println("===")
	fmt.Println(shirt)
	fmt.Println(shoes)
	fmt.Println(dress)
	fmt.Println(phone)
	fmt.Println("===")
}

func printItem2() {
	p := &Item{"Men's Slim-Fit Shirt", 25000, 3}
	fmt.Println("===")
	fmt.Println(p)
	fmt.Println(p.Cost())
	fmt.Println("===")
}

func printItem3() {
	item := new(Item)
	item.name = "Men's Slim-Fit Shirt"
	item.price = 25000
	item.quantity = 3
	fmt.Println("===")
	fmt.Println(item)
	fmt.Println(item.Cost())
	fmt.Println("===")
}

func printRect5() {
	r1 := rect{1, 2}
	r1.width, r1.height = 2, 1
	r2 := new(rect)
	r2.width, r2.height = 3, 4
	r3 := &rect{}
	r4 := &rect{5, 6}
	fmt.Println("===")
	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
	fmt.Println(r4)
	fmt.Println("===")
}

func printRect6() {
	rects := []struct{ w, h int }{{1, 2}, {}, {-1, -2}}
	fmt.Println("===")
	for _, r := range rects {
		fmt.Printf("(%d, %d)", r.w, r.h)
	}
	fmt.Println("\n===")
}

func printItem4() {
	var t Item
	t.name = "Men's Slim-Fit Shirt"
	t.price = 25000
	t.quantity = 3
	fmt.Println("===")
	fmt.Println(t.name)
	fmt.Println(t.price)
	fmt.Println(t.quantity)
	fmt.Println(t.Cost())
	fmt.Println("===")
}

type dimension struct {
	width, height, length float64
}

type Item2 struct {
	name             string
	price            float64
	quantity         int
	packageDimension dimension
	itemDimension    dimension
}

func printItem5() {
	shoes := Item2{
		"Sports Shoes", 30000, 2,
		dimension{30, 270, 20},
		dimension{50, 300, 30},
	}
	fmt.Println("===")
	fmt.Printf("%#v\n", shoes.itemDimension)
	fmt.Printf("%#v\n", shoes.packageDimension)
	fmt.Println(shoes.packageDimension.width)
	fmt.Println(shoes.packageDimension.height)
	fmt.Println(shoes.packageDimension.length)
	fmt.Println("===")
}

func printItem6() {
	tType := reflect.TypeOf(Item{})
	fmt.Println("===")
	for i := 0; i < tType.NumField(); i++ {
		fmt.Println(tType.Field(i).Tag)
	}
	fmt.Println("===")
}

type Option struct {
	name  string
	value string
}

type Item3 struct {
	name     string
	price    float64
	quantity int
	Option
}

func printItem7() {
	shoes := Item3{"Sports Shoes", 30000, 2, Option{"color", "red"}}
	fmt.Println("===")
	fmt.Println(shoes)
	fmt.Println(shoes.name)
	fmt.Println(shoes.value)
	fmt.Println(shoes.Option.name)
	fmt.Println("===")
}

type DiscountItem struct {
	Item
	discountRate float64
}

func printDiscountItem1() {
	shoes := Item{"Women's Walking Shoes", 30000, 2}
	eventShoes := DiscountItem{
		Item{"Sports Shoes", 5000, 3},
		10.00,
	}
	fmt.Println("===")
	fmt.Println(shoes.Cost())
	//fmt.Println(eventShoes.Cost()) // DiscountItem.Cost() 메서드 추가 이후 원가 표시가 되지 않음
	fmt.Println(eventShoes.Item.Cost())
	fmt.Println("===")
}

func (t DiscountItem) Cost() float64 {
	return t.Item.Cost() * (1.0 - t.discountRate/100)
}

func printDiscountItem2() {
	shoes := Item{"Women's Walking Shoes", 30000, 2}
	eventShoes := DiscountItem{
		Item{"Sports Shoes", 5000, 3},
		10.00,
	}
	fmt.Println("===")
	fmt.Println(shoes.Cost())
	fmt.Println(eventShoes.Cost())
	fmt.Println(eventShoes.Item.Cost())
	fmt.Println("===")
}

func NewItem(name string, price float64, quantity int) *Item {
	if price <= 0 || quantity <= 0 || len(name) == 0 {
		return nil
	}
	return &Item{name, price, quantity}
}

func printNewItem() {
	t := NewItem("Men's Slim-Fit Shirt", 25000, 3)
	fmt.Println("===")
	fmt.Println(t)
	fmt.Println("===")
}

func (t *Item) Name() string {
	return t.name
}

func (t *Item) SetName(n string) {
	if len(n) != 0 {
		t.name = n
	}
}

func (t *Item) Price() float64 {
	return t.price
}

func (t *Item) SetPrice(p float64) {
	if p > 0 {
		t.price = p
	}
}

func (t *Item) Quantity() int {
	return t.quantity
}

func (t *Item) SetQuantity(q int) {
	if q > 0 {
		t.quantity = q
	}
}

func printItem8() {
	shirt := NewItem("Men's Slim-Fit Shirt", 25000, 3)
	fmt.Println("===")
	fmt.Println(shirt)
	shirt.SetPrice(30000)
	shirt.SetQuantity(2)
	fmt.Println("Name: ", shirt.Name())
	fmt.Println("Price: ", shirt.Price())
	fmt.Println("Quantity: ", shirt.Quantity())
	fmt.Println("===")
}

func display3(s interface{ show() }) {
	s.show()
}

func (r rect) show() {
	fmt.Printf("width: %f, height: %f\n", r.width, r.height)
}

type circle struct{ radius float64 }

func (c circle) show() {
	fmt.Printf("radius: %f\n", c.radius)
}

func printCircle1() {
	r := rect{3, 4}
	c := circle{2.5}
	fmt.Println("===")
	display3(r)
	display3(c)
	fmt.Println("===")
}

func display4(s interface{}) {
	fmt.Println(s)
}

func printCircle2() {
	r := rect{3, 4}
	c := circle{2.5}
	fmt.Println("===")
	display4(r)
	display4(c)
	display4(2.5)
	display4("rect struct")
	fmt.Println("===")
}

type Coster interface {
	Cost() float64
}

func displayCost(c Coster) {
	fmt.Println("cost: ", c.Cost())
}

func printCost1() {
	shoes := Item{"Sports Shoes", 30000, 2}
	evnetShoes := DiscountItem{
		Item{"Women's Walking Shoes", 50000, 3},
		10.00,
	}
	fmt.Println("===")
	displayCost(shoes)
	displayCost(evnetShoes)
	fmt.Println("===")
}

type Rental struct {
	name         string
	feePerDay    float64
	periodLength int
	RentalPeriod
}

type RentalPeriod int

const (
	Days RentalPeriod = iota
	Weeks
	Months
)

func (p RentalPeriod) ToDays() int {
	switch p {
	case Weeks:
		return 7
	case Months:
		return 30
	default:
		return 1
	}
}

func (r Rental) Cost() float64 {
	return r.feePerDay * float64(r.ToDays()*r.periodLength)
}

func printRental1() {
	shirt := Item{"Men's Slim-Fit Shirt", 25000, 3}
	video := Rental{"Intersteller", 1000, 3, Days}
	fmt.Println("===")
	displayCost(shirt)
	displayCost(video)
	fmt.Println("===")
}

type Items []Coster

func (ts Items) Cost() (c float64) {
	for _, t := range ts {
		c += t.Cost()
	}
	return
}

func printCost2() {
	shirt := Item{"Men's Slim-Fit Shirt", 25000, 3}
	video := Rental{"Interstellar", 1000, 3, Days}
	eventShoes := DiscountItem{
		Item{"Women's Walking Shoes", 50000, 3},
		10.00,
	}
	items := Items{shirt, video, eventShoes}
	fmt.Println("===")
	displayCost(items)
	fmt.Println("===")
}

func (t Item) String() string {
	return fmt.Sprintf("[%s] %.0f", t.name, t.Cost())
}

func (t DiscountItem) String() string {
	return fmt.Sprintf("%s => %.0f(%.0f%s DC)", t.Item.String(), t.Cost(), t.discountRate, "%")
}

func (t Rental) String() string {
	return fmt.Sprintf("[%s] %.0f", t.name, t.Cost())
}

func (ts Items) String() string {
	var s []string
	for _, t := range ts {
		s = append(s, fmt.Sprint(t))
	}
	return fmt.Sprintf("%d items. total: %.0f\n\t- %s", len(ts), ts.Cost(), strings.Join(s, "\n\t- "))
}

func printStringer1() {
	shirt := Item{"Men's Slim-Fit Shirt", 25000, 3}
	video := Rental{"Interstellar", 1000, 3, Days}
	eventShoes := DiscountItem{
		Item{"Women's Walking Shoes", 50000, 3},
		10.00,
	}
	items := Items{shirt, video, eventShoes}
	fmt.Println("===")
	fmt.Println(shirt)
	fmt.Println(video)
	fmt.Println(eventShoes)
	fmt.Println(items)
	fmt.Println("===")
}

func handle(w io.Writer, msg string) {
	fmt.Fprintln(w, msg)
}

func printWriter1() {
	msg := []string{"This", "is", "an", "example", "of", "io.Writer"}
	fmt.Println("===")
	for _, s := range msg {
		time.Sleep(100 * time.Millisecond)
		handle(os.Stdout, s)
	}
	fmt.Println("===")
}

// 브라우저에서 http://localhost:4000 으로 접속하여 테스트한다.
// http://localhost:4000/aaa => 브라우저 화면에 aaa 가 표시됨
func printWriter2() {
	fmt.Println("===")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handle(w, r.URL.Path[1:])
	})
	fmt.Println("start listening on prot 4000")
	http.ListenAndServe(":4000", nil)
	fmt.Println("===")
}

type Itemer interface {
	Coster
	fmt.Stringer
}

type Order struct {
	Itemer
	taxRate float64
}

func (o Order) Cost() float64 {
	return o.Itemer.Cost() * (1.0 + o.taxRate/100)
}

func (o Order) String() string {
	return fmt.Sprintf("Total price: %.0f(tax rate: %.2f)\n\tOrder details: %s", o.Cost(), o.taxRate, o.Itemer.String())
}

func printOrder1() {
	shirt := Item{"Men's Slim-Fit Shirt", 25000, 3}
	video := Rental{"Interstellar", 1000, 3, Days}
	eventShoes := DiscountItem{
		Item{"Women's Walking Shoes", 50000, 3},
		10.00,
	}
	order1 := Order{Items{shirt, eventShoes}, 10.00}
	order2 := Order{video, 5.00}
	fmt.Println("===")
	fmt.Println(order1)
	fmt.Println(order2)
	fmt.Println("===")
}

func checkType(v interface{}) {
	switch v.(type) {
	case bool:
		fmt.Printf("%t is a bool\n", v)
	case int, int8, int16, int32, int64:
		fmt.Printf("%d is an int\n", v)
	case uint, uint8, uint16, uint32, uint64:
		fmt.Printf("%d is an unsigned int\n", v)
	case float64:
		fmt.Printf("%f is a float64\n", v)
	case complex64, complex128:
		fmt.Printf("%f is a complex\n", v)
	case string:
		fmt.Printf("%s is a string\n", v)
	case nil:
		fmt.Printf("%v is nil\n", v)
	default:
		fmt.Printf("%v is unknown type\n", v)
	}
}

func printType4() {
	fmt.Println("===")
	checkType(3)
	checkType(1.5)
	checkType(complex(1, 5))
	checkType(true)
	checkType("s")
	checkType(struct{}{})
	checkType(nil)
	fmt.Println("===")
}

func main() {
	printShirt()
	printType1()
	printType2()
	printType3()
	printDescribe()
	printRect1()
	printRect2()
	printQuantity()
	printNumberMap()
	printRect3()
	printRect4()
	printItem1()
	printItem2()
	printItem3()
	printRect5()
	printRect6()
	printItem4()
	printItem5()
	printItem6()
	printItem7()
	printDiscountItem1()
	printDiscountItem2()
	printNewItem()
	printItem8()
	printCircle1()
	printCircle2()
	printCost1()
	printRental1()
	printCost2()
	printStringer1()
	printWriter1()
	//printWriter2() // 브라우저 테스트는 일반 테스트 시 주석 처리
	printOrder1()
	printType4()
}
