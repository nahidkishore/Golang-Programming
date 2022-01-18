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

}