package main


type Person struct {
	name string;
	age int;
	isMarried bool
}

func main(){

	var p1 Person
	p1.name ="Shushant"
	p1.age=19
	p1.isMarried = false

	//shorthand way
	p2 := Person{
		name: "Ram",
		age: 20,
		isMarried: true,
	}

	//new keyword
	var p3 = new(Person)
	p3.name ="Bheem"
	p3.age=60
	p3.isMarried=true

}