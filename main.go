package main

import (
	"fmt"
	"runtime"
)

type Student struct {
	id   int
	name string
}

type ProductInfo struct {
	price  string
	origin string
}

type Product struct {
	id   int
	name string
	info ProductInfo
}

func main() {
	fmt.Println("==================================")
	fmt.Println("Hello Golang..........")
	fmt.Println("==================================")
	showInfo()
	fmt.Println("==================================")
	fmt.Println(sum(3))
	fmt.Println("==================================")
	// var a int;
	// a = 3;
	var a = 3
	// var (
	// 	name    string = "a"
	// 	address string = "b"
	// 	age     int    = 20
	// )
	// var name, address, age = "nameA", "adrress1", 20
	fmt.Println(a)
	fmt.Println("==================================")
	const PI = 3.14
	fmt.Println("PI...... is ", PI)
	fmt.Println("==================================")
	var sum = 0
	for i := 0; i < 4; i++ {
		sum += i
	}
	fmt.Println("sum is :", sum)
	// for ; sum < 1000; {
	// 	sum += sum
	// }
	fmt.Println("==================================")
	// not use while
	test := 1
	for test < 4 {
		test += test
	}
	fmt.Println("Test is :", test)
	fmt.Println("==================================")
	// for {
	// 	fmt.Println("infinite loop")
	// }
	fmt.Println("==================================")
	max := 10
	v := 9
	if value := 3; v < max {
		fmt.Println(value)
	} else {
		fmt.Println(max)
	}
	fmt.Println("==================================")
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}
	fmt.Println("==================================")
	w, h := multiReturnValues(100, 200)
	fmt.Printf("with and hight : %d, %d", w, h)
	fmt.Println("==================================")

	st1 := Student{
		id:   123,
		name: "Ngoc Hung",
	}
	fmt.Println(st1)
	fmt.Println(st1.id)
	fmt.Println(st1.name)
	st2 := Student{456, "Hung Hoang"}
	fmt.Println(st2)
	// anonymous struct
	var anonymous = struct {
		address string
	}{
		address: "struct annonymous",
	}
	fmt.Println(anonymous)
	// ananymous filds
	type NoName struct {
		string
		int
	}
	fmt.Println("==================================")
	n := NoName{"Quang Tri", 74}
	fmt.Println(n)
	fmt.Println("==================================")
	// struct in struct =- nested struct
	product := Product{
		id:   1,
		name: "Milo",
		info: ProductInfo{
			price:  "563 VND",
			origin: "Quang Tri",
		},
	}
	fmt.Println(product)
}

func showInfo() {
	fmt.Println("call function .................")
}

// func showInfo(name string) {
// 	fmt.Println("test overload .................")
// }

func multiReturnValues(w, h int) (int, int) {
	return w, h
}

func sum(number int) int {
	sum := 0
	for i := 0; i < number; i++ {
		sum += i
	}
	return sum
}
