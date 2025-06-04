package main

import "fmt"

type nameString string

func showProductInfo(product nameString) {
	fmt.Println(product)
}

func (n nameString) showInfo() {
	fmt.Println(n)
}

// func (n nameString) update(newName string) nameString {
// 	n = nameString(newName)
// 	return n
// }

func (n *nameString) update(newName string) {
	*n = nameString(newName)
}

// type studentType struct {
// 	name  string
// 	score scoreType
// }

// 埋め込みフィールド
// フィールドのメソッドにアクセスできるだけで同じデータ型として扱うことはできない
type studentType struct {
	name string
	scoreType
}

type scoreType struct {
	subject string
	points  int
}

func (s scoreType) isPassed() bool {
	if s.points >= 70 {
		return true
	} else {
		return false
	}
}

type englishScoreType int
type mathScoreType int

func (e englishScoreType) isPassed() bool {
	return e >= 80
}

func (m mathScoreType) isPassed() bool {
	return m >= 70
}

type passChecker interface {
	isPassed() bool
}

func showResult(score passChecker) {
	if score.isPassed() {
		fmt.Println("Pass!")
	} else {
		fmt.Println("Fail...")
	}
}

func showResult2(score any) {
	// v, ok := score.(englishScoreType)
	// if ok {
	// 	if v >= 80 {
	// 		fmt.Println("Pass!")
	// 	} else {
	// 		fmt.Println("Fail...")
	// 	}
	// }
	if v, ok := score.(englishScoreType); ok {
		if v >= 80 {
			fmt.Println("Pass!")
		} else {
			fmt.Println("Fail...")
		}
		return
	}

	if v, ok := score.(mathScoreType); ok {
		if v >= 80 {
			fmt.Println("Pass!")
		} else {
			fmt.Println("Fail...")
		}
		return
	}

	fmt.Println("Invalid score!")
}

func showResult3(score any) {
	switch v := score.(type) {
	case englishScoreType:
		if v >= 80 {
			fmt.Println("Pass!")
		} else {
			fmt.Println("Fail...")
		}
	case mathScoreType:
		if v >= 80 {
			fmt.Println("Pass!")
		} else {
			fmt.Println("Fail...")
		}
	default:
		fmt.Println("Invalid score!")
	}
}

func main() {
	var productName nameString = "Banana"

	showProductInfo(productName)

	productName.showInfo()
	// productName = productName.update("Apple")
	productName.update("Apple")
	productName.showInfo()

	englishScore := scoreType{
		subject: "English",
		points:  60,
	}

	fmt.Println(englishScore.subject)
	fmt.Println(englishScore.points)
	fmt.Println(englishScore.isPassed())

	// taro := studentType{
	// 	name: "Taro",
	// 	score: scoreType{
	// 		subject: "Math",
	// 		points:  85,
	// 	},
	// }

	// fmt.Println(taro.name)
	// fmt.Println(taro.score.subject)
	// fmt.Println(taro.score.points)
	// fmt.Println(taro.score.isPassed())

	taro := studentType{
		name: "Taro",
		scoreType: scoreType{
			subject: "Math",
			points:  85,
		},
	}

	fmt.Println(taro.name)
	fmt.Println(taro.subject)
	fmt.Println(taro.points)
	fmt.Println(taro.isPassed())

	englishScore2 := englishScoreType(90)
	mathScore2 := mathScoreType(50)

	fmt.Println(englishScore2.isPassed())
	fmt.Println(mathScore2.isPassed())
	showResult(englishScore2)
	showResult(mathScore2)

	showResult2(englishScore2)
	showResult2(mathScore2)
	showResult2("OK")

	showResult3(englishScore2)
	showResult3(mathScore2)
	showResult3("OK")
}
