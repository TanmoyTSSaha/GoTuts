package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")
	
	// SLICES PART 1
	// var fruitList = []string{"Apple", "Tomato", "Peach"}
	
	// fmt.Printf("Type of fruitList is: %T\n", fruitList)
	
	// fruitList = append(fruitList, "Mango", "Banana")
	// fmt.Println(fruitList)

	// fruitList = append(fruitList[1:3])
	// fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0]=234
	highScores[1]=367
	highScores[2]=356
	highScores[3]=122

	highScores = append(highScores, 567, 4783, 4673, 6573)
	// fmt.Println("Is this Slice sorted: ", sort.IntsAreSorted(highScores))

	sort.Ints(highScores)

	// fmt.Println(highScores)
	// fmt.Println("Is this Slice sorted: ", sort.IntsAreSorted(highScores))

	// SLICES PART 2
	var courses = []string{"reactjs", "javascript", "typescript", "swift", "flutter", "python", "ruby", "go"}

	fmt.Println(courses)

	var index int = 3

	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)
}