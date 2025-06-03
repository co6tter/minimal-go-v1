package main

import "fmt"

func addScore(pScore *int) {
	*pScore++
}

type student struct {
	name    string
	subject string
	score   int
}

func addScore2(pStudent *student) {
	pStudent.score++
}

func updateScoresArray(pScores *[3]int) {
	pScores[0] = 100
}

func updateScoresSlice(pScores []int) {
	pScores[0] = 100
}

func updateScoresMap(pScores map[string]int) {
	pScores["english"] = 100
}

func main() {
	score := 70
	// var pScore *int
	// pScore = &score
	pScore := &score

	fmt.Println(score)   // 70
	fmt.Println(&score)  // address
	fmt.Println(pScore)  // address
	fmt.Println(*pScore) // 70

	addScore(pScore)
	fmt.Println(score)

	taro := student{
		name:    "Taro",
		subject: "Math",
		score:   40,
	}
	addScore2(&taro)
	fmt.Println(taro.score)

	scoresArray := [3]int{10, 20, 30}
	scoresSlice := []int{40, 50, 60}
	scoresMap := map[string]int{
		"english": 70,
		"math":    80,
	}

	updateScoresArray(&scoresArray)
	updateScoresSlice(scoresSlice)
	updateScoresMap(scoresMap)

	fmt.Println("scoresArray:", scoresArray)
	fmt.Println("scoresSlice:", scoresSlice)
	fmt.Println("scoresMap:", scoresMap)

	score2 := 70
	var pScore2 *int
	fmt.Println(score2)
	if pScore2 != nil {
		fmt.Println(*pScore2)
	} else {
		fmt.Println("pScore2 is nil")
	}
}
