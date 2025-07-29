package main

import ("fmt")


func main() {
	x := "1234"
	y := "4567"
	x1 :=Atoi(x)
	y1 :=Atoi(y)
	fmt.Print( x1 + y1)

}

func Atoi(s string) int {
	nbr := 0
	for _, v := range s {
		nbr = nbr*10 + int(v-48)
	}
	return nbr
}
