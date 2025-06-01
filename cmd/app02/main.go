package main

import "fmt"

func main() {
	// for {
	// 	var password string
	// 	fmt.Println("Password?")
	// 	fmt.Scanf("%s", &password)

	// 	if password == "he11o" {
	// 		fmt.Println("Password matched")
	// 		break
	// 	}
	// }

	var password string
	for password != "he11o" {
		fmt.Println("Password?")
		fmt.Scanf("%s", &password)
	}
	fmt.Println("Password matched")

	// isLoggedIn := 0 暗黙的に真偽値に変換はされない
	isLoggedIn := false

	if isLoggedIn {
		fmt.Println("User logged in.")
	} else {
		fmt.Println("User not logged in.")
	}
}
