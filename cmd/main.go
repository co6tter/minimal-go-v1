package main

import "fmt"

func main() {
	fmt.Println(10 / 3)   // 3
	fmt.Println(10.0 / 3) // 3.3333333333333335

	var price1a int // ゼロ値
	fmt.Println(price1a)
	price1a = 150
	fmt.Println(price1a * 120)

	var price1b = 150
	fmt.Println(price1b)

	price1c := 150
	fmt.Println(price1c)

	// var rate float64
	// rate = 1.1
	// rate := 1.1
	const rate = 1.1 // 定数
	fmt.Println(float64(price1c) * float64(120) * rate)

	var name string
	fmt.Println("Your name?")
	fmt.Scanf("%s", &name)
	fmt.Printf("Hi, %s\n", name)
}
