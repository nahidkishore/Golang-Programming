package main

import (
	"fmt"
)

func main() {
	test := "hello"
	fmt.Println(test)

	x := 0
	for x < 5 {
		fmt.Println("value of x is ", x)
		x++ // x=x+1

	}
	fmt.Println("for loop comming ********************************")
	for i := 0; i < 5; i++ {
		fmt.Println("value of i is:", i)
	}
	// array
	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])

	}
	for index, val := range names {
		fmt.Printf("the value at pos %v is %v \n", index, val)
		val = "new string"
	}

	for _, val := range names {
		fmt.Print(val, ",")
		val = "new string"
	}

	// changing val in a loop does not mutate the original slice
	fmt.Println(names)


}
