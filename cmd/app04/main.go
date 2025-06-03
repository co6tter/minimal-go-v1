package main

import "fmt"

type student struct {
	name    string
	subject string
	score   int
}

func main() {
	// var scores [3]int
	// scores[0] = 70
	// scores[1] = 90
	// scores[2] = 80

	// var scores = [3]int{70, 90, 80}
	// scores := [3]int{70, 90, 80}
	// scores := [...]int{70, 90, 80}

	// スライス
	// 後から要素を追加できる
	scores := []int{70, 90, 80}
	scores = append(scores, 100, 60)

	moreScores := []int{60, 75, 82, 40, 55}
	scores = append(scores, moreScores...)

	fmt.Println(scores[0])
	fmt.Println(scores[1])
	fmt.Println(scores[2])
	fmt.Println(scores)
	fmt.Println(len(scores))

	for i := 0; i < len(scores); i++ {
		fmt.Println(scores[i])
	}

	for i, v := range scores {
		fmt.Println(i, v)
	}

	for _, v := range scores {
		fmt.Println(v)
	}

	// スライス式
	fmt.Println(moreScores[0:3])
	fmt.Println(moreScores[:2])
	fmt.Println(moreScores[3:])
	fmt.Println(moreScores[:])

	// var scores2 map[string]int
	// scores2 = make(map[string]int)
	// var scores2 = make(map[string]int)
	// scores2["english"] = 80
	// scores2["math"] = 70
	scores2 := map[string]int{
		"english": 80,
		"math":    70,
	}
	scores2["physics"] = 90
	delete(scores2, "math")

	fmt.Println(scores2)
	fmt.Println(len(scores2))
	fmt.Println(scores2["history"]) // エラーにならずゼロ値が返る

	// カンマokイディオム
	// キーが存在するかどうかを確認する
	v, ok := scores2["history"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("history is not found")
	}

	scores3 := map[string][]int{
		"english": []int{80, 40},
		"math":    []int{40, 50, 60},
	}

	for subject, scoresSlice := range scores3 {
		fmt.Println(subject)
		for _, score := range scoresSlice {
			fmt.Println(score)
		}
	}

	// var taro student
	// taro.name = "Taro"
	// taro.subject = "Math"
	// taro.score = 40
	taro := student{
		name:    "Taro",
		subject: "Math",
		// score:   40, // 省略するとゼロ値が入る
	}

	fmt.Println(taro.name)
	fmt.Println(taro)

}
