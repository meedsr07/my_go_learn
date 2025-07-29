package main

import "fmt"

func test(str string) (string,string) {
	x:=""
	y:=""
	x=str[:5]
	y=str[5:]

	return x,y
}

func main() {
	fmt.Print(test("Hello, World!"))
}
