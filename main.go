package main

import "fmt"

func main() {

	// variable types
	var username string = "Prathis"
	fmt.Println(username)
	fmt.Printf("variable is of the type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of the type: %T \n", isLoggedIn)

	var smallval uint16 = 8376
	fmt.Println(smallval)
	fmt.Printf("variable is of the type: %T \n", smallval)

	var float4 float32 = 83.78678687
	fmt.Println(float4)
	fmt.Printf("variable is of the type: %T \n", float4)
}
