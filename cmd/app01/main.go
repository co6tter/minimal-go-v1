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

	score := 80
	switch {
	case score >= 90:
		fmt.Println("A!")
	case score >= 70:
		fmt.Println("B!")
	default:
		fmt.Println("C!")
	}

	score1a := 3
	switch score1a {
	case 0:
		fmt.Println("Fail!")
	case 1:
		fmt.Println("C!")
	case 2:
		fmt.Println("B!")
	case 3, 4, 5:
		fmt.Println("A!")
	default:
		fmt.Println("Invalid score!")
	}

	money := 100.0
	for year := 1; year <= 30; year++ {
		money *= 1.05
		if money > 300 {
			break
		}
		if money < 200 {
			continue
		}
		fmt.Printf("Year %d: %f\n", year, money)
	}
}
