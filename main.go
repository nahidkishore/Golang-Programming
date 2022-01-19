package main
import "fmt"
func main() {
fmt.Println("Hello world!, Go is comming .....")
// Strings concept
var nameOne string ="naima"
var nameTwo = "nahid"
var nameThree string 
fmt.Println(nameOne, nameTwo,nameThree)
nameOne="test"
nameThree="browser"
fmt.Println(nameOne, nameTwo,nameThree)
// := meaning var
nameFour := "chrome"
fmt.Println(nameFour)

// ints
var ageOne int =26
var ageTwo=30
ageThree :=32
fmt.Println(ageOne, ageTwo,ageThree)
// floats
// bits & memory 
// link ----https://go.dev/ref/spec#Numeric_types
var numOne int8=25 // 127
var numTwo int8=127 // error 127 correct
var numThree uint16=256

var scoreOne float32=25.77
var scoreTwo float64=7376473864384623843.8346346
scoreThree :=1.6
fmt.Println(numOne, numTwo,numThree)
fmt.Println( scoreOne,scoreTwo,scoreThree)

// print new line
fmt.Print("hello, ")
fmt.Print("world! \n")
fmt.Print("new line from here \n")

// println 
age:=35
name:= "nahid"
fmt.Println("age is ", age ,"and name is", name)
// Printf (formatted string), %_ = format specifier
fmt.Printf("my age is %v and my name is %v \n", age, name) // %v = value in default format
fmt.Printf("my age is %q and my name is %q \n", age, name) // %q = quotes
fmt.Printf("age is of type %T \n", age)                    // %T is the type
fmt.Printf("you scored %f points! \n", 225.55)             // %f = float format
fmt.Printf("you scored %0.1f points! \n", 225.55)          // %0.2f = float with 2 decimal points

// Array and slice

// Array
// var ages [3]int = [3]int{20, 25, 30}
var ages = [3]int {20,30,40}
fmt.Println(ages,len((ages)))

names:=[4]string{"nahid","kishore","Hridoy","test"}
fmt.Println(names,len((names)))
// slices (use arrays under the hood)
var scores = []int{100, 50, 60}
scores[2] = 25
scores = append(scores, 85)

fmt.Println(scores, len(scores))
// slice ranges
rangeOne := names[1:4] // doesn't include pos 4 element
rangeTwo := names[2:]  //includes the last element
rangeThree := names[:3]

fmt.Println(rangeOne, rangeTwo, rangeThree)
fmt.Printf("the type of rangeOne is %T \n", rangeOne)

rangeOne = append(rangeOne, "koopa")
fmt.Println(rangeOne)
}