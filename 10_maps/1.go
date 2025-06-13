package main

import "fmt"

// Maps => unordered key value pair
func printInfo(studentInfo map[string]int){
	for name,marks := range studentInfo{
		fmt.Printf("%s) = %d\n",name,marks)
	}
}

func main(){

	studentRoll := make(map[int]string)

	studentRoll[1] = "Shushant"
	studentRoll[2] = "Manjistha"
	studentRoll[3] = "KV"

	studentMarks := map[string]int{
		"Shushant":98,
		"Raj":99,
		"Shubham":90,
	}

	printInfo(studentMarks)

	//delete any entry
	delete(studentRoll,3)
	//printInfo(studentRoll)

	//checks if exist
	Value , Exist :=studentRoll[3]
	fmt.Printf("Value : %s exist: %t",Value,Exist)

}