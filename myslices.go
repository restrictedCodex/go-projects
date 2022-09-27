package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	fmt.Println("Welcome to Slices")

// 	var fruitList = []string{"Apple", "grape", "watermelon"}
// 	fmt.Printf("Type of data: %T \n", fruitList)

// 	fruitList = append(fruitList, "Mango")
// 	fmt.Println(fruitList)

// 	fruitList = append(fruitList[2:])
// 	fmt.Println(fruitList)

// 	highScore := make([]int, 4)

// 	highScore[0] = 234
// 	highScore[1] = 994
// 	highScore[2] = 114
// 	highScore[3] = 444

// 	highScore = append(highScore, 456, 677)

// 	sort.Ints(highScore)

// 	fmt.Println(highScore)

// 	//how to remove value from slice

// 	var courses = []string{"react", "swift", "js", "kotlin", "c++"}
// 	fmt.Println(courses)
// 	var index int = 1
// 	courses = append(courses[:index], courses[index+1:]...)
// 	fmt.Println(courses)
// }
